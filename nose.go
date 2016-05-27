package main

import (
	"fmt"
	"math/rand"
)

type nose struct {
	length int // 5 [4,6]
	width  int // 6 [4,8]
	x      int
	y      int
}

var noseDefaults = nose{5, 5, 65, 50}

func newNose(r *rand.Rand) (n nose) {
	n.length = randint(r, -1, 1)
	n.width = randint(r, -1, 1)
	return
}

func noseToSVG(d diffs) (svg string) {
	x := noseDefaults.x
	y := noseDefaults.y
	l := noseDefaults.length + d.nose.length
	w := noseDefaults.width + d.nose.width
	svg += `<g id="nose" class="wiggle" style="stroke:#000;stroke-width:1;fill:pink;">`
	svg += fmt.Sprintf(`<path d="M%d,%d l%d,%d l%d,0 Z"/>`, x, y, -w/2, -l, w)
	svg += `</g>`
	return
}
