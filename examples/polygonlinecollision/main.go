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
	rectangleLocation = physics.New(300, 300)
)

func main() {
	game := NewGame()

	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Vertices")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	showFullRectangle bool
	userLocation      physics.V
}

func (g *Game) Update() error {
	//return errors.New("implement update")
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.showFullRectangle = !g.showFullRectangle
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.userLocation = g.userLocation.Add(physics.New(5, 0))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.userLocation = g.userLocation.Add(physics.New(-5, 0))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.userLocation = g.userLocation.Add(physics.New(0, 5))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.userLocation = g.userLocation.Add(physics.New(0, -5))
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "POLYGON COLLISION WITH LINE INTERSECTION", 10, 10)
	ebitenutil.DebugPrintAt(screen, "Press space to show rectangle", 10, 25)

	r := physics.NewRectangle(rectangleLocation, 100, 150, 0)
	r.Rotate(physics.Degree(45))
	r.Update()

	user := physics.NewRectangle(g.userLocation, 100, 150, 0)
	user.Rotate(physics.Degree(70))
	user.Update()

	g.drawRect(screen, r)
	g.drawRect(screen, user)

	collision := physics.CollisionRigidPolyDiagonalLine(user, r)
	if collision == nil {
		return
	}

	draw.Vector(screen, collision.Start(), color2.Purple, 2)
	draw.Vector(screen, collision.End(), color2.Red, 2)

	resolved := physics.NewRectangle(user.Location().Add(collision.Resolve()), user.Width(), user.Height(), 0, physics.WithAngle(user.Direction()))
	resolved.Update()

	draw.Rectangle(screen, resolved, color2.RGBA(0, 255, 0, 50), false)

}

func (g *Game) drawRect(screen *ebiten.Image, r physics.Rectangle) {
	if g.showFullRectangle {
		draw.Rectangle(screen, r, color2.Blue, false)
	}

	for _, v := range r.Vertices() {
		draw.Vector(screen, v, color2.Green, 4)
	}

	for _, line := range r.Edges() {
		draw.Line(screen, line.Start(), line.End(), color2.Green)
	}

	for _, line := range r.Radiuses() {
		draw.Line(screen, line.Start(), line.End(), color2.Green)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func NewGame() *Game {
	return &Game{
		userLocation: physics.New(200, 200),
	}
}
