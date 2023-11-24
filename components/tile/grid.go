package tile

import (
	"github.com/google/uuid"
	"github.com/kaschula/twod/physics"
	"github.com/kaschula/twod/physics/collision"
)

func NewTileBoard(center physics.V, tileW, tileH, numTileW, numTileH int) *TileGrid {

	inBoardCenter := physics.New(float64(tileW*numTileW/2), float64(tileH*numTileH/2))
	topLeftOfBoard := center.Sub(inBoardCenter)
	centerFirstTile := topLeftOfBoard.Add(physics.New(float64(tileW/2), float64(tileH/2)))

	//fmt.Println(centerFirstTile)

	tiles := []*Tile{}
	tilesByID := map[string]*Tile{}

	for wI := 0; wI < numTileW; wI++ {
		for hI := 0; hI < numTileH; hI++ {
			centerTile := centerFirstTile.Add(physics.New(float64(wI*tileW), float64(hI*tileH)))

			t := NewTile(uuid.NewString(), wI, hI, centerTile, tileW, tileH)
			tiles = append(tiles, t)
			tilesByID[t.ID] = t
		}
	}
	return &TileGrid{
		tileW: tileW, tileH: tileH, numTileW: numTileW, numTileH: numTileH,
		Tiles: tiles,
	}
}

type TileGrid struct {
	tileW, tileH, numTileW, numTileH int
	Tiles                            []*Tile
	TilesByID                        map[string]*Tile
}

func (g *TileGrid) GetTileBYGridReference(x, y int) (*Tile, bool) {
	for _, tile := range g.Tiles {
		if tile.x == x && tile.y == y {
			return tile, true
		}
	}

	return nil, false
}

func (g *TileGrid) GetTileBYGridPoint(p GridPoint) (*Tile, bool) {
	return g.GetTileBYGridReference(p.X, p.Y)
}

func (g *TileGrid) Foreach(fn func(i int, t *Tile)) {
	for i, tile := range g.Tiles {
		fn(i, tile)
	}
}

// PointInTile returns nil if notclicked
func (g *TileGrid) PointInTile(point physics.V) (*Tile, bool) {
	//todo optimise
	for _, tile := range g.Tiles {

		rect := tile.RigidBody()
		collision := collision.PointInRectangle(rect, point)

		if collision {
			return tile, collision
		}

	}

	return nil, false
}

type Tile struct {
	ID string
	// x, y board coordinates, how tiles across and down a tile is, x and y will never more than the number of tiles in the board width or hight
	x, y int
	poly physics.Rectangle
}

func (t *Tile) RigidBody() physics.Rectangle {
	return t.poly
}

func (t *Tile) GridPoint() GridPoint {
	return NewGridPoint(t.x, t.y)
}

func (t *Tile) X() int {
	return t.x
}
func (t *Tile) Y() int {
	return t.y
}

func NewTile(ID string, x, y int, center physics.V, width, height int) *Tile {
	rect := physics.NewRectangle(center, width, height, 0, physics.WithSpeed(0))
	return &Tile{ID: ID, poly: rect, x: x, y: y}
}
