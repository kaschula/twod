package components

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kaschula/twod/slice"
	"image/color"
)

// Todo untested
// is this need as a tile board exists
// Grid will probably need a set colour but add later
type Grid struct {
	tileW, tileH int
	gridW, gridH int
	color        color.Color
}

func NewGrid(tileW, tileH, gridW, gridH int, c color.Color) Grid {

	if tileW == 0 || tileH == 0 || gridW == 0 || gridH == 0 {
		panic(fmt.Sprintf("%v %v %v %v none of these can be zero", tileW, tileH, gridW, gridH))
	}

	return Grid{tileW: tileW, tileH: tileH, gridW: gridW, gridH: gridH, color: c}
}

// todo this doesn;t need to be a method and should be draw function, the colour property is not needed
func (g Grid) Draw(screen *ebiten.Image) {
	width, height, tileWidth, tileHeight := g.gridW, g.gridH, g.tileW, g.tileH

	// code this up for flexibility
	verticalLineMax := float64(tileHeight * height)
	horizontalLineMax := float64(tileHeight * width)

	// use 0 in the range min if you want an out right line
	// vertical
	for i := range slice.Range(0, width) {
		x := float64(i * tileWidth)

		ebitenutil.DrawLine(screen, x, 0, x, verticalLineMax, g.color)
	}
	// horizontal
	for i := range slice.Range(0, height) {
		y := float64(i * tileHeight)
		ebitenutil.DrawLine(screen, 0, y, horizontalLineMax, y, g.color)
	}
}
