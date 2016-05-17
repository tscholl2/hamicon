package main

import (
	"math/rand"
	"strconv"
)

// returns random value in [a,b] (inclusive)
func randint(a, b int) int {
	return rand.Intn(b-a+1) + a
}

func randcolor() string {
	r := strconv.FormatInt(int64(randint(0, 15)), 16)
	g := strconv.FormatInt(int64(randint(0, 15)), 16)
	b := strconv.FormatInt(int64(randint(0, 15)), 16)
	return "#" + r + g + b
}
