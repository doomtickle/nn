package train

import "github.com/doomtickle/nn/rand"

type Point struct {
	X, Y, Bias  float64
	Label int
}

func f(x float64) float64 {
  return 0.8 * x + 0.2
}

func TrainingData(n int) []Point {
	var (
		p  Point
		td []Point
	)

	for i := 0; i < n; i++ {
		rp := rand.RandFloats(-1, 1, 2)
		p.X = rp[0]
		p.Y = rp[1]
    p.Bias = 1
		if p.Y > f(p.X) {
			p.Label = 1
		} else {
			p.Label = -1
		}

		td = append(td, p)
	}

	return td
}
