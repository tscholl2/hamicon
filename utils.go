package main

import (
	"math/rand"
	"strconv"
)

// returns random value in [a,b] (inclusive)
func randint(r *rand.Rand, a, b int) int {
	return r.Intn(b-a+1) + a
}

func randcolor(rnd *rand.Rand) string {
	r := strconv.FormatInt(int64(randint(rnd, 0, 15)), 16)
	g := strconv.FormatInt(int64(randint(rnd, 0, 15)), 16)
	b := strconv.FormatInt(int64(randint(rnd, 0, 15)), 16)
	return "#" + r + g + b
}
