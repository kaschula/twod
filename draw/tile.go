package draw

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	color2 "github.com/kaschula/twod/color"
	"github.com/kaschula/twod/components/tile"
	"image/color"
)

// TiledGrid used for debug of tiles
func TiledGrid(screen *ebiten.Image, board *tile.TileGrid, cs ...color.Color) {
	var tileColor color.Color = color2.White
	var lineColor color.Color = color2.Black

	fmt.Println(tileColor, lineColor)

	if len(cs) >= 1 {
		tileColor = cs[0]
	}

	if len(cs) >= 2 {
		lineColor = cs[1]
	}

	for _, t := range board.Tiles {
		Rectangle(screen, t.RigidBody(), tileColor, false)
		WireFrame(screen, t.RigidBody(), lineColor)
	}
}

func TiledGridNoFill(screen *ebiten.Image, board *tile.TileGrid, cs ...color.Color) {
	var lineColor color.Color = color2.Black

	for _, t := range board.Tiles {
		WireFrame(screen, t.RigidBody(), lineColor)
	}
}

// TiledGridCoordinates used for debug of tiles
func TiledGridCoordinates(screen *ebiten.Image, board *tile.TileGrid, cs ...color.Color) {
	for _, t := range board.Tiles {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%v:%v", t.X(), t.Y()), t.RigidBody().Location().XInt(), t.RigidBody().Location().YInt())
	}
}
