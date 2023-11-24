package physics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterfaces(t *testing.T) {
	var body RigidBody = NewRigidBody(New(1, 1), 1)
	var rectangle Rectangle = NewRectangle(New(1, 1), 10, 10, 0)
	var circle RigidCircle = NewCircle(New(1, 1), 1, 1)

	var triange RigidPoly = NewPoly(New(1, 1), 1, 50)

	assert.NotNil(t, body)
	assert.NotNil(t, rectangle)
	assert.NotNil(t, circle)
	assert.NotNil(t, triange)
}
