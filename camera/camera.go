package camera

import "github.com/kaschula/twod/physics"

// develop a view port

type ViewPort struct {
	width, Height int
	pointInWorld  physics.V
	mapW, mapH    int
}

/*
The viewport will have set size, as its point moves it can be passed a list of rigid bodies

each rigid body will have a point inside the map world

If a rigid point falls within the viewport W and H then it is return in a list
and V offset is returns so that these rigid bodies can be draw on the screen
view port can be moved by moving the pointInWorld

*/

// CameraOffset returns the movement required for a body to remain a set point.
// Return V can be a applied to all other bodies in the scene to move them while the center body appears to stay in the same poisition
func Offset(centerBody physics.RigidBody, cameraFocus physics.V) physics.V {
	cameraOffset := cameraFocus.Sub(centerBody.Location())
	//offSetY := physics.New(0, cameraOffset.Y())
	return cameraOffset
}

// CameraOffsetY returns V to only adjust the Y axis
func OffsetY(centerBody physics.RigidBody, cameraFocus physics.V) physics.V {
	cameraOffset := Offset(centerBody, cameraFocus)
	return physics.New(0, cameraOffset.Y())
}
