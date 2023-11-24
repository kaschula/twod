package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	color2 "github.com/kaschula/twod/color"
	"github.com/kaschula/twod/draw"
	"github.com/kaschula/twod/inputs"
	"github.com/kaschula/twod/physics"
	"log"
)

const (
	screenW, screenH = 1600, 1600
)

func main() {
	game := NewGame()

	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("New Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type TransposesRotates interface {
	Transpose(to physics.V)
	Rotate(byDegrees physics.Degree)
	Update() (bool, physics.V)
}

type Game struct {
	selectedI int
	selected  string
	shapes    []string
	bodies    map[string]TransposesRotates
	drawables map[string]func(screen *ebiten.Image)
}

func (g *Game) Update() error {
	// select a shape to transpose
	if inputs.KeyJustPressed(ebiten.KeyLeft) {
		nextI := g.selectedI - 1

		if nextI < 0 {
			nextI = len(g.shapes) - 1
		}

		g.selectedI = nextI
		g.selected = g.shapes[nextI]
	}

	if inputs.KeyJustPressed(ebiten.KeyRight) {
		nextI := g.selectedI + 1

		if nextI > len(g.shapes)-1 {
			nextI = 0
		}

		g.selectedI = nextI
		g.selected = g.shapes[nextI]
	}

	if inputs.KeyJustPressed(ebiten.KeyUp) {
		for _, shape := range g.bodies {
			shape.Rotate(physics.Degree(10))
		}
	}

	if moveTo, pressed := inputs.MouseLeftPressed(); pressed {
		g.bodies[g.selected].Transpose(moveTo)
	}

	for _, shape := range g.bodies {
		shape.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "press <- and -> toggle selected shape", 10, 10)
	ebitenutil.DebugPrintAt(screen, "press ^ to rotate shapes", 10, 25)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("selected: %v", g.selected), 10, 40)

	for _, f := range g.drawables {
		f(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func NewGame() *Game {
	shapes := []string{"circle", "square", "hex", "triangle"}

	selected := "circle"

	circle := physics.NewCircle(physics.New(300, 300), 100, 0)
	square := physics.NewRectangle(physics.New(600, 1000), 100, 100, 0)
	hex := physics.NewPoly(physics.New(800, 300), 6, 50)
	triangle := physics.NewPoly(physics.New(1000, 1000), 3, 70)

	bodies := map[string]TransposesRotates{
		"circle":   circle,
		"square":   square,
		"triangle": triangle,
		"hex":      hex,
	}

	drawables := map[string]func(screen *ebiten.Image){
		"circle": func(screen *ebiten.Image) {
			draw.Circle(screen, circle, color2.WinterSoftBlue)
		},
		"square": func(screen *ebiten.Image) {
			draw.Rectangle(screen, square, color2.WinterBlue, true)
		},
		"triangle": func(screen *ebiten.Image) {
			draw.Polygon(screen, triangle, color2.WinterBeige)
		},
		"hex": func(screen *ebiten.Image) {
			draw.Polygon(screen, hex, color2.WinterOffWhite)
		},
	}

	return &Game{selectedI: 0, selected: selected, bodies: bodies, shapes: shapes, drawables: drawables}
}
