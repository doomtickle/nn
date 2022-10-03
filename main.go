package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/doomtickle/nn/matrix"
	"github.com/doomtickle/nn/perceptron"
	"github.com/doomtickle/nn/train"
)

type NeuralNetwork struct {
	InputNodes,
	HiddenNodes,
	OutputNodes int

	WeightsIH,
	WeightsHO,
	BiasH,
	BiasO *matrix.Matrix
}

func NewNeuralNet(inputs, neurons, outputs int) *NeuralNetwork {
	var nn NeuralNetwork

	nn.InputNodes = inputs
	nn.HiddenNodes = neurons
	nn.OutputNodes = outputs

	nn.WeightsIH = matrix.NewMatrix(nn.HiddenNodes, nn.InputNodes)
	nn.WeightsHO = matrix.NewMatrix(nn.OutputNodes, nn.HiddenNodes)

	nn.WeightsIH.Randomize()
	nn.WeightsHO.Randomize()

	nn.BiasH = matrix.NewMatrix(nn.InputNodes, 1)
	nn.BiasO = matrix.NewMatrix(nn.HiddenNodes, 1)

	nn.BiasH.Randomize()
	nn.BiasO.Randomize()

	return &nn
}

func (nn *NeuralNetwork) FeedForward(in []float64) []float64 {
	inputs := matrix.FromArray(in)
	hidden := matrix.Multiply(nn.WeightsIH, inputs)
	hidden.AddMatrix(nn.BiasH)
	hidden.Map(Sigmoid)

	op := matrix.Multiply(nn.WeightsHO, hidden)
	op.AddMatrix(nn.BiasO)
	op.Map(Sigmoid)

	return op.ToArray()
}

func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	nn := NewNeuralNet(3, 3, 3)
	in := []float64{1, 0, 0.5}
	out := nn.FeedForward(in)
	fmt.Printf("%v", out)
}

func simple() {
	p := perceptron.NewPerceptron(3)
	data := train.TrainingData(100)
	correct := 0
	for correct < 100 {
		correct = 0
		for _, d := range data {
			inputs := []float64{d.X, d.Y, d.Bias}
			target := d.Label
			p.Train(inputs, target)
			guess := p.Guess(inputs)
			if guess == target {
				correct = correct + 1
			}
		}
		fmt.Printf("Correct: %d/100\n", correct)
	}

}
