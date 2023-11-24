package physics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBody_ImplementsRigidBody(t *testing.T) {
	var b RigidBody = NewRigidBody(New(0, 0), 0)
	assert.NotNil(t, b)
}

func TestDegree_ToRadian(t *testing.T) {
	s := func(v float64) string { return fmt.Sprintf("%.2f", v) }

	assert.Equal(t, s(-1.0), s(Degree(57.2958).ToRadian().ToFloat64()))
	assert.Equal(t, s(-1.5708), s(Degree(90).ToRadian().ToFloat64()))
	assert.Equal(t, s(-3.1415), s(Degree(180).ToRadian().ToFloat64()))
	assert.Equal(t, s(-5.7595), s(Degree(330).ToRadian().ToFloat64()))
}

func TestRadians_ToDegrees(t *testing.T) {
	s := func(v float64) string { return fmt.Sprintf("%.2f", v) }

	assert.Equal(t, s(57.2958), s(Radian(1.0).ToDegree().ToFloat64()))
	assert.Equal(t, s(90), s(Radian(1.5708).ToDegree().ToFloat64()))
	assert.Equal(t, s(179.99), s(Radian(3.1415).ToDegree().ToFloat64()))
	assert.Equal(t, s(330), s(Radian(5.7595).ToDegree().ToFloat64()))
}

var (
	left  = New(-1, 0)
	right = New(1, 0)
	up    = New(0, -1)
	down  = New(0, 1)
)

func TestBody_Update_move_by_one(t *testing.T) {
	body := NewRigidBody(New(10, 10), 0)

	body.Move(left)
	hasMoved, movedBy := body.Update()

	assert.True(t, hasMoved)
	assert.Equal(t, New(9, 10), body.center)
	assert.Equal(t, New(-1, 0), movedBy)
	assert.Equal(t, 0, len(body.moves))
}

func TestBody_Update_move_by_many(t *testing.T) {
	body := NewRigidBody(New(10, 10), 0)

	body.Move(left, left, right, up, down)
	hasMoved, movedBy := body.Update()

	assert.True(t, hasMoved)
	assert.Equal(t, New(9, 10), body.center)
	assert.Equal(t, New(-1, 0), movedBy)
	assert.Equal(t, 0, len(body.moves))
}

func TestBody_Update_move_by_many_2(t *testing.T) {
	body := NewRigidBody(New(10, 10), 0)

	body.Move(left, left, left, down, down)
	hasMoved, movedBy := body.Update()

	assert.True(t, hasMoved)
	assert.Equal(t, New(7, 12), body.center)
	assert.Equal(t, New(-3, 2), movedBy)
	assert.Equal(t, 0, len(body.moves))
}

func TestBody_Drag(t *testing.T) {

	tests := []struct {
		name                               string
		body                               *Body
		startingAcceleration, by, expected float64
	}{
		{name: "drag decreases", startingAcceleration: 10, by: 6, expected: 4},
		{name: "drag decreases to 0", startingAcceleration: 4, by: 6, expected: 0},
		{name: "drag has no effect", startingAcceleration: 0, by: 6, expected: 0},
		{name: "drag increase", startingAcceleration: -10, by: 6, expected: -4},
		{name: "drag increase to 0", startingAcceleration: -4, by: 6, expected: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewRigidBody(New(0, 0), 0)
			b.acceleration = tt.startingAcceleration
			b.Drag(tt.by)

			assert.Equal(t, tt.expected, b.Acceleration())
		})
	}
}
