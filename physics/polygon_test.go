package physics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_setZeroCornerPoints(t *testing.T) {
	// square
	expected := NewMatrix2d(New(10, 0), New(20, 10), New(10.000000000000002, 20), New(0, 10.000000000000002))
	assert.Equal(t, expected, setZeroCornerPoints(4, 10, New(10, 10)))

	// triangle

	// hexigon
}

func Test_faces_match_rectangle(t *testing.T) {
	// square
	square := NewPoly(New(10, 10), 4, 20)
	rectangle := NewRectangle(New(10, 10), 20, 20, 0, WithAngle(Degree(45)))
	rectangle.Update()

	assert.Equal(t, rectangle.Faces()[0].ToString(), square.Faces()[0].ToString())
	assert.Equal(t, rectangle.Faces()[1].ToString(), square.Faces()[1].ToString())
	assert.Equal(t, rectangle.Faces()[2].ToString(), square.Faces()[2].ToString())
	assert.Equal(t, rectangle.Faces()[3].ToString(), square.Faces()[3].ToString())
}
