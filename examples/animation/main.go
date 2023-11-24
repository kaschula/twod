package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kaschula/twod/components/animation"
	"github.com/kaschula/twod/components/gameobject"
	"github.com/kaschula/twod/inputs"
	"github.com/kaschula/twod/physics"
	"image"
	_ "image/png"
	"log"
)

const (
	screenW, screenH = 500, 240
)

//go:embed runner.png
var spriteSheet []byte

func main() {
	game := NewAnimationExample()

	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Animation")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type AnimationExample struct {
	//animation1 *animation.Animation
	//animation2 *animation.Animation
	player *gameobject.GameObject
}

func (a AnimationExample) Update() error {
	if inputs.KeyPressAndHold(ebiten.KeyR) {
		//a.player.Transpose()
	}

	a.player.Update()

	return nil
}

func (a AnimationExample) Draw(screen *ebiten.Image) {
	a.player.Draw(screen)

	ebitenutil.DebugPrintAt(screen, "Press <- and -> keys to move player", 10, 10)
}

func (a AnimationExample) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func NewAnimationExample() *AnimationExample {
	img, _, err := image.Decode(bytes.NewReader(spriteSheet))
	if err != nil {
		panic(err)
	}

	spriteSheetImg := ebiten.NewImageFromImage(img)

	frameWidth, frameHeight := 32, 32

	// first row of sprites
	idleAnimation := animation.New(
		6, 32, 32,
		// todo wrap the ebiten.image package in your own image package that works with physics.V
		// coordinates here are for the vector of the top left and bottom right of a none rotated rectangle
		spriteSheetImg.SubImage(image.Rect(0+(frameWidth*0), 0, frameWidth+(frameWidth*0), frameHeight)).(*ebiten.Image),
		spriteSheetImg.SubImage(image.Rect(0+(frameWidth*1), 0, frameWidth+(frameWidth*1), frameHeight)).(*ebiten.Image),
		spriteSheetImg.SubImage(image.Rect(0+(frameWidth*2), 0, frameWidth+(frameWidth*2), frameHeight)).(*ebiten.Image),
		spriteSheetImg.SubImage(image.Rect(0+(frameWidth*3), 0, frameWidth+(frameWidth*3), frameHeight)).(*ebiten.Image),
		spriteSheetImg.SubImage(image.Rect(0+(frameWidth*4), 0, frameWidth+(frameWidth*4), frameHeight)).(*ebiten.Image),
	)

	// second row
	runningAnimationRight := animation.New(
		5, 32, 32,
		spriteSheetImg.SubImage(image.Rect(0+(frameWidth*0), frameWidth, frameWidth+(frameWidth*0), frameHeight*2)).(*ebiten.Image),
		spriteSheetImg.SubImage(image.Rect(0+(frameWidth*1), frameWidth, frameWidth+(frameWidth*1), frameHeight*2)).(*ebiten.Image),
		spriteSheetImg.SubImage(image.Rect(0+(frameWidth*2), frameWidth, frameWidth+(frameWidth*2), frameHeight*2)).(*ebiten.Image),
		spriteSheetImg.SubImage(image.Rect(0+(frameWidth*3), frameWidth, frameWidth+(frameWidth*3), frameHeight*2)).(*ebiten.Image),
		spriteSheetImg.SubImage(image.Rect(0+(frameWidth*4), frameWidth, frameWidth+(frameWidth*4), frameHeight*2)).(*ebiten.Image),
		spriteSheetImg.SubImage(image.Rect(0+(frameWidth*5), frameWidth, frameWidth+(frameWidth*5), frameHeight*2)).(*ebiten.Image),
		spriteSheetImg.SubImage(image.Rect(0+(frameWidth*6), frameWidth, frameWidth+(frameWidth*6), frameHeight*2)).(*ebiten.Image),
		spriteSheetImg.SubImage(image.Rect(0+(frameWidth*7), frameWidth, frameWidth+(frameWidth*7), frameHeight*2)).(*ebiten.Image),
	)

	playerBody := physics.NewRectangle(
		physics.New(0, float64(screenH/2)),
		32, 32, 0,
	)

	player := gameobject.NewGameObject(
		[]string{"IDLE", "RUNNING_RIGHT"},
		map[string]animation.SpriteRenderer{"IDLE": idleAnimation, "RUNNING_RIGHT": runningAnimationRight},
		func(gameObject *gameobject.GameObject) {
			if inputs.KeyPressAndHold(ebiten.KeyRight) {

				fmt.Println("running")
				gameObject.SwitchState("RUNNING_RIGHT")

				gameObject.Body().Move(physics.New(2, 0))

			} else if inputs.KeyPressAndHold(ebiten.KeyLeft) {

				fmt.Println("running")
				gameObject.SwitchState("RUNNING_RIGHT")

				gameObject.Body().Move(physics.New(-2, 0))

			} else {
				fmt.Println("idle")
				gameObject.SwitchState("IDLE")
			}

			gameObject.Body().Update()
		},
		gameobject.WithRigidPoly(playerBody),
	)

	return &AnimationExample{
		player: player,
	}
}
