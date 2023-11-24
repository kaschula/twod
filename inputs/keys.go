package inputs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func KeyPressAndHold(k ebiten.Key) bool {
	return inpututil.IsKeyJustPressed(k) || inpututil.KeyPressDuration(k) > 0
}

func KeyJustPressed(k ebiten.Key) bool {
	return inpututil.IsKeyJustPressed(k)
}
