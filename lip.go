package main

import "fmt"

// <path id="lip" d="M50,50 a 30,60 30 0,0 30,0"/>
type lip struct {
	Angle int    // 0 [0,90] 0=flat, 90=circluar
	Width int    // 30 [20,40]
	Frown int    // 0 [0,1]
	Style string // ""
}

func newLip() (l lip) {
	l.Angle = randint(0, 90)
	l.Width = randint(-10, 10)
	if randint(0, 4) == 0 {
		l.Frown = 1
	}
	return
}

func lipToSVG(d diffs) (svg string) {
	w := 30 + d.lip.Width
	x := 65 - w/2
	a := d.lip.Angle
	f := d.lip.Frown
	s := d.lip.Style
	y := 50
	if f > 0 {
		y += 3
	}
	svg = `<path id="lip"`
	if s != "" {
		svg += ` style="` + s + `"`
	}
	svg += fmt.Sprintf(` d="M%d,%d a %d,60 %d 0,%d %d,0"/>`,
		x, y, w, a, f, w)
	return
}
