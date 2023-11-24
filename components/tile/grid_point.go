package tile

type GridPoint struct {
	X, Y int
}

func NewGridPoint(x int, y int) GridPoint {
	return GridPoint{X: x, Y: y}
}

func (p GridPoint) Sub(other GridPoint) GridPoint {
	return NewGridPoint(p.X-other.X, p.Y-other.Y)
}

func (p GridPoint) Add(other GridPoint) GridPoint {
	return NewGridPoint(p.X+other.X, p.Y+other.Y)
}
