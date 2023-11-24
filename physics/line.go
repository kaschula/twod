package physics

import "C"
import (
	"fmt"
	"github.com/kaschula/twod/math"
)

type LineSegment struct {
	start, end V
}

func NewLineSegment(start V, end V) *LineSegment {
	return &LineSegment{start: start, end: end}
}

func (l *LineSegment) ID() string {
	return fmt.Sprintf("%s:%s", l.start.ToString(), l.end.ToString())
}

// GetSharedVector return true and a vector if any to points at the end of either line match
func (l *LineSegment) GetSharedVector(l2 *LineSegment) (bool, V) {
	if l.start == l2.start {
		return true, l.start
	}

	if l.end == l2.start {
		return true, l.end
	}

	if l.start == l2.end {
		return true, l.start
	}

	if l.end == l2.end {
		return true, l.end
	}

	return false, ZeroVector
}

func (l *LineSegment) Start() V {
	return l.start
}

func (l *LineSegment) End() V {
	return l.end
}

func (l *LineSegment) Direction() V {
	return l.start.Sub(l.end).Normalize()
}

func (l *LineSegment) Magnitude() float64 {
	return l.start.Sub(l.end).Magnitude()
}

func (l *LineSegment) NearestEnd(to V) V {
	distanceFromStart := l.Start().Sub(to).Length()
	distanceFromEnd := l.End().Sub(to).Length()

	if distanceFromStart <= distanceFromEnd {
		return l.start
	}

	return l.end
}
func (l *LineSegment) MoveStart(by V) *LineSegment {
	return &LineSegment{
		start: l.start.Add(by),
		end:   l.end,
	}
}

func (l *LineSegment) MoveEnd(by V) *LineSegment {
	return &LineSegment{
		start: l.start,
		end:   l.end.Add(by),
	}
}

func (l *LineSegment) Transpose(by V) *LineSegment {
	return &LineSegment{
		start: l.start.Add(by),
		end:   l.end.Add(by),
	}
}

func GetLineIntersection(l1, l2 *LineSegment) Touch {
	return GetIntersection(l1.Start(), l1.End(), l2.Start(), l2.End())
}

// todo write tests for this
func GetIntersection(a, b, c, d V) Touch {
	tTop := (d.x-c.x)*(a.y-c.y) - (d.y-c.y)*(a.x-c.x)
	uTop := (c.y-a.y)*(a.x-b.x) - (c.x-a.x)*(a.y-b.y)
	bottom := (d.y-c.y)*(b.x-a.x) - (d.x-c.x)*(b.y-a.y)

	if bottom != 0 {
		t := tTop / bottom
		u := uTop / bottom

		if t >= 0 && t <= 1 && u >= 0 && u <= 1 {
			return Touch{
				X:      math.Lerp(a.x, b.x, t),
				Y:      math.Lerp(a.y, b.y, t),
				LineA:  NewLineSegment(a, b),
				LineB:  NewLineSegment(c, d),
				Offset: t,
			}
		}
	}

	// no touch
	return Touch{}
}

var (
	EmptyTouch = Touch{
		X:      0,
		Y:      0,
		Offset: 0,
		LineA:  nil,
		LineB:  nil,
	}
)

// Touch represents the interesction between 2 lines
// Offset represents how car the collision is from Line A's start Vector
// todo convert this to a collision
type Touch struct {
	X, Y, Offset float64
	// LineA is the line being tested against
	LineA *LineSegment
	// LineB is an actor line. Has LineA comein to contact with LineB
	LineB *LineSegment
}

func (t Touch) Empty() bool {
	return t.X == 0 && t.Y == 0 && t.Offset == 0 && t.LineA == nil && t.LineB == nil
}

func (t Touch) Vector() V {
	return New(t.X, t.Y)
}

// LineAStartToOffSet returns a line segement from the start of line A to the touch point
func (t Touch) LineAStartToOffSet() *LineSegment {
	return NewLineSegment(t.LineA.Start(), t.Vector())
}

// LineAEndToOffSet returns a line segement from the touch vector to line A end
func (t Touch) LineAEndToOffSet() *LineSegment {
	return NewLineSegment(t.Vector(), t.LineA.End())
}

func (t Touch) Transpose(by V) {
	current := t.Vector()
	moved := current.Add(by)

	t.X = moved.X()
	t.Y = moved.Y()

	if t.LineA != nil {
		t.LineA = NewLineSegment(t.LineA.Start().Add(by), t.LineA.End().Add(by))
	}

	if t.LineB != nil {
		t.LineB = NewLineSegment(t.LineB.Start().Add(by), t.LineB.End().Add(by))
	}
}
