package main

import (
	"fmt"
	"math/rand"
)

/*
  <g class="wiggle" id="nose" style="stroke:#000;stroke-width:1;fill:#fff;">
    <path d="M65,46 l-3,-5 l6,0 z">
  </g>
*/
type nose struct {
	length int // 5 [4,6]
	width  int // 6 [4,8]
}

func newNose(r *rand.Rand) (n nose) {
	n.length = randint(r, -1, 1)
	n.width = randint(r, -1, 1)
	return
}

func noseToSVG(d diffs) (svg string) {
	l := 5 + d.nose.length
	w := 5 + d.nose.width
	svg += `<g id="nose" class="wiggle" style="stroke:#000;stroke-width:1;fill:none;">`
	svg += fmt.Sprintf(`<path d="M65,50 l%d,%d l%d,0 Z"/>`, -w/2, -l, w)
	svg += `</g>`
	return
}
