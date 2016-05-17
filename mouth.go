package main

import "fmt"

/*
<g id="mouth" stroke-linecap="round" stroke="#000" stroke-width="4" fill-opacity="0" fill="#000">
	<path id="lip" d="M50,50 a 30,60 0 0,0 30,0"/>
</g>
*/
// <path id="lip" d="M50,50 a 30,60 30 0,0 30,0"/>
type mouth struct {
	angle int    // 0 [0,90] 0=flat, 90=circluar
	width int    // 30 [20,40]
	frown int    // 0 [0,1]
	style string // ""
}

func newMouth() (m mouth) {
	m.angle = randint(0, 90)
	m.width = randint(-10, 10)
	if randint(0, 4) == 0 {
		m.frown = 1
	}
	m.style = "stroke:#000;stroke-width:4;stroke-linecap:round;fill-opacity:0;"
	return
}

func mouthToSVG(d diffs) (svg string) {
	w := 30 + d.mouth.width
	x := 65 - w/2
	a := d.mouth.angle
	f := d.mouth.frown
	y := 50
	if f > 0 {
		y += 3
	}
	svg += `<g id="mouth" style="` + d.mouth.style + `">`
	svg += `<path id="lip" ` + fmt.Sprintf(`d="M%d,%d a %d,60 %d 0,%d %d,0"/>`,
		x, y, w, a, f, w)
	svg += `</g>`
	return
}
