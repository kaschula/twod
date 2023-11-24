package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kaschula/twod/draw"
	physicsvector "github.com/kaschula/twod/physics"
	"image/color"
)

type BroadPhaseCollisionGame struct {
	count     int
	rectangle physicsvector.RigidPoly
	circle    physicsvector.RigidCircle

	// DrawBounding box
	// Add a bounding box draw mode and draw function draw
}

func NewBroadPhaseCollisionGame() *BroadPhaseCollisionGame {
	return &BroadPhaseCollisionGame{
		rectangle: physicsvector.NewRectangle(physicsvector.New(300, 300), 70, 40, 80),
		circle:    physicsvector.NewCircle(physicsvector.New(500, 300), 30, 40),
	}
}

func (g *BroadPhaseCollisionGame) Update() error {
	g.count++

	if inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.KeyPressDuration(ebiten.KeyA) > 0 {
		fmt.Printf("key a move left \n")
		g.rectangle.Move(physicsvector.New(-10, 0))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.KeyPressDuration(ebiten.KeyD) > 0 {
		fmt.Printf("key d move right \n")
		g.rectangle.Move(physicsvector.New(10, 0))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		fmt.Printf("key right \n")
		g.rectangle.Rotate(10)
		g.circle.Rotate(10)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		fmt.Printf("key right \n")
		g.rectangle.Rotate(-10)
		g.circle.Rotate(-10)

	}

	g.rectangle.Update()
	//g.circle.Update()

	collision := g.circle.BoundingCollidesWith(g.rectangle)
	if collision {
		fmt.Printf("objects are colliding \n")
	}

	return nil
}

func (g *BroadPhaseCollisionGame) Draw(screen *ebiten.Image) {
	g.drawRigidBodies(screen)
}

func (g *BroadPhaseCollisionGame) drawRigidBodies(screen *ebiten.Image) {
	// todo add text label to explain demo and the keys

	collisionColor := color.RGBA{1, 230, 230, 230}

	var (
		circleColor color.Color = color.RGBA{1, 186, 239, 255}
		rectColor   color.Color = color.White
	)

	if g.circle.IsColliding() {
		circleColor = collisionColor
	}

	if g.rectangle.IsColliding() {
		rectColor = collisionColor
	}

	// bounding colliding shapes
	draw.Circle(screen, g.circle, circleColor)
	draw.Polygon(screen, g.rectangle, rectColor)

	draw.PolygonFaces(screen, g.rectangle, rectColor)

}

func (g *BroadPhaseCollisionGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}
