package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kaschula/twod/physics"
	"image/color"
)

// Background Scrolling background, resets moves when it y == 0. This makes it feel continuos
type Background struct {
	vertical bool
	color    color.Color
	body     physics.Rectangle
	reset    float64
	draw     func(r physics.Rectangle, screen *ebiten.Image)
}

func (b Background) Draw(screen *ebiten.Image) {
	b.draw(b.body, screen)
}

func NewBackground(r physics.Rectangle, reset float64, draw func(r physics.Rectangle, screen *ebiten.Image)) *Background {
	return &Background{vertical: true, reset: reset, body: r, draw: draw}
}

func (b Background) Body() physics.Rectangle {
	return b.body
}

func (b Background) Transpose(by physics.V) {
	if by.Equal(physics.ZeroVector) {
		return
	}

	b.body.Move(by)
	b.body.Update()

	//fmt.Println("----", b.body.TopLeft())

	if b.vertical {
		// background reaches y reset
		if b.body.TopLeft().Y() >= 0 {
			b.body.Move(physics.New(0, b.reset))
			b.body.Update()
		}
	}
}
