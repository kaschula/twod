package physics

import (
	"fmt"
	"math"
)

// TODO: Refactor and orginisation
//1. create a test file in each example, this will allow for compilation checks
//2. refactor vector naming
//3. add the word Collision infront of all collision functions
//4. continue work on the self driving car

var (
	ZeroVector = New(0, 0)
)

// V Todo change to D2 (2 dimensional)
type V struct {
	x, y float64
}

func New(x, y float64) V {
	return V{x: x, y: y}
}

func (v V) ToString() string {
	return fmt.Sprintf("(%.2f:%.2f)", v.x, v.y)
}

func (v V) ToFloat64() (float64, float64) {
	return v.x, v.y
}

func (v V) ToInt() (int, int) {
	return int(v.x), int(v.y)
}

func (v V) WithToFloat64(v2 V) (float64, float64, float64, float64) {
	return v.x, v.y, v2.x, v2.y
}

func (v V) X() float64 {
	return v.x
}

func (v V) Y() float64 {
	return v.y
}

func (v V) X32() float32 {
	return float32(v.x)
}

func (v V) Y32() float32 {
	return float32(v.y)
}

func (v V) XInt() int {
	return int(v.x)
}

func (v V) YInt() int {
	return int(v.y)
}

func (v V) Direction() float64 {
	return v.x
}

func (v V) Add(v2 V) V {
	return New(v.x+v2.x, v.y+v2.y)
}

func (v V) Sub(v2 V) V {
	return New(v.x-v2.x, v.y-v2.y)
}

func (v V) Multiply(v2 V) V {
	return New(v.x*v2.x, v.y*v2.y)
}

func (v V) Scale(r float64) V {
	return New(v.x*r, v.y*r)
}

func (v V) Equal(v2 V) bool {
	return v.x == v2.x && v.y == v2.y
}

func (v V) Dot(v2 V) float64 {
	return v.x*v2.x + v.y*v2.y
}

// Magnitude or Length
func (v V) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v V) Length() float64 {
	return v.Magnitude()
}

// Invert Rotates the vector 180 degrees
func (v V) Invert() V {
	return New(-v.x, -v.y)
}

// Perp Rotates the vector 90 degrees anti clock wise
func (v V) Perp() V {
	return New(v.y, -v.x)
}

func (v V) Normalize() V {
	length := v.Magnitude()
	if length > 0 {
		length = 1.0 / length
	}

	return v.Scale(length)
}

func (v V) Rotate(center V, degrees Degree) V {
	if degrees == 0 {
		return v
	}

	return v.RotateRadian(center, degrees.ToRadian())
}

func (v V) RotateRadian(center V, radians Radian) V {
	if radians == 0 {
		return v
	}
	//clockwise
	radians64 := radians.ToFloat64()
	cos := math.Cos(radians64)
	sin := math.Sin(radians64)

	// offset vector around center
	offSetX := v.x - center.x
	offSetY := v.y - center.y

	// apply angles
	nx := (cos * (offSetX)) + (sin * (offSetY))
	ny := (cos * (offSetY)) - (sin * (offSetX))

	// transpose back original position
	nx += center.x
	ny += center.y

	return New(nx, ny)
}

func (v V) OutOfBounds(width, height int) bool {
	x, y := v.ToFloat64()

	if x < 0 || x > float64(width) {
		return true
	}

	if y < 0 || y > float64(height) {
		return true
	}

	return false
}

//Vec2.prototype.cross = function (vec) {
//return (this.x * vec.y - this.y * vec.x);
//};

//Vec2.prototype.distance = function (vec) {
//var x = this.x - vec.x;
//var y = this.y - vec.y;
//return Math.sqrt(x * x + y * y);
//};

type Matrix2d []V

func NewMatrix2d(vs ...V) Matrix2d {
	return vs
}

func (m Matrix2d) Map(fn func(v V) V) Matrix2d {
	newM := make(Matrix2d, len(m))
	for i := range m {
		newM[i] = fn(m[i])
	}

	return newM
}

func (m Matrix2d) Rotate(center V, degree Degree) Matrix2d {
	return m.Map(func(v V) V {
		return v.Rotate(center, degree)
	})
}

func (m Matrix2d) Add(by V) Matrix2d {
	return m.Map(func(v V) V {
		return v.Add(by)
	})
}

func (m Matrix2d) Sum() V {
	sum := New(0, 0)
	for _, v := range m {
		sum = sum.Add(v)
	}

	return sum
}
func (m Matrix2d) Clear() Matrix2d {
	return Matrix2d{}
}
