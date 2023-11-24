package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kaschula/twod/draw"
	"github.com/kaschula/twod/physics"
	"image/color"
)

type circleCollision struct {
	circle1, circle2     physics.RigidCircle
	collision            *physics.Collision
	resolveCollisionMode bool
}

func NewCircleCollision() *circleCollision {
	return &circleCollision{
		circle1:              physics.NewCircle(physics.New(600, 600), 100, 0),
		circle2:              physics.NewCircle(physics.New(600, 600), 100, 0),
		collision:            nil,
		resolveCollisionMode: false,
	}
}

func (c *circleCollision) Update() error {
	if keyPressAndHold(ebiten.KeyA) {
		c.circle1.Move(physics.New(-10, 0))
	}

	if keyPressAndHold(ebiten.KeyW) {
		c.circle1.Move(physics.New(0, -10))
	}

	if keyPressAndHold(ebiten.KeyD) {
		c.circle1.Move(physics.New(10, 0))
	}

	if keyPressAndHold(ebiten.KeyS) {
		c.circle1.Move(physics.New(0, 10))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		c.resolveCollisionMode = !c.resolveCollisionMode
	}

	c.collision = c.circle1.Collision(c.circle2)

	if c.collision != nil && c.resolveCollisionMode {
		coll := c.collision
		coll.Normal().Scale(coll.Depth())

		c.circle1.Move(coll.Normal().Scale(coll.Depth()))
	}

	c.circle1.Update()
	c.circle2.Update()
	return nil
}

func (c *circleCollision) Draw(screen *ebiten.Image) {
	isCollisionNotInResolvedMode := c.collision != nil && !c.resolveCollisionMode

	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	circle1, circle2 := c.circle1, c.circle2
	ebitenutil.DebugPrintAt(screen, "CIRCLE CIRCLE COLLISION TEST", 10, 10)

	collidingColor := color.RGBA{R: 230, G: 230, B: 255, A: 255}
	circle1Color := color.RGBA{R: 200, G: 230, B: 30, A: 255}
	circle2Color := color.RGBA{R: 60, G: 100, B: 230, A: 255}

	if isCollisionNotInResolvedMode {
		circle1Color = collidingColor
		circle2Color = collidingColor
	}

	draw.Circle(screen, circle1, circle1Color)
	draw.Circle(screen, circle2, circle2Color)

	ebitenutil.DrawLine(screen, circle1.Location().X(), circle1.Location().Y(), circle2.Location().X(), circle2.Location().Y(), color.RGBA{R: 230, G: 255, B: 255, A: 255})

	// data
	printText := fmt.Sprintf(`
resolve  : %v
distance : %.2f,
radiusSum: %v`, Ternary(c.resolveCollisionMode, "On", "Off"), circle1.Location().Sub(circle2.Location()).Length(), circle1.Radius()+circle2.Radius())
	ebitenutil.DebugPrintAt(screen, printText, 10, 30)

	// controls
	inputText := `Controls:
AWSD to move circle 1
spacebar: toggle resolve mode`
	ebitenutil.DebugPrintAt(screen, inputText, screenW-200, 10)

	// print collision data
	if isCollisionNotInResolvedMode {
		collision := c.collision
		collisionStart := physics.NewCircle(c.collision.Start(), 5, 0)
		collisionEnd := physics.NewCircle(c.collision.End(), 5, 0)

		draw.Circle(screen, collisionStart, color.Black)
		draw.Circle(screen, collisionEnd, color.Black)

		ebitenutil.DrawLine(screen, c.collision.Start().X(), c.collision.Start().Y(), c.collision.End().X(), c.collision.End().Y(), red)

		collisionPrint := fmt.Sprintf(`Collision
Depth : %.2f,
Start : %v,
End   : %v,
Normal: %v,
`, collision.Depth(), collision.Start().ToString(), collision.End().ToString(), collision.Normal().ToString())

		ebitenutil.DebugPrintAt(screen, collisionPrint, 10, 100)
	}
}

func (c *circleCollision) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}

	return falseValue
}
