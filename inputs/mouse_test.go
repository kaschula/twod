package inputs

import (
	"github.com/kaschula/twod/physics"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMouseThrottler_Pressed(t *testing.T) {
	throttler := NewMousePressedThrottler(time.Duration(500)*time.Millisecond, alwaysReturnsPressed)
	resultStream := make(chan bool)

	go func(t *MouseThrottler, r chan<- bool) {
		for i := 0; i < 6; i++ {
			_, pressed := t.Pressed()
			r <- pressed

			time.Sleep(100 * time.Millisecond)
		}

		close(resultStream)
	}(throttler, resultStream)

	results := []bool{}
	for res := range resultStream {
		results = append(results, res)
	}

	expected := []bool{true, false, false, false, false, true}

	assert.Equal(t, expected, results)
}

func alwaysReturnsPressed() (physics.V, bool) {
	return physics.ZeroVector, true
}
