package training

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kaschula/twod/ai/neuralnetwork"
	"github.com/kaschula/twod/file"
	"os"
	"strings"
	"time"
)

type SimulationConfig struct {
	// directory to store generation data
	FilePath        string `json:"FilePath"`
	NetworkFileName string `json:"NetworkFileName"`
	LevelID         string `json:"levelID"`
	// number of games per generation
	Population int `json:"Population"`
	// number of Generations
	Generations  int     `json:"Generations"`
	MutationRate float64 `json:"MutationRate"`
	// number of network inputs
	NetworkInputs int `json:"NetworkInputs"`
	CreateGame    func(config SimulationConfig, network *neuralnetwork.NeuralNetwork) TrainingGame
}

type SimulationResult struct {
	generation, topScore int
}

func (c SimulationConfig) String() string {
	return fmt.Sprintf(
		"Population: %v | Generations: %v | MutationRate: %.2f | #Inputs: %v",
		c.Population,
		c.Generations,
		c.MutationRate,
		c.NetworkInputs,
	)
}

func RunSimulation(config SimulationConfig) error {
	if config.CreateGame == nil {
		return errors.New("create game config parameter required")
	}

	start := time.Now()
	var results []SimulationResult

	defer func() {
		for _, result := range results {
			fmt.Printf("generation: %v topScore: %v \n", result.generation, result.topScore)
		}
	}()

	bestTrainedNetwork := TrainedNetwork{
		score:   0,
		network: neuralnetwork.CreateRandomizedNeuralNetwork(config.NetworkInputs, 5, 4), // start with random network
	}
	fmt.Printf("******* Starting simulation on Level '%v' running %v generations with populations of %v *******\n", config.LevelID, config.Generations, config.Population)
	for generation := 0; generation < config.Generations; generation++ {
		if generation != 0 && bestTrainedNetwork.score == 0 {
			// start new random network if best score is 0.
			bestTrainedNetwork.network = neuralnetwork.CreateRandomizedNeuralNetwork(config.NetworkInputs, 5, 4)
		}

		fmt.Printf(">>>>>>>>> generation %v running games %v \n", generation, config.Population)
		trainedNetwork, err := trainGeneration(config, bestTrainedNetwork.network, generation)
		if err != nil {
			return fmt.Errorf("%w trainGenerationError", err)
		}

		if trainedNetwork.network == nil {
			fmt.Println("warning........ topNetworkInGeneration network is nil")
		}

		if trainedNetwork.score > bestTrainedNetwork.score {
			bestTrainedNetwork = trainedNetwork
		}

		err = storeNetwork(config, generation, bestTrainedNetwork)
		if err != nil {
			return fmt.Errorf("%w persist nerual network error for generation %v", err, generation)
		}

		results = append(results, SimulationResult{generation: generation, topScore: bestTrainedNetwork.score})
	}

	writeResults(config, start, time.Now(), results)

	return nil
}

func storeNetwork(config SimulationConfig, generation int, trainedNetwork TrainedNetwork) error {
	filePath := fmt.Sprintf("%v/%v_%v", config.FilePath, generation, config.NetworkFileName)
	file.CreateDirIfNotExists(fmt.Sprintf("%v", config.FilePath))

	return NewPersistedNetwork(trainedNetwork).Write(filePath)
}

type TrainingGame interface {
	Update() error
	GameOver() bool
	Score() int
	GetNetwork() *neuralnetwork.NeuralNetwork
}

type PersistedNetwork struct {
	Score         int                          `json:"score"`
	NeuralNetwork *neuralnetwork.NeuralNetwork `json:"neuralNetwork"`
}

func NewPersistedNetwork(trainedNetwork TrainedNetwork) *PersistedNetwork {
	return &PersistedNetwork{Score: trainedNetwork.score, NeuralNetwork: trainedNetwork.network}
}

func (pn *PersistedNetwork) Write(filePath string) error {
	byt, err := json.Marshal(pn)
	if err != nil {
		return fmt.Errorf("%w network bytes marshall", err)
	}

	return os.WriteFile(filePath, byt, 0644)
}

func LoadPersistedNetwork(filePath string) (*PersistedNetwork, error) {
	contents, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("%w reading file error", err)
	}

	pn := &PersistedNetwork{}
	err = json.Unmarshal(contents, pn)
	if err != nil {
		return nil, fmt.Errorf("%w json unmarshall", err)
	}

	return pn, nil
}

func getNetworks(config SimulationConfig, latestBestNetwork *neuralnetwork.NeuralNetwork) ([]*neuralnetwork.NeuralNetwork, error) {
	if latestBestNetwork != nil {
		nns := neuralnetwork.MutateToMany(config.Population, latestBestNetwork, config.MutationRate)
		nns[0] = latestBestNetwork // always keep the orginal highest score netwirk in case muatation end up this should prevent the mutations kill off a scoring generation

		return nns, nil
	}

	return make([]*neuralnetwork.NeuralNetwork, config.Population), nil
}

type TrainedNetwork struct {
	score   int
	network *neuralnetwork.NeuralNetwork
}

func trainNetwork(config SimulationConfig, generation, populationI int, nn *neuralnetwork.NeuralNetwork) TrainedNetwork {
	if nn == nil {
		fmt.Println("warning........ trainNetwork() network is nil")
	}

	game := config.CreateGame(config, nn)

	frameLimit := 2400 // 40s at 60 Frames/second
	frameCount := 0

	for {
		if frameCount > frameLimit {
			break
		}

		updateErr := game.Update()
		if updateErr != nil {
			fmt.Printf("%s update error for generation %v Population %v\n", updateErr.Error(), generation, populationI)
			break
		}

		if game.GameOver() {
			break
		}

		frameCount++
	}

	return TrainedNetwork{
		score:   game.Score(),
		network: game.GetNetwork(),
	}
}

func trainGeneration(config SimulationConfig, nn *neuralnetwork.NeuralNetwork, generation int) (TrainedNetwork, error) {
	mutatedNetworks, err := getNetworks(config, nn)
	if err != nil {
		return TrainedNetwork{}, fmt.Errorf("%w error creating networks generation %v", err, generation)
	}

	trainedNetwork := trainPopulation(config, generation, mutatedNetworks)

	fmt.Printf("Generation: %v score %v NetworkID = %v Games played %v \n", generation, trainedNetwork.score, trainedNetwork.network.ID)

	return trainedNetwork, nil
}

func trainPopulation(config SimulationConfig, generation int, mutatedNetworks []*neuralnetwork.NeuralNetwork) TrainedNetwork {
	bestNetwork := TrainedNetwork{
		score: -1,
	}

	for populationI := 0; populationI < config.Population; populationI++ {
		network := mutatedNetworks[populationI]

		trainedNetwork := trainNetwork(config, generation, populationI, network)
		if trainedNetwork.score > bestNetwork.score {
			bestNetwork = trainedNetwork
		}
	}

	return bestNetwork
}

func writeResults(config SimulationConfig, start, end time.Time, results []SimulationResult) {
	// prints as well
	runDuration := end.Sub(start)
	file := []string{}

	file = append(file, fmt.Sprintf("Simulation Run at %v, total duration %v", start.Format(time.RFC3339), runDuration.String()))
	file = append(file, config.String())
	file = append(file, "-----------")
	for _, result := range results {
		line := fmt.Sprintf("generation: %v topScore: %v", result.generation, result.topScore)
		file = append(file, line)
	}

	f, err := os.Create(config.FilePath + "/output.txt")
	if err != nil {
		fmt.Printf("warning results file create failed %v", err.Error())
	}

	_, err = f.Write([]byte(strings.Join(file, "\n")))
	if err != nil {
		fmt.Printf("warning results file write failed %v", err.Error())
	}
}
