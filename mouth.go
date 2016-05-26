package main

import (
	"fmt"
	"math/rand"
)

/*
<g id="mouth" style="stroke:#000;stroke-width:4;stroke-linecap:round;fill-opacity:0;">
	<path id="lip" d="M50,50 a 30,60 0 0,0 30,0"/>
</g>
*/
type mouth struct {
	angle int  // 0 [0,90] 0=flat, 90=circluar
	width int  // 30 [20,40]
	frown int  // 0 [0,1]
	fill  bool // false
}

func newMouth(r *rand.Rand) (m mouth) {
	m.angle = randint(r, 0, 90)
	m.width = randint(r, -10, 10)
	if randint(r, 0, 4) == 0 {
		m.frown = 1
	}
	if randint(r, 0, 4) == 0 {
		m.fill = true
	}
	return
}

func mouthToSVG(d diffs) (svg string) {
	w := 30 + d.mouth.width
	x := 65 - w/2
	a := d.mouth.angle
	f := d.mouth.frown
	y := 55
	if f > 0 {
		y += 3
	}
	var z, s string
	if d.mouth.fill {
		z = "Z"
		s = `fill="#fff" stroke-width="2" fill-opacity="1"`
	}
	svg += `<g class="talk" id="mouth" style="stroke:#000;stroke-width:4;stroke-linecap:round;fill-opacity:0;">`
	svg += fmt.Sprintf(`<path id="lip" d="M%d,%d a %d,60 %d 0,%d %d,0%s" %s/>`, x, y, w, a, f, w, z, s)
	svg += `</g>`
	return
}
