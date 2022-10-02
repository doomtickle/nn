package main

import (
	"fmt"
	"math/rand"
	"time"
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

// activation function
func sign(n float64) int {
	if n >= 0 {
		return 1
	}

	return -1
}

var (
	inputs = []float64{-1, 0.5}
)

func main() {
  rand.Seed(time.Now().UnixNano())

	for i := 0; i < 150; i++ {
    p := NewPerceptron(len(inputs))
		fmt.Println(p.Guess(inputs))
	}
}
