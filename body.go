package main

import (
	"fmt"
	"math/rand"
)

/*
<ellipse id="body" cx="50" cy="50" rx="45" ry="30" style="fill:#fff;fill-opacity:1;stroke:#000;stroke-width:2;"/>
*/
type body struct {
	rx    int    // 45 [42,45]
	ry    int    // 30 [25,35]
	color string // "#fff"
}

func newBody(r *rand.Rand) (b body) {
	b.rx = randint(r, -3, 0)
	b.ry = randint(r, -5, 5)
	b.color = randcolor(r)
	return
}

func bodyToSVG(d diffs) string {
	rx := 45 + d.body.rx
	ry := 30 + d.body.ry
	cx := 50
	cy := 50
	c := d.body.color
	if c == "" {
		c = "#fff"
	}
	return fmt.Sprintf(`<ellipse id="body" cx="%d" cy="%d" rx="%d" ry="%d" style="fill-opacity:1;stroke:#000;stroke-width:2;fill:%s;"/>`,
		cx, cy, rx, ry, c)
}
