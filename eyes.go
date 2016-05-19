package main

import (
	"fmt"
	"math/rand"
)

/*
<g id="eyes" style="stroke:#000;stroke-width:2;fill:#fff;">
	<ellipse id="eye1" rx="3" ry="3" cy="35" cx="60"/>
	<ellipse id="eye2" rx="3" ry="3" cy="35" cx="70"/>
</g>
*/
type eyes struct {
	r1      int  // 3 [2,6]
	h1      int  // 35 [34,36]
	r2      int  // 3 [2,6]
	h2      int  // 35 [34,36]
	w       int  // 4 [2,6]
	glasses bool // false
}

func newEyes(r *rand.Rand) (e eyes) {
	e.r1 = randint(r, 0, 2)
	e.h1 = randint(r, -1, 1)
	e.r2 = randint(r, 0, 2)
	e.h2 = randint(r, -1, 1)
	e.w = randint(r, -2, 2)
	e.glasses = randint(r, 0, 1) == 0
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
	if !d.eyes.glasses {
		svg += `<g id="eyes" style="stroke:#000;stroke-width:2;fill:#fff;">`
		svg += fmt.Sprintf(`<ellipse class="blink" id="eye1" cx="%d" cy="%d" rx="%d" ry="%d"/>`, x1, y1, r1, r1)
		svg += fmt.Sprintf(`<ellipse class="blink" id="eye2" cx="%d" cy="%d" rx="%d" ry="%d"/>`, x2, y2, r2, r2)
		svg += `</g>`
	} else {
		gr := max(r1, r2) + 3
		gy := (y1 + y2) / 2
		gx1 := x1
		gx2 := x2
		if gr > (x2-x1)/2 {
			gx1 -= gr - (x2-x1)/2
			gx2 += gr - (x2-x1)/2
		}
		svg += `<g id="eyes" style="stroke:#000;stroke-width:1;fill:#fff;">`
		svg += `<g id="glasses" style="stroke-width:2;fill-opacity:0;">`
		svg += fmt.Sprintf(`<circle cx="%d" cy="%d" r="%d"/>`, gx1, gy, gr)
		svg += fmt.Sprintf(`<circle cx="%d" cy="%d" r="%d"/>`, gx2, gy, gr)
		svg += fmt.Sprintf(`<path d="M%d,%d l%d,0"/>`, gx1+gr, gy, gx2-gx1-2*gr)
		svg += fmt.Sprintf(`<path d="M%d,%d l-10,0" transform="rotate(15,%d,%d)"/>`, gx1-gr, gy, gx1-gr, gy)
		svg += fmt.Sprintf(`<path d="M%d,%d l10,0" transform="rotate(-15,%d,%d)"/>`, gx2+gr, gy, gx2+gr, gy)
		svg += `</g>`
		svg += fmt.Sprintf(`<ellipse class="blink" id="eye1" cx="%d" cy="%d" rx="%d" ry="%d"/>`, x1, y1, r1, r1)
		svg += fmt.Sprintf(`<ellipse class="blink" id="eye2" cx="%d" cy="%d" rx="%d" ry="%d"/>`, x2, y2, r2, r2)
		svg += `</g>`
	}
	return
}
