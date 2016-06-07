package main

import (
	"encoding/xml"
	"fmt"
	"math"
	"math/rand"
)

/*
<svg width="100" height="100" preserveAspectRatio="xMidYMid meet" xmlns="http://www.w3.org/2000/svg" xmlns:svg="http://www.w3.org/2000/svg">
<style type="text/css" >
<![CDATA[
  ... css ...
]]>
</style>
*/

type hamicon struct {
	svg
	XMLName             xml.Name    `xml:"svg"`
	Width               int         `xml:"width,attr"`
	Height              int         `xml:"height,attr"`
	PreserveAspectRatio string      `xml:"preserveAspectRatio,attr"`
	XMLNS               string      `xml:"xmlns,attr"`
	XMLNSSVG            string      `xml:"xmlns:svg,attr"`
	Seed                int64       `xml:"seed,attr"`
	Style               interface{} `xml:"style"`
	Icon                group
}

type randomizable struct {
	body struct {
		rx, ry int
		color  string
	}
	legs struct {
		length, y int
	}
	eyes struct {
		r1, h1, r2, h2, w int
	}
	glasses int // none, round, square
	nose    struct {
		length, width int
	}
	mouth struct {
		width int
	}
	cheeks struct {
		radius int
		inward bool
	}
}

func newRandomizable(seed int64) (r randomizable) {
	rnd := rand.New(rand.NewSource(seed))
	// body options
	r.body.rx = randint(rnd, -3, 0)
	r.body.ry = randint(rnd, -5, 5)
	r.body.color = randcolor(rnd)
	// leg options
	r.legs.length = randint(rnd, -2, 5)
	if r.body.ry > 0 && r.legs.length <= 0 {
		r.legs.length = 2
	}
	if r.body.ry < 0 {
		r.legs.y = -3
	}
	// eyes options
	r.eyes.r1 = randint(rnd, 0, 2)
	r.eyes.h1 = randint(rnd, -1, 1)
	r.eyes.r2 = randint(rnd, 0, 2)
	r.eyes.h2 = randint(rnd, -1, 1)
	r.eyes.w = randint(rnd, -2, 2)
	// glasses options
	r.glasses = randint(rnd, 0, 2)
	// nose options
	r.nose.length = randint(rnd, -1, 1)
	r.nose.width = randint(rnd, -1, 1)
	// mouth options
	r.mouth.width = randint(rnd, -2, 2)
	// cheek options
	r.cheeks.radius = randint(rnd, -4, 4)
	r.cheeks.inward = randint(rnd, 0, 1) == 0
	// ears options

	return
}

type options struct {
	seed  int64
	scale int
	blank bool
}

func newIcon(opt options) (h hamicon) {
	opt.scale = 4
	/*
	  TODO: build all randomizable variables out of options
	*/
	h.Seed = opt.seed
	h.Width = 100 * max(opt.scale, 1)
	h.Height = 100 * max(opt.scale, 1)
	if opt.scale > 0 {
		h.Icon.svg.Transform = optionalSVGAttr(fmt.Sprintf("scale(%d)", opt.scale))
	}
	h.PreserveAspectRatio = "xMidYMid meet"
	h.XMLNS = "http://www.w3.org/2000/svg"
	h.XMLNSSVG = "http://www.w3.org/2000/svg"
	h.Style = struct {
		Type string `xml:"type,attr"`
		CSS  string `xml:",cdata"`
	}{"text/css", basicCSS}
	h.svg.Stroke = "#000"
	h.svg.StrokeWidth = "2"
	h.svg.FillOpacity = "1"
	var r randomizable
	if !opt.blank {
		r = newRandomizable(opt.seed)
	} else {
		r.body.color = "#fff"
	}
	Legs :=
		// TODO prove this stays within the bounds 100x100
		group{svg: svg{ID: "legs", Class: "walk"}, Children: []interface{}{
			path{svg: svg{ID: "bleg1"}}.moveAbs(30, 75+r.legs.y).line(int(float64(10+r.legs.length)*math.Sin(math.Pi/18)), int(float64(13+r.legs.length)*math.Cos(math.Pi/18))),
			path{svg: svg{ID: "bleg2"}}.moveAbs(35, 75+r.legs.y).line(int(float64(10+r.legs.length)*math.Sin(-math.Pi/18)), int(float64(13+r.legs.length)*math.Cos(-math.Pi/18))),
			path{svg: svg{ID: "fleg1"}}.moveAbs(65, 75+r.legs.y).line(int(float64(10+r.legs.length)*math.Sin(math.Pi/18)), int(float64(13+r.legs.length)*math.Cos(math.Pi/18))),
			path{svg: svg{ID: "fleg2"}}.moveAbs(70, 75+r.legs.y).line(int(float64(10+r.legs.length)*math.Sin(-math.Pi/18)), int(float64(13+r.legs.length)*math.Cos(-math.Pi/18))),
		}}
	Body :=
		group{svg: svg{ID: "body", Fill: optionalSVGAttr(r.body.color)}, Children: []interface{}{
			ellipse{CX: 50, CY: 50, RX: 45 + r.body.rx, RY: 30 + r.body.ry},
		}}
	Ears :=
		group{svg: svg{ID: "ears", Class: "twitch", Style: "stroke:#000;stroke-width:1;fill:#fff;"}, Children: []interface{}{
			path{svg: svg{ID: "ear1"}}.moveAbs(53, 28).arc(5, 3, 25, 0, 0, -6, 7).close(),
			path{svg: svg{ID: "ear2"}}.moveAbs(75, 28).arc(5, 3, -25, 0, 1, 6, 7).close(),
		}}
	Eyes :=
		group{svg: svg{ID: "eyes", Style: "stroke:#000;stroke-width:2;fill:#fff;"}, Children: []interface{}{
			ellipse{svg: svg{ID: "eye1"}, CX: 60 - r.eyes.r1 - r.eyes.w/2, CY: 35 + r.eyes.h1, RX: 3 + r.eyes.r1, RY: 3 + r.eyes.r1},
			ellipse{svg: svg{ID: "eye2"}, CX: 70 + r.eyes.r2 + r.eyes.w/2, CY: 35 + r.eyes.h2, RX: 3 + r.eyes.r2, RY: 3 + r.eyes.r1},
		}}
	Nose :=
		group{svg: svg{ID: "nose", Class: "wiggle", Style: "stroke:#000;stroke-width:1;fill:pink;"}, Children: []interface{}{
			path{}.moveAbs(65, 50).line(-(6+r.nose.width)/2, -(5 + r.nose.length)).horiz(6 + 2*(r.nose.width>>1)).close(),
		}}
	Mouth :=
		// TODO split into cheeks and lips and mouth
		group{svg: svg{ID: "mouth", StrokeWidth: "1", FillOpacity: "0"}, Children: []interface{}{
			path{svg: svg{ID: "lip1"}}.moveAbs(65, 50).arc(6+r.mouth.width/2, 5+r.mouth.width/2, 0, 0, 1, 10-r.mouth.width, 0),
			path{svg: svg{ID: "lip2"}}.moveAbs(65, 50).arc(6+r.mouth.width/2, 5+r.mouth.width/2, 0, 0, 0, 10+r.mouth.width, 0),
			ellipse{svg: svg{ID: "speaker", Class: "talk", Style: "fill:#000;fill-opacity:1;"}, CX: 65, CY: 54, RX: 5, RY: 3},
		}}
	cheekDir := 0
	if r.cheeks.inward {
		cheekDir = 1
	}
	Cheeks :=
		/*
			func mouthToSVG(d diffs) (svg string) {
				w := mouthDefaults.width + d.mouth.width
				cr := mouthDefaults.cheekRadius + d.mouth.cheekRadius
				var cd int
				if d.mouth.cheekInward {
					cd = 1
				}
				x := noseDefaults.x + d.nose.x
				y := noseDefaults.y + d.nose.y
				s := "stroke:#000;stroke-width:1;fill-opacity:0;"
				svg += fmt.Sprintf(`<g id="mouth" style="%s">`, s)
				svg += fmt.Sprintf(`<path id="lip1" d="M%d,%d a%d,%d 0 0,1 %d,0"/>`, x, y, w/2+1, w/2, -w)
				svg += fmt.Sprintf(`<path id="lip2" d="M%d,%d a%d,%d 0 0,0 %d,0"/>`, x, y, w/2+1, w/2, w)
				svg += fmt.Sprintf(`<path id="cheek1" class="swell" d="M%d,%d a%d,%d 180 0,%d 0,%d"/>`, x-w-cr-2, y-cr, cr, cr, cd, 2*cr)
				svg += fmt.Sprintf(`<path id="cheek2" class="swell" d="M%d,%d a%d,%d 180 0,%d 0,%d"/>`, x+w+cr+2, y-cr, cr, cr, 1-cd, 2*cr)
				svg += fmt.Sprintf(`<ellipse id="speaker" class="talk" cx="%d" cy="%d" rx="%d" ry="%d" style="fill:#000;fill-opacity:1;"/>`, x, y+w/2-2+1, w/2, w/2-2)
				svg += `</g>`
				return
			}
		*/
		group{svg: svg{ID: "cheeks", StrokeWidth: "1"}, Children: []interface{}{
			path{svg: svg{ID: "cheek1"}}.moveAbs(65-(5+r.cheeks.radius)-(5+r.mouth.width)-2, 65-(5+r.cheeks.radius)).arc(5+r.cheeks.radius, 5+r.cheeks.radius, 0, 0, cheekDir, 0, 2*(5+r.cheeks.radius)),
			path{svg: svg{ID: "cheek2"}}.moveAbs(65+(5+r.cheeks.radius)+(5+r.mouth.width)+2, 65-(5+r.cheeks.radius)).arc(5+r.cheeks.radius, 5+r.cheeks.radius, 0, 0, 1-cheekDir, 0, 2*(5+r.cheeks.radius)),
		}}
	h.Icon.Children = []interface{}{Legs, Body, Ears, Eyes, Nose, Mouth, Cheeks}
	// optional attachments
	if r.glasses > 0 {
		gr := 5 + max(r.eyes.r1, r.eyes.r2)
		gy := (35 + r.eyes.h1 + 35 + r.eyes.h2) / 2
		gx1 := 60 - r.eyes.r1 - r.eyes.w/2
		gx2 := 70 + r.eyes.r1 + r.eyes.w/2
		if gr > (gx2-gx1)/2 {
			gx1, gx2 = gx1-gr+(gx2-gx1)/2, gx2+gr-(gx2-gx1)/2
		}
		glasses := group{svg: svg{ID: "glasses", FillOpacity: "0"}, Children: []interface{}{
			path{}.moveAbs(gx1+gr, gy).horiz(gx2 - gx1 - 2*gr),
			path{}.moveAbs(gx1-gr, gy).line(-8, 6),
			path{}.moveAbs(gx2+gr, gy).line(8, 6),
		}}
		if r.glasses == 1 {
			glasses.Children = append(glasses.Children, []interface{}{
				circle{CX: gx1, CY: gy, R: gr},
				circle{CX: gx2, CY: gy, R: gr},
			})
		}
		if r.glasses == 2 {
			glasses.Children = append(glasses.Children, []interface{}{
				path{}.moveAbs(gx1-gr, gy-gr).vert(2 * gr).horiz(2 * gr).vert(-2 * gr).close(),
				path{}.moveAbs(gx2-gr, gy-gr).vert(2 * gr).horiz(2 * gr).vert(-2 * gr).close(),
			})
		}
		h.Icon.Children = append(h.Icon.Children, glasses)
	}
	return
}
