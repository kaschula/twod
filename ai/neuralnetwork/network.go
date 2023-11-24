package neuralnetwork

import "github.com/google/uuid"

type NeuralNetwork struct {
	ID         string   `json:"id"`
	Levels     []*Level `json:"levels"`
	PreviousID string   `json:"previousID"`
}

// CreateRandomizedNeuralNetwork creates a new neural network by taking the number of nodes at each layer
func CreateRandomizedNeuralNetwork(levelCounts ...int) *NeuralNetwork {
	if len(levelCounts) < 2 {
		panic("must have at least enough levelCounts for 1 level")
	}

	for _, count := range levelCounts {
		if count == 0 {
			panic("can not have a count of 0")
		}
	}

	levels := []*Level{}
	for i := 0; i < len(levelCounts)-1; i++ {
		levels = append(levels, Randomize(NewLevel(levelCounts[i], levelCounts[i+1])))
	}

	return NewNeuralNetwork("", levels)
}

func NewNeuralNetwork(previousID string, levels []*Level) *NeuralNetwork {
	return &NeuralNetwork{
		Levels:     levels,
		ID:         uuid.NewString(),
		PreviousID: previousID,
	}
}

func (network *NeuralNetwork) FeedForward(inputs []float64) []float64 {
	outputs := inputs
	for _, level := range network.Levels {
		outputs = level.FeedForward(outputs)
	}

	return outputs
}

func (network *NeuralNetwork) GetLevels() []*Level {
	return network.Levels
}
