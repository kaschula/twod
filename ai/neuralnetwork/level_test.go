package neuralnetwork

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_randomFloat64(t *testing.T) {
	for i := 0; i < 100; i++ {
		randomFloat := randomFloat64()

		fmt.Println(randomFloat)

		if randomFloat < -1 || randomFloat > 1 {
			t.Fatalf("expect %v ti be between -1 and 1", randomFloat)
		}
	}
}

func TestLevelRandomise(t *testing.T) {
	level := Randomize(NewLevel(5, 4))

	fmt.Println("level", level)
}

func TestFeedForward(t *testing.T) {
	level := setWeightsAndBiases(NewLevel(5, 2))

	outputs := feedForward([]float64{0.8, 0.3, 0.2, 0.2, 0.2}, level)

	assert.Len(t, outputs, 2)
	assert.Equal(t, float64(1), outputs[0])
	assert.Equal(t, float64(1), outputs[1])
}

func setWeightsAndBiases(level *Level) *Level {
	for i := range level.Inputs {
		for j := range level.Outputs {
			level.Weights[i][j] = 0.4
		}
	}

	for i := range level.Biases {
		level.Biases[i] = 0.5
	}

	return level
}
