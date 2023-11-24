package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kaschula/twod/components"
	"image/color"
	"log"
)

const (
	screenW, screenH = 1200, 1200
)

// Next Add text and information to other games
// Start rectangle SAT2 collision

func main() {
	broadPhaseGame := NewBroadPhaseCollisionGame()
	vectorsRotateToMouse := NewRotateAndMoveGame()
	circleColliding := NewCircleCollision()
	spaceShipGame := NewSpaceShipGame()
	diagonalCollisionPoly := NewPolyGonLineCollision()

	games := components.NewGameContainer([]ebiten.Game{
		diagonalCollisionPoly,
		// todo Add annotations
		broadPhaseGame,
		// todo Add annotations
		vectorsRotateToMouse,
		circleColliding,
		spaceShipGame,
	}, ebiten.KeyLeftBracket, ebiten.KeyRightBracket)

	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Vectors Tutorial")
	if err := ebiten.RunGame(games); err != nil {
		log.Fatal(err)
	}
}

func keyPressAndHold(k ebiten.Key) bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.KeyPressDuration(k) > 0
}

func colourRed() color.Color {
	return color.RGBA{R: 255, G: 0, B: 0, A: 255}
}

func colourPink() color.Color {
	return color.RGBA{R: 255, G: 0, B: 255, A: 255}
}

func colourBlue() color.Color {
	return color.RGBA{R: 0, G: 0, B: 255, A: 255}
}
