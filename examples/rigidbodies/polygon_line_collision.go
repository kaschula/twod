package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kaschula/twod/color"
	"github.com/kaschula/twod/draw"
	"github.com/kaschula/twod/physics"
)

type PolyGonLineCollision struct {
	// todo update this be interface
	user, obstacle1, obstacle2 physics.RigidPoly
}

func NewPolyGonLineCollision() *PolyGonLineCollision {

	return &PolyGonLineCollision{
		user:      physics.NewPoly(physics.New(screenW/2, screenH/2), 3, 50),
		obstacle1: physics.NewPoly(physics.New(200, 200), 5, 120),
		obstacle2: physics.NewPoly(physics.New(800, 1000), 6, 120),
	}
}

func (g *PolyGonLineCollision) Update() error {
	// movement
	if keyPressAndHold(ebiten.KeyA) {
		g.user.Move(physics.New(-10, 0))
	}

	if keyPressAndHold(ebiten.KeyW) {
		g.user.Move(physics.New(0, -10))
	}

	if keyPressAndHold(ebiten.KeyD) {
		g.user.Move(physics.New(10, 0))
	}

	if keyPressAndHold(ebiten.KeyS) {
		g.user.Move(physics.New(0, 10))
	}

	// rotating
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		fmt.Printf("key right \n")
		g.user.Rotate(10)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		fmt.Printf("key right \n")
		g.user.Rotate(-10)
	}

	g.user.Update()

	return nil
}

func (g *PolyGonLineCollision) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "DIAGONAL COLLISION TEST", 10, 10)
	ebitenutil.DebugPrintAt(screen, "Transpose with WASD, rotate with <- -> ", 10, 25)

	// DrawUnit all polys

	draw.Polygon(screen, g.obstacle1, color.WinterBlue)
	draw.Polygon(screen, g.obstacle2, color.WinterSoftBlue)
	draw.Polygon(screen, g.user, color.WinterBeige)

	var collision *physics.Collision
	for _, obstacle := range []physics.RigidPoly{g.obstacle1, g.obstacle2} {
		c := physics.CollisionRigidPolyDiagonalLine(g.user, obstacle)
		if c != nil {
			collision = c
		}
	}

	if collision == nil {
		return
	}

	resolved := physics.NewPoly(g.user.Location().Add(collision.Resolve()), g.user.Sides(), g.user.Width(), physics.WithAngle(g.user.Direction()))
	resolved.Update()

	draw.Polygon(screen, resolved, color.White)

}

func (g *PolyGonLineCollision) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenW
}
