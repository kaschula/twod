package inputs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kaschula/twod/physics"
	"time"
)

func MouseCursor() physics.V {
	mouseX, mouseY := ebiten.CursorPosition()
	return physics.New(float64(mouseX), float64(mouseY))
}

func MouseLeftPressed() (physics.V, bool) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		return physics.New(float64(mouseX), float64(mouseY)), true
	}

	return physics.ZeroVector, false
}

func MouseRightPressed() (physics.V, bool) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		mouseX, mouseY := ebiten.CursorPosition()
		return physics.New(float64(mouseX), float64(mouseY)), true
	}

	return physics.ZeroVector, false
}

type MousePressedFN func() (physics.V, bool)
type MouseThrottler struct {
	throttleBy time.Duration
	throttleFN MousePressedFN
	nextClick  int64 // unix time Stamp
}

// warning: throttle does not allow for a holding action, for example if you throttle a mouse button you cant drag with that button
func NewMousePressedThrottler(throttleBy time.Duration, throttleFN MousePressedFN) *MouseThrottler {
	return &MouseThrottler{throttleBy: throttleBy, nextClick: 0, throttleFN: throttleFN}
}

func (throttler *MouseThrottler) Pressed() (physics.V, bool) {
	now := time.Now()
	nowMS := now.UnixMilli()

	if nowMS < throttler.nextClick {
		return physics.ZeroVector, false
	}

	mouseV, pressed := throttler.throttleFN()
	if !pressed {
		return mouseV, false
	}

	throttler.nextClick = now.Add(throttler.throttleBy).UnixMilli()

	return mouseV, pressed
}

// todo next Mouse Point Object

// has Update to track frames
// does the throttling
// tracks left right button clicks and releases
