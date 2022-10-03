package rand

import "math/rand"

func RandFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = RandFloat(min, max)
	}
	return res
}

func RandFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandInt(n int) int {
	return rand.Intn(n)
}
