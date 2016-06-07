package main

import (
	"encoding/xml"
	"fmt"
	"math"
	"math/rand"
)

//go:generate embd -n basicCSS static/basic.1.css

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

type wigglable struct {
	bodyCX, bodyCY, bodyRX, bodyRY float64
	bodyColor                      optionalSVGAttr

	legsLength, legsY float64

	eye1R, eye1CX, eye1CY float64
	eye2R, eye2CX, eye2CY float64
	eyesW                 float64

	glasses int // none, round, square

	noseX, noseY          float64
	noseLength, noseWidth float64

	mouthWidth float64

	cheeksRadius, cheeksInward float64
}

func defaults() wigglable {
	return wigglable{
		bodyCX: 50, bodyCY: 50, bodyRX: 45, bodyRY: 30,
		bodyColor: "#fff",

		legsLength: 10, legsY: 75,

		eye1R: 3, eye1CX: 60, eye1CY: 35,
		eye2R: 3, eye2CX: 70, eye2CY: 35,
		eyesW: 0,

		glasses: 0,

		noseX: 65, noseY: 50,
		noseLength: 4, noseWidth: 6,

		mouthWidth: 8,

		cheeksInward: 0, cheeksRadius: 5,
	}
}

func newRandomizable(seed int64) wigglable {
	rnd := rand.New(rand.NewSource(seed))
	r := defaults()
	r.bodyRX += randintf(rnd, -3, 0)
	/*
		// body options
		r.body.rx = randintf(rnd, -3, 0)
		r.body.ry = randintf(rnd, -5, 5)
		r.body.color = randcolor(rnd)
		// leg options
		r.legs.length = randintf(rnd, -2, 5)
		if r.body.ry > 0 && r.legs.length <= 0 {
			r.legs.length = 2
		}
		if r.body.ry < 0 {
			r.legs.y = -3
		}
		// eyes options
		r.eyes.r1 = randintf(rnd, 0, 2)
		r.eyes.h1 = randintf(rnd, -1, 1)
		r.eyes.r2 = randintf(rnd, 0, 2)
		r.eyes.h2 = randintf(rnd, -1, 1)
		r.eyes.w = randintf(rnd, -2, 2)
		// glasses options
		r.glasses = randint(rnd, 0, 2)
		// nose options
		r.nose.length = randintf(rnd, -1, 1)
		r.nose.width = randintf(rnd, -1, 1)
		// mouth options
		r.mouth.width = randintf(rnd, -2, 2)
		// cheek options
		r.cheeks.radius = randintf(rnd, -4, 4)
		r.cheeks.inward = randintf(rnd, 0, 1)
		// ears options
	*/
	return r
}

type options struct {
	seed  int64
	scale int
	blank bool
}

func newIcon(opt options) (h hamicon) {
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
	w := defaults()
	if !opt.blank {
		w = newRandomizable(opt.seed)
	}
	Legs :=
		// TODO prove this stays within the bounds 100x100
		group{svg: svg{ID: "legs", Class: "walk", StrokeLinecap: "round"}, Children: []interface{}{
			path{svg: svg{ID: "bleg1"}}.moveAbs(30, w.legsY).line(w.legsLength*math.Sin(-math.Pi/18), w.legsLength*math.Cos(math.Pi/18)),
			path{svg: svg{ID: "bleg2"}}.moveAbs(35, w.legsY).line(w.legsLength*math.Sin(math.Pi/18), w.legsLength*math.Cos(-math.Pi/18)),
			path{svg: svg{ID: "fleg1"}}.moveAbs(65, w.legsY).line(w.legsLength*math.Sin(-math.Pi/18), w.legsLength*math.Cos(math.Pi/18)),
			path{svg: svg{ID: "fleg2"}}.moveAbs(70, w.legsY).line(w.legsLength*math.Sin(math.Pi/18), w.legsLength*math.Cos(-math.Pi/18)),
		}}
	Body :=
		group{svg: svg{ID: "body", Fill: w.bodyColor}, Children: []interface{}{
			ellipse{CX: w.bodyCX, CY: w.bodyCY, RX: w.bodyRX, RY: w.bodyRY},
		}}
	Ears :=
		group{svg: svg{ID: "ears", Class: "twitch", Fill: w.bodyColor, StrokeWidth: "1"}, Children: []interface{}{
			path{svg: svg{ID: "ear1"}}.moveAbs(53, 28).arc(5, 3, 25, 0, 0, -6, 7).close(),
			path{svg: svg{ID: "ear2"}}.moveAbs(75, 28).arc(5, 3, -25, 0, 1, 6, 7).close(),
		}}
	Eyes :=
		group{svg: svg{ID: "eyes", Class: "blink", Fill: "#fff", StrokeWidth: "1"}, Children: []interface{}{
			ellipse{svg: svg{ID: "eye1"}, CX: w.eye1CX, CY: w.eye1CY, RX: w.eye1R, RY: w.eye1R},
			ellipse{svg: svg{ID: "eye2"}, CX: w.eye2CX, CY: w.eye2CY, RX: w.eye2R, RY: w.eye2R},
		}}
	Nose :=
		group{svg: svg{ID: "nose", Class: "wiggle", Fill: "pink", StrokeWidth: "1"}, Children: []interface{}{
			path{}.moveAbs(65, 50).line(-w.noseWidth/2, -w.noseLength).horiz(w.noseWidth).close(),
		}}
	Mouth :=
		group{svg: svg{ID: "mouth", StrokeWidth: "1", FillOpacity: "0"}, Children: []interface{}{
			path{svg: svg{ID: "lip1"}}.moveAbs(w.noseX, w.noseY).arc(w.mouthWidth/2, w.mouthWidth/2, 0, 0, 1, -w.mouthWidth, 0),
			path{svg: svg{ID: "lip2"}}.moveAbs(w.noseX, w.noseY).arc(w.mouthWidth/2, w.mouthWidth/2, 0, 0, 0, w.mouthWidth, 0),
			ellipse{svg: svg{ID: "speaker", Class: "talk", Fill: "#000", FillOpacity: "1"}, CX: 65, CY: 54, RX: 5, RY: 3},
		}}
	Cheeks :=
		group{svg: svg{ID: "cheeks", Class: "swell", StrokeWidth: "1", FillOpacity: "0"}, Children: []interface{}{
			path{svg: svg{ID: "cheek1"}}.moveAbs(w.noseX-w.cheeksRadius-w.mouthWidth-2, w.noseY-w.cheeksRadius).arc(w.cheeksRadius, w.cheeksRadius, 0, 0, w.cheeksInward, 0, 2*w.cheeksRadius),
			path{svg: svg{ID: "cheek2"}}.moveAbs(w.noseX+w.cheeksRadius+w.mouthWidth+2, w.noseY-w.cheeksRadius).arc(w.cheeksRadius, w.cheeksRadius, 0, 0, 1-w.cheeksInward, 0, 2*w.cheeksRadius),
		}}
	h.Icon.Children = []interface{}{Legs, Body, Ears, Eyes, Nose, Mouth, Cheeks}
	// optional attachments
	/*
		if r.glasses > 0 {
			gr := 5 + math.Max(r.eyes.r1, r.eyes.r2)
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
	*/
	return
}
