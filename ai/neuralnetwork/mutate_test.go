package neuralnetwork

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMutate(t *testing.T) {
	nn := CreateRandomizedNeuralNetwork(2, 3, 1)
	copyNN := *nn

	mutated := Mutate(nn, 0.1)

	assert.NotEqual(t, &copyNN.Levels[0].Biases, mutated.Levels[0].Biases)
	assert.NotEqual(t, &copyNN.Levels[1].Biases, mutated.Levels[1].Biases)
	assert.NotEqual(t, &copyNN.Levels[0].Weights, mutated.Levels[0].Weights)
	assert.NotEqual(t, &copyNN.Levels[1].Weights, mutated.Levels[1].Weights)
}

func TestMutate_by_zero_returns_the_same(t *testing.T) {
	nn := CreateRandomizedNeuralNetwork(2, 3, 1)
	copyNN := *nn

	mutated := Mutate(nn, 0)

	assert.NotEqual(t, &copyNN.Levels[0].Biases, mutated.Levels[0].Biases)
	assert.NotEqual(t, &copyNN.Levels[1].Biases, mutated.Levels[1].Biases)
	assert.NotEqual(t, &copyNN.Levels[0].Weights, mutated.Levels[0].Weights)
	assert.NotEqual(t, &copyNN.Levels[1].Weights, mutated.Levels[1].Weights)
}

func TestMutateAndClone(t *testing.T) {
	original := CreateRandomizedNeuralNetwork(3, 2, 4)

	cloned := CloneMutate(original, 0.1)

	assert.Equal(t, original.ID, cloned.PreviousID)
	assert.NotEqual(t, original.ID, cloned.ID)
	assert.Equal(t, cloned.Levels[0].Inputs, original.Levels[0].Inputs)
	assert.Equal(t, cloned.Levels[1].Inputs, original.Levels[1].Inputs)
	assert.Equal(t, cloned.Levels[0].Outputs, original.Levels[0].Outputs)
	assert.Equal(t, cloned.Levels[1].Outputs, original.Levels[1].Outputs)
	assert.NotEqual(t, cloned.Levels[0].Biases, original.Levels[0].Biases)
	assert.NotEqual(t, cloned.Levels[1].Biases, original.Levels[1].Biases)
	assert.NotEqual(t, cloned.Levels[0].Weights, original.Levels[0].Weights)
	assert.NotEqual(t, cloned.Levels[1].Weights, original.Levels[1].Weights)
}
