package main

import (
	"fmt"
	"math/rand"
)

// returns random value in [a,b] (inclusive)
func randint(r *rand.Rand, a, b int) int {
	return r.Intn(b-a+1) + a
}

func randintf(r *rand.Rand, a, b int) float64 {
	return float64(randint(r, a, b))
}

func randcolor(rnd *rand.Rand) string {
	return fmt.Sprintf("#%x%x%x", randint(rnd, 0, 15), randint(rnd, 0, 15), randint(rnd, 0, 15))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
