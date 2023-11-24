package physics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircle_Collision_on_same_y(t *testing.T) {
	circle1 := NewCircle(New(200, 200), 30, 0)
	circle2 := NewCircle(New(250, 200), 30, 0)

	collision := circle1.Collision(circle2)

	assert.Equal(t, 10.0, collision.Depth())
	assert.Equal(t, New(230, 200), collision.Start())
	assert.Equal(t, New(220, 200), collision.End())
	assert.Equal(t, New(-1, 0), collision.Normal())
}

func TestCircle_Collision_on_different_y(t *testing.T) {
	circle1 := NewCircle(New(60, 60), 30, 0)
	circle2 := NewCircle(New(80, 20), 30, 0)

	collision := circle1.Collision(circle2)

	assert.Equal(t, 15.278640450004204, collision.Depth())
	assert.Equal(t, New(73.42, 33.17).ToString(), collision.Start().ToString())
	assert.Equal(t, New(66.58, 46.83).ToString(), collision.End().ToString())
	assert.Equal(t, New(-0.45, 0.89).ToString(), collision.Normal().ToString())
}

func Test_Implements_RigidCircleInterface(t *testing.T) {
	var c RigidCircle = NewCircle(New(0, 0), 5, 0)
	assert.NotNil(t, c)
}
