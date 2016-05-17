package main

import "fmt"

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

func newLegs() (l legs) {
	l.length = randint(-2, 5)
	l.style = "stroke:#000;stroke-width:2;stroke-linecap:round;"
	return
}

func legsToSVG(d diffs) (svg string) {
	svg += `<g id="legs" style="` + d.legs.style + `">`
	svg += fmt.Sprintf(`<path id="bleg1" d="M30,75 l0,10" transform="rotate(15,30,75)"/>`)
	svg += fmt.Sprintf(`<path id="bleg2" d="M35,75 l0,10" transform="rotate(-15,35,75)"/>`)
	svg += fmt.Sprintf(`<path id="fleg1" d="M65,75 l0,10" transform="rotate(15,65,75)"/>`)
	svg += fmt.Sprintf(`<path id="fleg2" d="M70,75 l0,10" transform="rotate(-15,70,75)"/>`)
	svg += `</g>`
	return
}
