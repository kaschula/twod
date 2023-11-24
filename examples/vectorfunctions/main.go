package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	color2 "github.com/kaschula/twod/color"
	"github.com/kaschula/twod/draw"
	"github.com/kaschula/twod/physics"
	"log"
)

const (
	screenW, screenH = 600, 600
)

var (
	startingPosition = physics.New(200, 200)
)

func main() {
	game := NewGame(startingPosition)

	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Vector functions")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	userPosition, direction physics.V
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.direction = g.direction.Invert()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.direction = g.direction.Perp()
	}

	g.userPosition = g.userPosition.Add(g.direction)

	if g.userPosition.OutOfBounds(screenW, screenH) {
		g.userPosition = startingPosition
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Vector Functions", 10, 10)
	ebitenutil.DebugPrintAt(screen, "A = Invert, S = Perp", 10, 20)

	draw.Rectangle(screen, physics.NewRectangle(g.userPosition, 10, 10, 0), color2.Green, false)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func NewGame(userPosition physics.V) *Game {
	return &Game{
		direction:    physics.New(3, 3),
		userPosition: userPosition,
	}
}
