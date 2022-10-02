package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	lr = 0.1
)

type Perceptron struct {
	Weights []float64
}

func NewPerceptron(numWeights int) *Perceptron {
	var p Perceptron

	// Initialize Weights randomly
	p.Weights = randFloats(-1, 1, numWeights)

	return &p
}

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

func (p *Perceptron) Guess(inputs []float64) int {
	sum := float64(0)
	for i := 0; i < len(p.Weights); i++ {
		sum += inputs[i] * p.Weights[i]
	}

	output := sign(sum)

	return output
}

func (p *Perceptron) Train(inputs []float64, target int) {
  var ws []float64
	guess := p.Guess(inputs)
	err := target - guess
	for i, w := range p.Weights {
		ws = append(ws, (w + float64(err) * inputs[i] * lr))
	}

  p.Weights = ws
}

// activation function
func sign(n float64) int {
	if n >= 0 {
		return 1
	}

	return -1
}

type TrainingPoint struct {
	X, Y  float64
	Label int
}

func trainingData(n int) []TrainingPoint {
	var (
		p  TrainingPoint
		td []TrainingPoint
	)

	for i := 0; i < n; i++ {
		rp := randFloats(0, 400, 2)
		p.X = rp[0]
		p.Y = rp[1]
		if p.X > p.Y {
			p.Label = 1
		} else {
			p.Label = -1
		}

		td = append(td, p)
	}

	return td
}

func main() {
	rand.Seed(time.Now().UnixNano())
	p := NewPerceptron(2)
	data := trainingData(100)
  correct := 0
	for i := 0; i < 10; i++ {
    correct = 0
		for _, d := range data {
			inputs := []float64{d.X, d.Y}
			target := d.Label
			p.Train(inputs, target)
      guess := p.Guess(inputs)
      fmt.Printf("%d, %d\n", guess, target)
			if guess == target {
				correct = correct + 1
			}
		}
		fmt.Printf("Correct: %d/100\n", correct)
	}
}
