package collision

import "github.com/kaschula/twod/physics"

// This file contains collisions functions for rectangles where there is no rotation. AABB type collisons
// todo move all collisions into the package and return proper collision objects from all functions

// PointInRectangle angle must be 0 degrees
func PointInRectangle(r physics.Rectangle, point physics.V) bool {
	if r.Direction() != physics.Degree(0) {
		panic("angle is not 0")
		return false
	}

	x, y := point.X(), point.Y()

	rX := float64(r.TopLeft().X())
	rY := float64(r.TopLeft().Y())
	rW := float64(r.Width())
	rH := float64(r.Height())

	return x >= rX && // right of the left edge AND
		x <= rX+rW && // left of the right edge AND
		y >= rY && // below the top AND
		y <= rY+rH // above the bottom
}

func RectRect(r1, r2 physics.Rectangle) bool {
	//
	//rect1 := r1.Location()
	//rect2 := r2.Location()
	//
	//return rect1.x < rect2.x + rect2.w &&
	//	rect1.x + rect1.w > rect2.x &&
	//	rect1.y < rect2.y + rect2.h &&
	//	rect1.h + rect1.y > rect2.y
	//

	return false
}
