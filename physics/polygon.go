package physics

import "fmt"

type Poly struct {
	sides, width        int
	body                RigidBody
	zeroCornerPoints    Matrix2d
	rotatedCornerPoints Matrix2d
	faces               Matrix2d
}

func (p *Poly) Sides() int {
	return p.sides
}

func (p *Poly) Width() int {
	return p.width
}

func (p *Poly) Radiuses() []*LineSegment {
	radiues := []*LineSegment{}
	for _, vertex := range p.rotatedCornerPoints {
		radiues = append(radiues, NewLineSegment(p.Location(), vertex))
	}

	return radiues
}

func (p *Poly) Edges() []*LineSegment {
	vertices := p.rotatedCornerPoints
	lastIndex := len(vertices) - 1
	edges := []*LineSegment{}
	for i := 0; i < len(vertices); i++ {
		nextI := i + 1
		if i == lastIndex {
			nextI = 0
		}

		edges = append(edges, NewLineSegment(vertices[i], vertices[nextI]))
	}

	return edges
}

func (p *Poly) Accelerate(incrementBy float64) {
	p.body.Accelerate(incrementBy)
}

func (p *Poly) Drag(by float64) {
	p.body.Drag(by)
}

func (p *Poly) Rotate(by Degree) {
	p.body.Rotate(by)
}

func (p *Poly) SetRotation(degrees Degree) {
	p.body.SetRotation(degrees)
}

func (p *Poly) RotateTo(point V) {
	p.body.RotateTo(point)
}

func (p *Poly) Move(by ...V) {
	p.body.Move(by...)
}

func (p *Poly) Update() (bool, V) {
	hasMoved, movedBy := p.body.Update()
	p.calculateEdgesAndRotate()

	return hasMoved, movedBy
}

func (p *Poly) PreviousAngle() V {
	return p.body.PreviousAngle()
}

func (p *Poly) Location() V {
	return p.body.Location()
}

func (p *Poly) Acceleration() float64 {
	return p.body.Acceleration()
}

func (p *Poly) Direction() Degree {
	return p.body.Direction()
}

func (p *Poly) HasBoundingCollisionOn() bool {
	return p.body.HasBoundingCollisionOn()
}

func (p *Poly) BoundingRadius() int {
	return p.body.BoundingRadius()
}

func (p *Poly) BoundingCollidesWith(other RigidBody) bool {
	return p.body.BoundingCollidesWith(other)
}

func (p *Poly) SetIsColliding(isColliding bool) {
	p.body.SetIsColliding(isColliding)
}

func (p *Poly) IsColliding() bool {
	return p.body.IsColliding()
}

func (p *Poly) Vertices() []V {
	return p.rotatedCornerPoints
}

func (p *Poly) Faces() []V {
	if len(p.faces) == 0 {
		p.setFaces()
	}

	return p.faces
}

func (p *Poly) setFaces() {
	faces := make([]V, p.sides)
	vertices := p.Vertices()

	for i := 0; i < p.sides; i++ {
		adjacentCornerStart := indexCycler(p.sides, i+1)
		adjacentCornerEnd := indexCycler(p.sides, i+2)

		faces[i] = vertices[adjacentCornerStart].Sub(vertices[adjacentCornerEnd]).Normalize()
	}

	p.faces = faces
}

func (p *Poly) RotatedEdgePoint() V {
	edgePoint := p.body.Location().Add(New(0, -float64(p.width)))
	edgePoint = edgePoint.Rotate(p.body.Location(), p.body.Direction())

	return edgePoint
}

func (p *Poly) Transpose(to V) {
	p.body.Transpose(to)

	p.zeroCornerPoints = setZeroCornerPoints(p.sides, p.width, p.body.Location())
	p.rotatedCornerPoints = p.zeroCornerPoints.Rotate(p.body.Location(), p.body.Direction())
	p.setFaces()
}

func (p *Poly) calculateEdgesAndRotate() {
	p.zeroCornerPoints = setZeroCornerPoints(p.sides, p.width, p.body.Location())
	p.rotatedCornerPoints = p.zeroCornerPoints.Rotate(p.body.Location(), p.body.Direction())
	p.setFaces()
}

func NewPoly(center V, sides, width int, option ...Option) *Poly {
	if sides < 3 {
		fmt.Println("poly must have at least 3 sides")
		sides = 3
	}

	defaultRB := defaultRigidBody{}
	defaultRB.load(option...)

	rb := NewRigidBody(center, defaultRB.boundingBody, option...)

	// get zero points
	// rotate by angle
	zeroEdges := setZeroCornerPoints(sides, width, center)

	p := &Poly{
		sides:               sides,
		width:               width,
		body:                rb,
		zeroCornerPoints:    zeroEdges,
		rotatedCornerPoints: zeroEdges,
	}

	p.calculateEdgesAndRotate()

	return p
}

func setZeroCornerPoints(sides, width int, center V) Matrix2d {

	// calculate the angle each pint rotates by
	// add each aedge to a matix

	betweenEach := 360 / sides
	m := Matrix2d{}
	starting := center.Add(New(0, -float64(width)))
	for i := 0; i < sides; i++ {
		if i == 0 {
			m = append(m, starting)
			continue
		}

		degree := Degree(betweenEach * (i))
		edge := starting.Rotate(center, degree)

		m = append(m, edge)
	}

	return m
}
