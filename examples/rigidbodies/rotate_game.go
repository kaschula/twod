package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kaschula/twod/draw"
	physicsvector "github.com/kaschula/twod/physics"
	"image/color"
)

type RotateAndMoveGame struct {
	count int

	// mouse rotate
	// todo add shooting projectiles when space pressed for both these objects
	mouseRotateCircle    *physicsvector.Circle
	mouseRotateRectangle physicsvector.RigidPoly

	// moving to mouse
	movingCircle *physicsvector.Circle
}

func NewRotateAndMoveGame() *RotateAndMoveGame {
	return &RotateAndMoveGame{
		mouseRotateCircle:    physicsvector.NewCircle(physicsvector.New(100, 500), 30, 40),
		mouseRotateRectangle: physicsvector.NewRectangle(physicsvector.New(400, 500), 40, 40, 80),
		movingCircle: physicsvector.NewCircle(
			physicsvector.New(100, 800),
			30, 40,
			physicsvector.WithAngle(90),
			physicsvector.WithSpeed(2),
			physicsvector.WithAcceleration(1),
		),
	}
}

func (g *RotateAndMoveGame) Update() error {
	g.count++

	g.updateMouseRotationObjects()
	g.updateCircleThatMovesTowardsMouse()

	return nil
}

func (g *RotateAndMoveGame) updateCircleThatMovesTowardsMouse() {
	mouseX, mouseY := ebiten.CursorPosition()
	mouseV := physicsvector.New(float64(mouseX), float64(mouseY))

	g.movingCircle.RotateTo(mouseV)
	g.movingCircle.Update()
}

func (g *RotateAndMoveGame) updateMouseRotationObjects() {
	mouseX, mouseY := ebiten.CursorPosition()
	mouseV := physicsvector.New(float64(mouseX), float64(mouseY))

	g.mouseRotateCircle.RotateTo(mouseV)
	g.mouseRotateRectangle.RotateTo(mouseV)

	g.mouseRotateCircle.Update()
	g.mouseRotateRectangle.Update()
}

func (g *RotateAndMoveGame) Draw(screen *ebiten.Image) {
	g.drawRigidBodies(screen)
}

func (g *RotateAndMoveGame) drawRigidBodies(screen *ebiten.Image) {
	// mouse rotating shapes based on mouse
	grey := color.RGBA{200, 200, 200, 255}
	draw.Circle(screen, g.mouseRotateCircle, grey)
	draw.Polygon(screen, g.mouseRotateRectangle, grey)

	// moving circle to mouse
	blue := color.RGBA{150, 50, 255, 255}
	draw.Circle(screen, g.movingCircle, blue)
}

func (g *RotateAndMoveGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}
