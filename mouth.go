package main

import (
	"fmt"
	"math/rand"
)

type mouth struct {
	width int
}

var mouthDefaults = mouth{10}

func newMouth(r *rand.Rand) (m mouth) {
	m.width = randint(r, -2, 2)
	return
}

func mouthToSVG(d diffs) (svg string) {
	w := mouthDefaults.width + d.mouth.width
	x := noseDefaults.x + d.nose.x
	y := noseDefaults.y + d.nose.y
	s := "stroke:#000;stroke-width:1;fill-opacity:0;"
	svg += fmt.Sprintf(`<g id="mouth" style="%s">`, s)
	svg += fmt.Sprintf(`<path id="lip1" d="M%d,%d a%d,%d 0 0,1 %d,0"/>`, x, y, w/2+1, w/2, -w)
	svg += fmt.Sprintf(`<path id="lip2" d="M%d,%d a%d,%d 0 0,0 %d,0"/>`, x, y, w/2+1, w/2, w)
	svg += fmt.Sprintf(`<ellipse id="speaker" class="talk" cx="%d" cy="%d" rx="%d" ry="%d" style="fill:#000;fill-opacity:1;"/>`, x, y+w/2-2+1, w/2, w/2-2)
	svg += `</g>`
	return
}
