package physics

import (
	"fmt"
	"math"
)

// This will probably go into a collision package

const (
	ShapeRectangle = "rectangle"
	ShapeCircle    = "circle"
	Fixed          = true
	UnFixed        = false

	// adjustForOriginAngle Atan2 calculate the angle from the horizontal plain.
	// a point on the y = 0 plain would have a0 degree rotation
	// This means we have adjust the angle by 90 degrees to have 0 point up along the x not the y
	// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/atan2
	adjustForOriginAngle = 90
)

type Degree float64

func (d Degree) ToRadian() Radian {
	// This is technically wrong should be 180 not -180 but when playing -180 works
	return Radian(d * (math.Pi / -180))
}

func (d Degree) ToFloat64() float64 {
	return float64(d)
}

type Radian float64

func (r Radian) ToDegree() Degree {
	return Degree(r * (180 / math.Pi))
}

func (r Radian) ToFloat64() float64 {
	return float64(r)
}

type Body struct {
	previousCenter, center V
	// degrees
	angle                              Degree
	fixed                              bool
	boundingCollisionRadius            int
	colliding                          bool
	acceleration, speed, maxAccelerate float64
	moves                              Matrix2d
}

func (b *Body) Transpose(to V) {
	prevoius := b.center
	b.previousCenter = prevoius
	b.center = to
}

func (b *Body) Accelerate(incrementBy float64) {
	b.acceleration += incrementBy

	if b.acceleration > b.maxAccelerate {
		b.acceleration = b.maxAccelerate
	}

	if b.acceleration < -b.maxAccelerate {
		b.acceleration = -b.maxAccelerate
	}
}

func (b *Body) Direction() Degree {
	return b.angle
}

// Drag by should always be positive
// if acceleration is negative then drag will increase to get closer to 0
// if acceleration is positive tehn drag decrease till 0
func (b *Body) Drag(by float64) {
	if by < 0 {
		panic(fmt.Sprintf("drag can not be less than zero, got %v", by))
	}

	if b.acceleration == 0 {
		return
	}

	if b.acceleration < 0 {
		b.acceleration += by
		if b.acceleration > 0 {
			b.acceleration = 0
		}
		return
	}

	// acceleration is positive
	b.acceleration -= by
	if b.acceleration < 0 {
		b.acceleration = 0
	}
}

func (b *Body) Acceleration() float64 {
	return b.acceleration
}

func (b *Body) Location() V {
	return b.center
}

//todo delete doesnt look like its being used
type RigidPolyCollides interface {
	SAT(other RigidPoly) *Collision
}

func (b *Body) SetIsColliding(isColliding bool) {
	b.colliding = isColliding
}

func (b *Body) RotateTo(point V) {
	direction := point.Sub(b.center)
	radians := math.Atan2(direction.Y(), direction.X())
	degrees := Radian(radians).ToDegree() + adjustForOriginAngle

	b.SetRotation(degrees)
}

func (b *Body) SetRotation(degrees Degree) {
	b.angle = Degree(Cycle360Float64(degrees.ToFloat64()))
}

func (b *Body) Rotate(byDegrees Degree) {
	updated64 := Cycle360Float64(b.angle.ToFloat64() + byDegrees.ToFloat64())
	b.angle = Degree(updated64)
}

// This is a bit broken
// This function should be in the update, private and should calculate the velocity based on speed
// speed should be a rigid body option
// add can be set
// same with acceleration
func (b *Body) velocity() V {
	// untested
	speed := b.speed
	acceleration := b.acceleration

	if speed == 0 {
		return ZeroVector
	}

	if acceleration == 0 {
		return ZeroVector
	}

	projectionPoint := b.center.Add(New(0, -10))
	rotated := projectionPoint.Rotate(b.center, b.angle)
	velocity := rotated.Sub(b.center).Normalize().Scale(speed * acceleration)

	return velocity
}

// todo this is not working so well try this in a new app where
// Update apply speed, angle and acceleration to the center
func (b *Body) Update() (bool, V) {
	previous := b.center

	// apply Moves
	totalMove := New(0, 0)
	totalMove = totalMove.Add(b.movesSum())

	// Applys
	speedAccelerationVelocity := b.velocity()
	totalMove = totalMove.Add(speedAccelerationVelocity)

	b.center = b.center.Add(totalMove)

	if previous.Equal(b.center) {
		return false, totalMove
	}
	//
	b.previousCenter = previous
	//b.center = b.center.Add(speedVelocity)
	return true, totalMove
}

// Transpose(to V)
// sub the B center to V, add it as a Transpose Force run and update

func (b *Body) Move(by ...V) {
	b.moves = append(b.moves, by...)
}

func (b *Body) Apply(by V) {

}

func (b *Body) PreviousAngle() V {
	return b.center.Sub(b.previousCenter)
}

func (b *Body) Previous() V {
	return b.previousCenter
}

func (b *Body) movesSum() V {
	sum := b.moves.Sum()
	b.moves = b.moves.Clear()

	return sum
}

type defaultRigidBody struct {
	acceleration, speed, maxAccelerate float64
	angle                              Degree
	fixed                              bool
	boundingBody                       int
}

type Option func(d *defaultRigidBody)

func (d *defaultRigidBody) load(os ...Option) {
	for _, o := range os {
		o(d)
	}
}

func WithAcceleration(a float64) Option {
	return func(d *defaultRigidBody) {
		d.acceleration = a
	}
}

func WithSpeed(s float64) Option {
	return func(d *defaultRigidBody) {
		d.speed = s
	}
}

func WithAngle(a Degree) Option {
	return func(d *defaultRigidBody) {
		d.angle = a
	}
}
func WithBoundingBody(b int) Option {
	return func(d *defaultRigidBody) {
		d.boundingBody = b
	}
}

func WithMaxAccelerate(ma float64) Option {
	return func(d *defaultRigidBody) {
		d.maxAccelerate = ma
	}
}

// with max speed
// with max accerlation

func NewRigidBody(center V, boundingCollisionRadius int, os ...Option) *Body {
	defaultValues := defaultRigidBody{}
	defaultValues.load(os...)

	// todo remove boundingCollisionRadius as argument and have option only
	if boundingCollisionRadius == 0 {
		boundingCollisionRadius = defaultValues.boundingBody
	}

	return &Body{
		center:                  center,
		angle:                   defaultValues.angle,
		fixed:                   defaultValues.fixed,
		speed:                   defaultValues.speed,
		maxAccelerate:           defaultValues.maxAccelerate,
		acceleration:            defaultValues.acceleration,
		boundingCollisionRadius: boundingCollisionRadius,
		colliding:               false,
	}
}
