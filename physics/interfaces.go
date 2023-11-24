package physics

type BoundingCollider interface {
	// todo rename
	HasBoundingCollisionOn() bool
	BoundingRadius() int
	BoundingCollidesWith(other RigidBody) bool
	SetIsColliding(isColliding bool)
	IsColliding() bool
}

// todo add a Transpose function that just moves all vector based data by an amount
type RigidBody interface {
	//todo
	//IsActive()

	Rotate(byDegrees Degree)
	SetRotation(degrees Degree)
	RotateTo(point V)
	// Move apply a vector to an object ignoring the angle or velocity of the rigid body similar to applying a force
	Move(by ...V)
	// Apply apply a force vector to a rigid body taking into account the angle and velocity of the body
	//Apply(by V)

	// Update returns if it has moved and how much by
	Update() (bool, V)
	PreviousAngle() V
	Location() V
	Direction() Degree
	// Accelerate increment add
	Accelerate(incrementBy float64)
	Acceleration() float64
	Drag(incrementBy float64)

	// Transpose Maintains angle and does not apply forces or speed moves center and edges to a point
	Transpose(to V)
	// set speed
	// Set acceleration
	BoundingCollider
}

type RigidPoly interface {
	RigidBody
	// Vertices these are the vectors of the corners of the object at its current rotation
	Vertices() []V
	Faces() []V

	// Lines are drawn between all the vertices
	Edges() []*LineSegment
	// RotatedEdgePoint return the point at the edge of the shape facing the angle
	RotatedEdgePoint() V
	Radiuses() []*LineSegment
	Sides() int
	Width() int
}

type RigidCircle interface {
	RigidBody
	Radius() int
	Radius64() float64
	StartPoint() V
	RotatedEdgePoint() V
	Collision(other RigidCircle) *Collision
}

type Rectangle interface {
	RigidPoly
	//Width() int
	Height() int
	TopLeft() V
	TopRight() V
	BottomLeft() V
	BottomRight() V
}
