package physics

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRectangle_Faces(t *testing.T) {
	r1 := NewRectangle(New(10, 10), 20, 20, 0)

	assert.Equal(t, []V{{x: 0, y: -1}, {x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}}, r1.Faces())
}

func TestRectangle_Update(t *testing.T) {
	location := New(10, 10)
	r1 := NewRectangle(location, 20, 20, 0)
	r1.Rotate(Degree(45))
	r1.Rotate(Degree(45))
	r1.Update()

	assert.Equal(t, location, r1.Location())
	assert.Equal(t, Degree(90), r1.Direction())
	assert.Equal(t, []V{{x: 20, y: 0}, {x: 20, y: 20}, {x: 0, y: 20}, {x: 0, y: 0}}, r1.Vertices())
	assert.Equal(t, []V{{x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}, {x: 0, y: -1}}, r1.Faces())
}

func TestRectangle_RotateTo(t *testing.T) {
	location := New(10, 10)
	r1 := NewRectangle(location, 20, 20, 0)
	r1.RotateTo(New(20, 10))
	r1.Update()

	assert.Equal(t, location, r1.Location())
	assert.Equal(t, Degree(90), r1.Direction())
	assert.Equal(t, []V{{x: 20, y: 0}, {x: 20, y: 20}, {x: 0, y: 20}, {x: 0, y: 0}}, r1.Vertices())
	assert.Equal(t, []V{{x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}, {x: 0, y: -1}}, r1.Faces())
}

func TestRectangle_Width_Height(t *testing.T) {
	r := NewRectangle(New(0, 0), 20, 30, 0)
	assert.Equal(t, 20, r.Width())
	assert.Equal(t, 30, r.Height())
}

func TestRect_Radiuses(t *testing.T) {

	r := NewRectangle(New(100, 100), 100, 100, 0)

	radiuses := r.Radiuses()

	assert.Equal(t, NewLineSegment(r.Location(), r.TopLeft()), radiuses[0])
	assert.Equal(t, NewLineSegment(r.Location(), r.TopRight()), radiuses[1])
	assert.Equal(t, NewLineSegment(r.Location(), r.BottomRight()), radiuses[2])
	assert.Equal(t, NewLineSegment(r.Location(), r.BottomLeft()), radiuses[3])

}

func TestRect_Edges(t *testing.T) {
	r := NewRectangle(New(100, 100), 100, 100, 0)

	edges := r.Edges()

	assert.Equal(t, NewLineSegment(r.TopLeft(), r.TopRight()), edges[0])
	assert.Equal(t, NewLineSegment(r.TopRight(), r.BottomRight()), edges[1])
	assert.Equal(t, NewLineSegment(r.BottomRight(), r.BottomLeft()), edges[2])
	assert.Equal(t, NewLineSegment(r.BottomLeft(), r.TopLeft()), edges[3])
}

func TestNewRectangleFromCorners(t *testing.T) {
	tl := New(10, 10)
	br := New(15, 16)

	r := NewRectangleFromCorners(tl, br, 0)

	require.Equal(t, 6, r.height)
	require.Equal(t, 5, r.width)
	require.True(t, r.Body.center.Equal(New(12.5, 13)))
}
