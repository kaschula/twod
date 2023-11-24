package physics

func (r *Rect) BoundingCollidesWith(other RigidBody) bool {
	return r.Body.BoundingCollidesWith(other)
}

func (r *Rect) HasBoundingCollisionOn() bool {
	return r.Body.HasBoundingCollisionOn()
}

func (r *Rect) BoundingRadius() int {
	return r.Body.BoundingRadius()
}

func (r *Rect) SetIsColliding(isColliding bool) {
	r.Body.SetIsColliding(isColliding)
}

func (r *Rect) IsColliding() bool {
	return r.Body.IsColliding()
}
