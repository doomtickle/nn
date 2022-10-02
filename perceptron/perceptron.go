package perceptron

import "github.com/doomtickle/nn/rand"

const (
	lr = 0.1
)

type Perceptron struct {
	Weights []float64
}

func NewPerceptron(numInputs int) *Perceptron {
	var p Perceptron

	// Initialize Weights randomly
	p.Weights = rand.RandFloats(-1, 1, numInputs)

	return &p
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
