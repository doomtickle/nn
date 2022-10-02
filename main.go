package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/doomtickle/nn/perceptron"
	"github.com/doomtickle/nn/train"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	p := perceptron.NewPerceptron(3)
	data := train.TrainingData(100)
  correct := 0
	for correct<100 {
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
