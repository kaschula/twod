package physics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestV_ToString(t *testing.T) {
	assert.Equal(t, "(11.00:12.00)", New(11, 12).ToString())
	assert.Equal(t, "(9.00:5.00)", New(9, 5).ToString())
}

func TestV_Add(t *testing.T) {
	a := New(10, 10)
	b := New(11, 12)
	assert.Equal(t, New(21, 22), a.Add(b))
}

func TestV_Sub(t *testing.T) {
	a := New(21, 22)
	b := New(10, 10)
	assert.Equal(t, New(11, 12), a.Sub(b))
}

func TestV_Scale(t *testing.T) {
	assert.Equal(t, New(20, 20), New(2, 2).Scale(10))
	assert.Equal(t, New(-1, -0), New(1, 0).Scale(-1))
}

func TestV_Invert(t *testing.T) {
	assert.Equal(t, New(-1, 1), New(1, -1).Invert())
}

func TestV_Perp(t *testing.T) {
	assert.Equal(t, New(10, -1), New(1, 10).Perp())
}

func TestV_Magnitude(t *testing.T) {
	assert.Equal(t, float64(5), New(4, 3).Magnitude())
}

func TestV_Normalize(t *testing.T) {
	assert.Equal(t, New(0.8, 0.6).ToString(), New(4, 3).Normalize().ToString())
}

func TestV_Rotate(t *testing.T) {
	a := New(10, 0).Rotate(New(0, 0), 90)
	assert.Equal(t, New(0, 10).ToString(), a.ToString())
}
