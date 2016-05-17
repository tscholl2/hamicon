package main

import "math/rand"

// returns random value in [a,b] (inclusive)
func randint(a, b int) int {
	return rand.Intn(b-a+1) + a
}
