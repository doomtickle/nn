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

  LearningRate float64
}

func NewNeuralNet(inputs, neurons, outputs int, lr float64) *NeuralNetwork {
	var nn NeuralNetwork

	nn.InputNodes = inputs
	nn.HiddenNodes = neurons
	nn.OutputNodes = outputs
  nn.LearningRate = lr

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

func DSigmoid(y float64) float64 {
	return y * (1 - y)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	nn := NewNeuralNet(2, 2, 1, 0.1)
	in := []float64{1, 0}
	targets := []float64{1}

	err := nn.Train(in, targets)
	fmt.Printf("%v", err)

	// out := nn.FeedForward(in)
	// fmt.Printf("%v", targets)
}

func (nn *NeuralNetwork) Train(in, targets []float64) []float64 {
	inputs := matrix.FromArray(in)
	hidden := matrix.Multiply(nn.WeightsIH, inputs)
	hidden.AddMatrix(nn.BiasH)
	hidden.Map(Sigmoid)

	op := matrix.Multiply(nn.WeightsHO, hidden)
	op.AddMatrix(nn.BiasO)
	op.Map(Sigmoid)
	ys := matrix.FromArray(targets)

	// Calculate Error
	outputErr := matrix.Subtract(ys, op)
  // Calculate Gradient
  gradients := matrix.StaticMap(op, DSigmoid) 
  fmt.Printf("%+v\n", gradients)
  fmt.Printf("%+v\n", outputErr)
  gradients.MultMatrixByElement(outputErr)
  gradients.Scale(nn.LearningRate)


  //Calculate Deltas
  hiddenT := matrix.Transpose(hidden)
  weight_HO_deltas := matrix.Multiply(gradients, hiddenT)


  nn.WeightsHO.AddMatrix(weight_HO_deltas)

	who_t := matrix.Transpose(nn.WeightsHO)
	hiddenErr := matrix.Multiply(who_t, outputErr)


  // Calculate Hidden gradients
  hidden_gradients := matrix.StaticMap(hidden, DSigmoid)
  fmt.Printf("%+v\n", hidden_gradients)
  fmt.Printf("%+v\n", hiddenErr)
  hidden_gradients.MultMatrixByElement(hiddenErr)
  hidden_gradients.Scale(nn.LearningRate)

  // Calculate Hidden deltas
  inputsT := matrix.Transpose(inputs)
  weight_IH_deltas := matrix.Multiply(hidden_gradients, inputsT)

  nn.WeightsIH.AddMatrix(weight_IH_deltas)

	return nn.WeightsHO.ToArray()
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
