package main

import (
	"fmt"
	"math/rand"
)

/*
<g id="legs" style="stroke:#000;stroke-width:2;stroke-linecap:round;">
  <path id="bleg1" d="M30,75 l0,10" transform="rotate(15,30,75)"/>
  <path id="bleg2" d="M35,75 l0,10" transform="rotate(-15,35,75)"/>
  <path id="fleg1" d="M65,75 l0,10" transform="rotate(15,65,75)"/>
  <path id="fleg2" d="M70,75 l0,10" transform="rotate(-15,70,75)"/>
</g>
*/
type legs struct {
	length int    // 10 [8,15]
	style  string // ""
}

func newLegs(r *rand.Rand) (l legs) {
	l.length = randint(r, -2, 5)
	l.style = "stroke:#000;stroke-width:2;stroke-linecap:round;"
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
	svg += `<g id="legs" style="` + d.legs.style + `">`
	svg += fmt.Sprintf(`<path id="bleg1" d="M30,%d l0,%d" transform="rotate(15,30,75)"/>`, y, l)
	svg += fmt.Sprintf(`<path id="bleg2" d="M35,%d l0,%d" transform="rotate(-15,35,75)"/>`, y, l)
	svg += fmt.Sprintf(`<path id="fleg1" d="M65,%d l0,%d" transform="rotate(15,65,75)"/>`, y, l)
	svg += fmt.Sprintf(`<path id="fleg2" d="M70,%d l0,%d" transform="rotate(-15,70,75)"/>`, y, l)
	svg += `</g>`
	return
}
