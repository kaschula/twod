package animation

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kaschula/twod/physics"
)

// think of a better name
type SpriteRenderer interface {
	Update()
	Draw(center physics.V, screen *ebiten.Image)
}

type staticSprite struct {
	img                     *ebiten.Image
	frameWidth, frameHeight int
}

func NewStaticImage(img *ebiten.Image, frameWidth int, frameHeight int) *staticSprite {
	return &staticSprite{img: img, frameWidth: frameWidth, frameHeight: frameHeight}
}

func (s *staticSprite) Update() {}

func (s *staticSprite) Draw(center physics.V, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	// This first translation off sets the image to the center minus the frame size
	op.GeoM.Translate(-float64(s.frameWidth)/2, -float64(s.frameHeight)/2)
	op.GeoM.Translate(center.X(), center.Y())
	screen.DrawImage(s.img, op)
}

// todo create an Animation that works time
type Animation struct {
	// frame is calculated on update
	//frame                                   int
	images                                  []*ebiten.Image
	numberOfFrames, frameWidth, frameHeight int
	// controls the speed, the higher the number the slower the animation / the longer each frame is drawn
	frameLength int
	// animationSequence the value in the slice is imageI index of the sequence to draw
	animationSequence             []int
	animationSequenceLastI        int
	currentAnimationSequenceFrame int
	isFirst                       bool
}

func New(frameLength, frameWidth, frameHeight int, images ...*ebiten.Image) *Animation {

	numberOfFrames := len(images)

	animationSequence := []int{}

	// this creates a linear sequence.
	// todo add a way so that the frameImageI can produced using tweening
	for frameImageI := 0; frameImageI < numberOfFrames; frameImageI++ {
		for i := 0; i < frameLength; i++ {
			animationSequence = append(animationSequence, frameImageI)
		}
	}

	return &Animation{
		images:                        images,
		numberOfFrames:                numberOfFrames,
		frameLength:                   frameLength,
		frameWidth:                    frameWidth,
		frameHeight:                   frameHeight,
		isFirst:                       true,
		animationSequence:             animationSequence,
		animationSequenceLastI:        len(animationSequence) - 1,
		currentAnimationSequenceFrame: 0,
	}
}

func (a *Animation) Update() {
	// the first frame needs to draw before you begin updating
	if a.isFirst {
		a.isFirst = false
		return
	}

	a.currentAnimationSequenceFrame++

	if a.currentAnimationSequenceFrame > a.animationSequenceLastI {
		a.currentAnimationSequenceFrame = 0
	}
}
func (a *Animation) Draw(center physics.V, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	// This first translation off sets the image to the center minus the frame size
	op.GeoM.Translate(-float64(a.frameWidth)/2, -float64(a.frameHeight)/2)
	op.GeoM.Translate(center.X(), center.Y())

	frameI := a.animationSequence[a.currentAnimationSequenceFrame]
	screen.DrawImage(a.images[frameI], op)
}
