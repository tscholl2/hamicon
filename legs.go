package main

import (
	"fmt"
	"math/rand"
)

type legs struct {
	length int // 10 [8,15]
}

func newLegs(r *rand.Rand) (l legs) {
	l.length = randint(r, -2, 5)
	return
}

func legsToSVG(d diffs) (svg string) {
	l := 10 + d.legs.length
	if d.body.ry > 0 && d.legs.length <= 0 {
		l = 12
	}
	y := 75
	if d.body.ry < 0 {
		y -= 3
	}
	svg += `<g id="legs" style="stroke:#000;stroke-width:2;stroke-linecap:round;">`
	svg += fmt.Sprintf(`<path id="bleg1" class="leg" d="M30,%d l0,%d" transform="rotate(15,30,75)"/>`, y, l)
	svg += fmt.Sprintf(`<path id="bleg2" class="leg" d="M35,%d l0,%d" transform="rotate(-15,35,75)"/>`, y, l)
	svg += fmt.Sprintf(`<path id="fleg1" class="leg" d="M65,%d l0,%d" transform="rotate(15,65,75)"/>`, y, l)
	svg += fmt.Sprintf(`<path id="fleg2" class="leg" d="M70,%d l0,%d" transform="rotate(-15,70,75)"/>`, y, l)
	svg += `</g>`
	return
}
