package neuralnetwork

import "github.com/kaschula/twod/math"

func Mutate(nn *NeuralNetwork, by float64) *NeuralNetwork {
	for _, level := range nn.Levels {
		for i, bias := range level.Biases {
			newBias := math.Lerp(bias, randomFloat64(), by)
			level.Biases[i] = newBias
		}

		for i, weights := range level.Weights {
			for j, weight := range weights {
				level.Weights[i][j] = math.Lerp(weight, randomFloat64(), by)
			}
		}
	}

	return nn
}

func CloneMutate(original *NeuralNetwork, mutateBy float64) *NeuralNetwork {
	levels := make([]*Level, len(original.Levels))
	for levelI, level := range original.Levels {
		newLevel := &Level{
			Inputs:  make([]float64, len(level.Inputs)),
			Outputs: make([]float64, len(level.Outputs)),
			Biases:  make([]float64, len(level.Outputs)),
			Weights: make([][]float64, len(level.Weights)),
		}

		for i, input := range level.Inputs {
			newLevel.Inputs[i] = input
		}

		for i, output := range level.Outputs {
			newLevel.Outputs[i] = output
		}

		for i, bias := range level.Biases {
			newBias := math.Lerp(bias, randomFloat64(), mutateBy)
			newLevel.Biases[i] = newBias
		}

		for i, weights := range level.Weights {

			newLevel.Weights[i] = make([]float64, len(weights))

			for j, weight := range weights {
				newLevel.Weights[i][j] = math.Lerp(weight, randomFloat64(), mutateBy)
			}
		}

		levels[levelI] = newLevel
	}

	return NewNeuralNetwork(original.ID, levels)
}

// MutateToMany creates a copy of the network then mutates the copy
func MutateToMany(count int, network *NeuralNetwork, by float64) []*NeuralNetwork {
	nns := make([]*NeuralNetwork, count)
	for i := 0; i < count; i++ {
		nns[i] = CloneMutate(network, by)
	}

	return nns
}
