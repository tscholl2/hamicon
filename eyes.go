package main

import "fmt"

// <circle id="eye1" r="3" cy="35" cx="60"/>
// <circle id="eye2" r="3" cy="35" cx="70"/>
type eyes struct {
	r1    int    // [-1,3]
	h1    int    // [-1,1]
	r2    int    // [-1,3]
	h2    int    // [-1,1]
	w     int    // [0,1]
	style string // ""
}

func newEyes() (e eyes) {
	e.r1 = randint(0, 2)
	e.h1 = randint(-1, 1)
	e.r2 = randint(0, 2)
	e.h2 = randint(-1, 1)
	e.w = randint(-2, 2)
	e.style = "stroke:#000;stroke-width:2;fill-opacity:0;"
	return
}

func eyesToSVG(d diffs) (svg string) {
	r1 := 3 + d.eyes.r1
	r2 := 3 + d.eyes.r2
	w := 4 + d.eyes.w
	x1 := 65 - r1 - w/2
	y1 := 35 + d.eyes.h1
	x2 := 65 + r2 + w/2
	y2 := 35 + d.eyes.h2
	svg += `<g id="eyes" style="` + d.eyes.style + `">`
	svg += fmt.Sprintf(`<circle id="eye1" r="%d" cx="%d" cy="%d"/>`, r1, x1, y1)
	svg += fmt.Sprintf(`<circle id="eye2" r="%d" cx="%d" cy="%d"/>`, r2, x2, y2)
	svg += `</g>`
	return
}
