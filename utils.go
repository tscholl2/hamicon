package main

import (
	"fmt"
	"math/rand"
)

// returns random value in [a,b] (inclusive)
func randint(r *rand.Rand, a, b int) int {
	return r.Intn(b-a+1) + a
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

func min(a, b int) int {
	return -max(-a, -b)
}
