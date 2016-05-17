package main

import (
	"fmt"
	"math/rand"
)

/*
<g id="eyes" style="stroke:#000;stroke-width:2;fill:#fff;">
	<circle id="eye1" r="3" cy="35" cx="60"/>
	<circle id="eye2" r="3" cy="35" cx="70"/>
</g>
*/
type eyes struct {
	r1    int    // 3 [2,6]
	h1    int    // 35 [34,36]
	r2    int    // 3 [2,6]
	h2    int    // 35 [34,36]
	w     int    // 4 [2,6]
	style string // ""
}

func newEyes(r *rand.Rand) (e eyes) {
	e.r1 = randint(r, 0, 2)
	e.h1 = randint(r, -1, 1)
	e.r2 = randint(r, 0, 2)
	e.h2 = randint(r, -1, 1)
	e.w = randint(r, -2, 2)
	e.style = "stroke:#000;stroke-width:2;fill:#fff;"
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
