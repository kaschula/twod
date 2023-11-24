package components

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	color2 "github.com/kaschula/twod/color"
	"github.com/kaschula/twod/draw"
	"github.com/kaschula/twod/math"
	"github.com/kaschula/twod/physics"
	"image/color"
	math2 "math"
)

// Sensor and object that shoots out rays from a given rigid body.
// This is used for detecting multiple nearby RigidBodies
type Sensor interface {
	Update(collisionEdges ...*physics.LineSegment)
	Draw(screen *ebiten.Image)
	Transpose(by physics.V)
	// Data should return slice size equal to number of rays
	Data() []float64
}

type NoSensor struct{}

// todo function to clear readings

func (n NoSensor) Update(collisionEdges ...*physics.LineSegment) {}
func (n NoSensor) Transpose(by physics.V)                        {}
func (n NoSensor) Draw(screen *ebiten.Image)                     {}
func (n NoSensor) Data() []float64                               { return []float64{} }

func NewSensor(body physics.RigidBody, rayCount int, rayLength float64, raySpread physics.Degree, color color.Color) *RaySensor {
	return &RaySensor{body: body, rayCount: rayCount, rayLength: rayLength, raySpread: raySpread, color: color}
}

type RaySensor struct {
	body      physics.RigidBody
	rayCount  int
	rayLength float64
	raySpread physics.Degree
	rays      []*physics.LineSegment
	color     color.Color
	// Line collisions between the collisionEdges and the sensor rays
	touches []physics.Touch
}

func (sensor *RaySensor) Data() []float64 {
	data := []float64{}
	for _, touch := range sensor.touches {
		data = append(data, touch.Offset)
	}

	return data
}

func (sensor *RaySensor) Update(collisionEdges ...*physics.LineSegment) {
	sensor.castRays()
	sensor.detectReadings(collisionEdges...)
}

// detectReadings records any line touches of the rays and the collisionEdges
func (sensor *RaySensor) detectReadings(collisionEdges ...*physics.LineSegment) {
	if len(collisionEdges) == 0 {
		// nothing to detect
		fmt.Println("nothing to detect")
		return
	}

	touches := make([]physics.Touch, len(sensor.rays))

	// todo this could be for range
	for i := 0; i < len(sensor.rays); i++ {
		touch := sensor.getReading(sensor.rays[i], collisionEdges)
		if !touch.Empty() {
			touches[i] = touch
		}
	}

	sensor.touches = touches
}

// getReading detches a touch using line intersectuion of a sensor ray and a collection of lines
// if a ray touches more than one line the closest touch is return.
func (sensor *RaySensor) getReading(ray *physics.LineSegment, collisionEdges []*physics.LineSegment) physics.Touch {
	touchWithSmallestOffset := physics.Touch{}

	for _, collisionEdge := range collisionEdges {
		touch := physics.GetLineIntersection(ray, collisionEdge)
		if touch.Empty() {
			continue
		}

		if touchWithSmallestOffset.Offset == 0 || touch.Offset < touchWithSmallestOffset.Offset {
			touchWithSmallestOffset = touch
		}
	}

	return touchWithSmallestOffset
}

// castRays builds the rays for the sensor based on the body direction and location
func (sensor *RaySensor) castRays() {
	rays := []*physics.LineSegment{}

	halfRaySpread := sensor.raySpread.ToRadian().ToFloat64() / 2
	start := sensor.body.Location()

	for i := 0; i < sensor.rayCount; i++ {
		rayAngle := math.Lerp(
			halfRaySpread,
			-halfRaySpread,
			RayTValue(i, sensor.rayCount-1),
		) + sensor.body.Direction().ToRadian().ToFloat64() // This rotates the center with car angle

		end := physics.New(
			start.X()-math2.Sin(rayAngle)*sensor.rayLength,
			start.Y()-math2.Cos(rayAngle)*sensor.rayLength,
		)

		rays = append(rays, physics.NewLineSegment(start, end))

	}
	sensor.rays = rays
}

func (sensor *RaySensor) Transpose(by physics.V) {
	// loop through rays
	// loop through touches
	for _, r := range sensor.rays {
		r.Transpose(by)
	}

	for _, touch := range sensor.touches {
		touch.Transpose(by)
	}
}

func (sensor *RaySensor) Draw(screen *ebiten.Image) {
	for _, r := range sensor.rays {
		draw.Line(screen, r.Start(), r.End(), sensor.color)
	}

	for _, touch := range sensor.touches {
		draw.Rectangle(screen, physics.NewRectangle(touch.Vector(), 6, 6, 0), color2.Green, false)
	}
}

// RayTValue used a T value in Lerp function, if there is only 1 Ray then it will always point forward
// else calculate the value
func RayTValue(index int, rayCount int) float64 {
	if rayCount == 1 {
		return 0.5
	}

	return float64(index) / float64(rayCount-1)
}
