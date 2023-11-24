package physics

func (b *Body) HasBoundingCollisionOn() bool {
	return b.boundingCollisionRadius > 0
}

func (b *Body) BoundingRadius() int {
	return b.boundingCollisionRadius
}

func (b *Body) IsColliding() bool {
	return b.colliding
}

// BoundingCollidesWith
// todo currently this sets the rigid body as colliding this could be bugg when multiplay things collide
// When checking lots of bodies make a collision body collection and set the colliding to false for everything then check
// if no collision do not update to false as previous collision will still be valid
// todo refactor this so that even to objects can be used this way
func (b *Body) BoundingCollidesWith(other RigidBody) bool {
	vFrom1to2 := other.Location().Sub(b.center)
	radiusSum := b.boundingCollisionRadius + other.BoundingRadius()
	var dist = vFrom1to2.Length()

	// todo reverse this so you can just return the condition
	if dist > float64(radiusSum) {
		b.SetIsColliding(false)
		other.SetIsColliding(false)
		return false
	}

	b.SetIsColliding(true)
	other.SetIsColliding(true)

	return true
}
