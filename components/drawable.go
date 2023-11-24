package components

import "github.com/hajimehoshi/ebiten/v2"

type Drawable interface {
	Draw(screen *ebiten.Image)
}
