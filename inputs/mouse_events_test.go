package inputs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kaschula/twod/physics"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewMouseInputProducer_pressed_and_released(t *testing.T) {
	currentV := physics.New(10, 10)
	leftPressed := physics.New(20, 10)

	leftPressedCount := 0

	mouseLeftPressed := func() (physics.V, bool) {
		if leftPressedCount == 0 {
			leftPressedCount++
			return leftPressed, true

		}

		return physics.ZeroVector, false
	}
	mouseRightPressed := func() (physics.V, bool) {
		return physics.ZeroVector, true
	}

	currentPosition := func() physics.V {
		return currentV
	}

	inputProducer := NewMouseInputProducer(currentPosition, mouseLeftPressed, mouseRightPressed, 0, time.Duration(1000)*time.Millisecond)

	inputsFrame1 := inputProducer.Update()
	inputsFrame2 := inputProducer.Update()

	expectedFrame1 := NewInputMouseEvent(NewMouseEvent(ebiten.MouseButtonLeft, MouseButtonPressed, leftPressed))
	expectedFrame2 := NewInputMouseEvent(NewMouseEvent(ebiten.MouseButtonLeft, MouseButtonReleased, currentV))

	assert.Equal(t, expectedFrame1, inputsFrame1[0])
	assert.Equal(t, expectedFrame2, inputsFrame2[0])

}
