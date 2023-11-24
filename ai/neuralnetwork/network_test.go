package neuralnetwork

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNeuralNetwork(t *testing.T) {
	nn := CreateRandomizedNeuralNetwork(4, 5, 3)

	assert.Len(t, nn.Levels, 2)
	assert.Len(t, nn.Levels[0].Inputs, 4)
	assert.Len(t, nn.Levels[0].Outputs, 5)
	assert.Len(t, nn.Levels[1].Inputs, 5)
	assert.Len(t, nn.Levels[1].Outputs, 3)
}

func TestNewNeuralNetwork_FeedForward(t *testing.T) {
	nn := CreateRandomizedNeuralNetwork(4, 5, 3)
	outputs := nn.FeedForward([]float64{0.1, 0.2, 0.3, 0.4})

	assert.Len(t, outputs, 3)
}
