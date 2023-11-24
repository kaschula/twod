package physics

import "math"

type Circle struct {
	body *Body
	// the edge which the circle at 0 degress points
	edgePoint V
	// the edge point the current angle points to
	rotatedEdge V
	radius      int
	shape       string
}

func (c *Circle) PreviousAngle() V {
	return c.body.PreviousAngle()
}

func (c *Circle) Accelerate(incrementBy float64) {
	c.body.Accelerate(incrementBy)
}

func (c *Circle) Drag(by float64) {
	c.body.Drag(by)
}

func (c *Circle) Acceleration() float64 {
	return c.body.Acceleration()
}

func (c *Circle) HasBoundingCollisionOn() bool {
	return c.body.HasBoundingCollisionOn()
}

func (c *Circle) BoundingRadius() int {
	return c.body.BoundingRadius()
}

func (c *Circle) BoundingCollidesWith(other RigidBody) bool {
	return c.body.BoundingCollidesWith(other)
}

func (c *Circle) SetIsColliding(isColliding bool) {
	c.body.SetIsColliding(isColliding)
}

func (c *Circle) IsColliding() bool {
	return c.body.IsColliding()
}

func NewCircle(center V, radius, boundingRadius int, os ...Option) *Circle {
	//startPoint := New(center.X(), center.Y()-float64(radius))
	c := &Circle{
		body:   NewRigidBody(center, boundingRadius, os...),
		radius: radius,
		shape:  ShapeCircle,
	}

	c.calculateEdges()

	return c
}

func (c *Circle) calculateEdges() {
	startPoint := c.body.center.Add(New(0, -float64(c.radius)))
	c.edgePoint = startPoint
	c.rotatedEdge = c.edgePoint.Rotate(c.body.center, c.body.angle)
}

// func (c *Circle) RotateTo(point V) {
//}

func (c *Circle) Vertices() []V {
	panic("not implemented")
	return nil
}

func (c *Circle) Location() V {
	return c.body.center
}

func (c *Circle) Update() (bool, V) {
	if c.body.fixed {
		return false, ZeroVector
	}

	hasMoved, moved := c.body.Update()
	c.calculateEdges()
	return hasMoved, moved
}

func (c *Circle) Radius() int {
	return c.radius
}

func (c *Circle) Radius64() float64 {
	return float64(c.radius)
}

func (c *Circle) StartPoint() V {
	return c.edgePoint
}

func (c *Circle) RotatedEdgePoint() V {
	return c.rotatedEdge
}

func (c *Circle) Rotate(by Degree) {
	c.body.Rotate(by)
}

func (c *Circle) RotateTo(p V) {
	c.body.RotateTo(p)
}

func (c *Circle) SetRotation(angle Degree) {
	c.body.SetRotation(angle)
}

func (c *Circle) Direction() Degree {
	return c.body.Direction()
}

func (c *Circle) Move(by ...V) {
	c.body.Move(by...)
}

func (c *Circle) Transpose(to V) {
	c.body.Transpose(to)
	c.calculateEdges()
}

func (c *Circle) Collision(other RigidCircle) *Collision {
	fromC1toC2 := c.Location().Sub(other.Location())
	fromC2toC1 := other.Location().Sub(c.Location())
	rSum := c.Radius64() + other.Radius64()
	distance := fromC1toC2.Length()
	if sqrtR := math.Sqrt(rSum * rSum); distance > sqrtR {
		//if distance > float64(rSum) {
		// no collision
		return nil
	}

	if distance != 0 {
		c2ToC1Normal := fromC2toC1.Normalize()
		// direction to move c1 to resolve collision
		normal := fromC1toC2.Normalize()
		start := c.Location().Add(c2ToC1Normal.Scale(float64(c.Radius())))
		depth := float64(rSum) - distance

		return NewCollision(start, normal, depth)
	}

	if !c.Location().Equal(other.Location()) {
		return nil
	}

	// if circles have the same location then move negatively along the Y using the largest radius
	normal := New(0, -1)
	if c.Radius64() > other.Radius64() {
		return NewCollision(c.Location().Add(New(0, c.Radius64())), normal, rSum)
	}

	return NewCollision(other.Location().Add(New(0, other.Radius64())), normal, rSum)
}
