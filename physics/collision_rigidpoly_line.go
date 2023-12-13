package physics

// CollisionRigidPolyDiagonalLine takes to rigid polygons and checks if the radius lines of A
// intersect with line B and vice versa
func CollisionRigidPolyDiagonalLine(a, b RigidPoly) *Collision {
	touch := lineSetCollision(a.Radiuses(), b.Edges())
	if !touch.Empty() {
		return NewCollisionFromTouch(touch)
	}

	// check the inverse
	touch = lineSetCollision(b.Radiuses(), a.Edges())
	if !touch.Empty() {
		return NewCollisionFromTouch(touch).ReverseDirection()
	}

	return nil
}

func CollisionRigidPolyDiagonalLineValue(a, b RigidPoly) Collision {
	if a == nil || b == nil {
		return emptyCollision()
	}

	touch := lineSetCollision(a.Radiuses(), b.Edges())
	if !touch.Empty() {
		return NewCollisionFromTouchValue(touch)
	}

	// check the inverse
	touch = lineSetCollision(b.Radiuses(), a.Edges())
	if !touch.Empty() {
		return NewCollisionFromTouchValue(touch).ReverseDirectionValue()
	}

	return emptyCollision()
}

// lineSetCollision Loops through to sets of lines to determine if any lines intersect
func lineSetCollision(lineSetA, lineSetB []*LineSegment) Touch {
	for _, lineA := range lineSetA {
		for _, lineB := range lineSetB {
			touch := GetLineIntersection(lineA, lineB)
			if !touch.Empty() {
				return touch
			}
		}
	}

	return EmptyTouch
}

func CollisionRigidPolyWithEdges(a RigidPoly, edges []*LineSegment) *Collision {
	touch := lineSetCollision(a.Edges(), edges)
	if !touch.Empty() {
		return NewCollisionFromTouch(touch)
	}

	return nil
}
