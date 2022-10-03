package main

import (
	"fmt"
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
}


func main() {
	rand.Seed(time.Now().UnixNano())
  mat1 := matrix.NewMatrix(2,3)
  mat1.Randomize()
	fmt.Printf("%+v\n", mat1)

  mat2 := matrix.NewMatrix(3,2)
  mat2.Randomize()
  fmt.Printf("%+v\n", mat2)

  mat3 := mat1.MultMatrix(mat2)
  fmt.Printf("%+v\n", mat3)

  mat4 := mat2.Transpose()
  fmt.Printf("%+v\n", mat4)
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
