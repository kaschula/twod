package draw

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	ebitenvector "github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kaschula/twod/physics"
	"image"
	"image/color"
	"math"
)

// could be an effeicny gain to use a sinle sub image

// emptySubImage is an internal sub image of emptyImage.
// Use emptySubImage at DrawTriangles instead of emptyImage in order to avoid bleeding edges.
var (
	emptyImage    = ebiten.NewImage(3, 3)
	emptySubImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	emptyImage.Fill(color.White)
}

type VectorImage struct {
	// todo remove start
	vertices []physics.V
	colour   color.Color
}

func NewVectorImage(vertices []physics.V, colour color.Color) *VectorImage {
	return &VectorImage{vertices: vertices, colour: colour}
}

func VectorFill(screen *ebiten.Image, v *VectorImage) {
	var path ebitenvector.Path
	start := v.vertices[0]
	path.MoveTo(start.X32(), start.Y32())
	for _, p := range v.vertices[1:] {
		path.LineTo(p.X32(), p.Y32())
	}

	drawVectorPath(screen, path, v.colour)
}

func ToVectorShape(poly physics.RigidPoly, c color.Color) *VectorImage {
	vertices := poly.Vertices()

	var vis []physics.V
	for i := range vertices {
		vis = append(vis, vertices[i])
	}

	return NewVectorImage(vis, c)
}

// RigidBody
//todo interfcae to RigidBodyPoly

func Polygon(screen *ebiten.Image, p physics.RigidPoly, c color.Color) {
	VectorFill(screen, ToVectorShape(p, c))

	// todo make this an option
	//edgePoint := p.RotatedEdgePoint()
	//x1, y1, x2, y2 := edgePoint.WithToFloat64(p.Location())
	//ebitenutil.DrawLine(screen, x1, y1, x2, y2, startPointShader(c))
}

// todo make drawAngle an option function
func Rectangle(screen *ebiten.Image, p physics.Rectangle, c color.Color, drawAngle bool) {
	VectorFill(screen, ToVectorShape(p, c))

	if drawAngle {
		edgePoint := p.RotatedEdgePoint()
		x1, y1, x2, y2 := edgePoint.WithToFloat64(p.Location())
		ebitenutil.DrawLine(screen, x1, y1, x2, y2, startPointShader(c))
	}
}

// PolygonFaces
// todo Project lines out from ceter in direction of faces
func PolygonFaces(screen *ebiten.Image, p physics.RigidPoly, c color.Color) {
	center := p.Location()

	for _, faceDirection := range p.Faces() {
		scaled := faceDirection.Scale(60)
		lineEnd := center.Add(scaled)

		x1, y1, x2, y2 := center.WithToFloat64(lineEnd)

		ebitenutil.DrawLine(screen, x1, y1, x2, y2, c)
	}
}

func WireFrame(screen *ebiten.Image, p physics.RigidPoly, c color.Color) {
	for _, edge := range p.Edges() {
		ebitenutil.DrawLine(screen, edge.Start().X(), edge.Start().Y(), edge.End().X(), edge.End().Y(), c)
	}
}

func Boundry(screen *ebiten.Image, rb physics.RigidBody, c color.Color) {
	drawArc(screen, rb.Location(), float64(rb.BoundingRadius()), c)
}

func Circle(screen *ebiten.Image, circle physics.RigidCircle, c color.Color) {
	drawArc(screen, circle.Location(), circle.Radius64(), c)

	x1, y1, x2, y2 := circle.RotatedEdgePoint().WithToFloat64(circle.Location())

	ebitenutil.DrawLine(screen, x1, y1, x2, y2, startPointShader(c))
}

// Vector draws a vector as 10 pixel dot
func Vector(screen *ebiten.Image, v physics.V, c color.Color, size ...float64) {
	var pointSize float64 = 10
	if len(size) > 0 {
		pointSize = size[0]
	}

	drawArc(screen, v, pointSize, c)
}

func Line(screen *ebiten.Image, start, end physics.V, c color.Color) {
	x1, y1, x2, y2 := start.WithToFloat64(end)

	ebitenutil.DrawLine(screen, x1, y1, x2, y2, startPointShader(c))
}

func drawArc(screen *ebiten.Image, point physics.V, radius float64, c color.Color) {
	var path ebitenvector.Path

	cx, cy := float32(point.X()), float32(point.Y()) // circle center and radius
	theta1 := math.Pi * float64(0.1) / 180

	theta2 := math.Pi * float64(0.1) / 180 / 3
	path.MoveTo(550, 100)
	path.Arc(cx, cy, float32(radius), float32(theta1), float32(theta2), ebitenvector.Clockwise)

	drawVectorPath(screen, path, c)

}

func drawVectorPath(screen *ebiten.Image, path ebitenvector.Path, c color.Color) {
	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}
	rd, g, b, alpha := c.RGBA()
	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = float32(uint8(rd)) / float32(255)
		vs[i].ColorG = float32(uint8(g)) / float32(255)
		vs[i].ColorB = float32(uint8(b)) / float32(255)
		vs[i].ColorA = float32(uint8(alpha)) / float32(255)
	}
	screen.DrawTriangles(vs, is, emptySubImage, op)
}

// todo draw polygon wire frame using lines to draw the corners

// BoundBox
// todo
// draw 20 dots rotated around an objects center in a 360 to show the bounding box

func startPointShader(orginal color.Color) color.Color {
	rd, g, b, alpha := orginal.RGBA()

	rd8, g8, b8, alpha8 := uint8(rd), uint8(g), uint8(b), uint8(alpha)

	rd8 += 150
	g8 += 150
	b8 += 150

	return color.RGBA{
		R: rd8,
		G: g8,
		B: b8,
		A: alpha8,
	}
}

func RGBA(r, g, b, a uint8) color.Color {
	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
