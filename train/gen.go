package train

import "github.com/doomtickle/nn/rand"

type Point struct {
	X, Y  float64
	Label int
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
		if p.X > p.Y {
			p.Label = 1
		} else {
			p.Label = -1
		}

		td = append(td, p)
	}

	return td
}
