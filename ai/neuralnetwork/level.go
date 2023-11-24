package neuralnetwork

import "math/rand"

type Level struct {
	Inputs  []float64   `json:"Inputs"`
	Outputs []float64   `json:"Outputs"`
	Biases  []float64   `json:"Biases"`
	Weights [][]float64 `json:"Weights"`
}

func NewLevel(inputs, outputs int) *Level {

	weights := [][]float64{}
	for i := 0; i < inputs; i++ {
		weights = append(weights, make([]float64, outputs))
	}

	return &Level{
		Inputs:  make([]float64, inputs),
		Outputs: make([]float64, outputs),
		Biases:  make([]float64, outputs),
		Weights: weights,
	}
}

func Randomize(level *Level) *Level {
	for i := range level.Inputs {
		for j := range level.Outputs {
			level.Weights[i][j] = randomFloat64()
		}
	}

	for i := range level.Biases {
		level.Biases[i] = randomFloat64()
	}

	return level
}

func (level *Level) FeedForward(givenInputs []float64) []float64 {
	return feedForward(givenInputs, level)
}

func (level *Level) GetInputs() []float64 {
	return level.Inputs
}

func (level *Level) GetOutputs() []float64 {
	return level.Outputs
}

// feedForward applies givenInputs to the level Inputs and calculates the Outputs
func feedForward(givenInputs []float64, level *Level) []float64 {
	if len(givenInputs) != len(level.Inputs) {
		panic("Inputs do not match Outputs")
	}

	for i := range level.Inputs {
		level.Inputs[i] = givenInputs[i]
	}

	// calculate Outputs
	for i := range level.Outputs {
		sum := float64(0)
		for j := range level.Inputs {
			input := level.Inputs[j]
			weight := level.Weights[j][i]

			sum += input * weight
		}

		output := float64(0)
		if sum > level.Biases[i] {
			output = 1
		}

		level.Outputs[i] = output
	}

	return level.Outputs
}

// randomFloat64 return number between -1 and 1
// todo rename
func randomFloat64() float64 {
	return float64(rand.Int63n(200)-100) / 100
}
