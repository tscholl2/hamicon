package main

import (
	"encoding/xml"
	"fmt"
	"math"
	"math/rand"
)

//go:generate embd -n basicCSS static/hamicon.css

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

	legsLength, legsY          float64
	leg1X, leg2X, leg3X, leg4X float64

	eye1R, eye1CX, eye1CY float64
	eye2R, eye2CX, eye2CY float64
	eyesW                 float64

	noseX, noseY          float64
	noseLength, noseWidth float64

	mouthWidth float64

	cheeksRadius, cheeksInward float64

	glasses int // none, round, square
	hat     int // none, top
}

func defaults() wigglable {
	return wigglable{
		bodyCX: 50, bodyCY: 50, bodyRX: 45, bodyRY: 30,
		bodyColor: "#fff",

		legsLength: 10, legsY: 75,
		leg1X: 30, leg2X: 35, leg3X: 65, leg4X: 70,

		eye1R: 3, eye1CX: 60, eye1CY: 35,
		eye2R: 3, eye2CX: 70, eye2CY: 35,
		eyesW: 0,

		noseX: 65, noseY: 50,
		noseLength: 4, noseWidth: 6,

		mouthWidth: 8,

		cheeksInward: 0, cheeksRadius: 5,

		glasses: 0,
		hat:     0,
	}
}

func newRandomizable(seed int64) wigglable {
	rnd := rand.New(rand.NewSource(seed))
	w := defaults()
	w.bodyRX += randintf(rnd, -3, 0)
	w.bodyRY += randintf(rnd, -5, 5)
	w.bodyColor = optionalSVGAttr(randcolor(rnd))

	w.legsLength += randintf(rnd, -1, 6)
	if w.bodyRY > 0 && defaults().legsLength < w.legsLength {
		w.legsLength = defaults().legsLength + 2
	}

	if w.bodyRY < defaults().bodyRY {
		w.legsY -= 3
	}

	w.eye1R += randintf(rnd, 0, 2)
	w.eye1CY += randintf(rnd, -1, 1)
	w.eye2R += randintf(rnd, 0, 2)
	w.eye2CY += randintf(rnd, -1, 1)
	w.eyesW += randintf(rnd, -1, 5)

	w.noseLength += randintf(rnd, -1, 1)
	w.noseWidth += randintf(rnd, -1, 1)

	w.mouthWidth += randintf(rnd, -2, 2)

	w.cheeksRadius += randintf(rnd, -4, 4)
	w.cheeksInward += randintf(rnd, 0, 1)

	w.glasses = randint(rnd, 0, 2)
	w.hat = randint(rnd, 0, 1)

	return w
}

type options struct {
	seed       int64
	scale      int
	blank      bool
	classes    bool
	includeCSS bool
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

	if opt.includeCSS {
		h.Style = struct {
			Type string `xml:"type,attr"`
			CSS  string `xml:",cdata"`
		}{"text/css", basicCSS}
	}

	h.Icon.svg.Stroke = "#000"
	h.Icon.svg.StrokeWidth = "2"
	h.Icon.svg.FillOpacity = "1"
	w := defaults()
	if !opt.blank {
		w = newRandomizable(opt.seed)
	}
	Legs :=
		// TODO prove this stays within the bounds 100x100
		group{svg: svg{ID: "legs", StrokeLinecap: "round"}, Children: []interface{}{
			path{svg: svg{ID: "bleg1"}}.moveAbs(w.leg1X, w.legsY).line(w.legsLength*math.Sin(-math.Pi/18), w.legsLength*math.Cos(math.Pi/18)),
			path{svg: svg{ID: "bleg2"}}.moveAbs(w.leg2X, w.legsY).line(w.legsLength*math.Sin(math.Pi/18), w.legsLength*math.Cos(-math.Pi/18)),
			path{svg: svg{ID: "fleg1"}}.moveAbs(w.leg3X, w.legsY).line(w.legsLength*math.Sin(-math.Pi/18), w.legsLength*math.Cos(math.Pi/18)),
			path{svg: svg{ID: "fleg2"}}.moveAbs(w.leg4X, w.legsY).line(w.legsLength*math.Sin(math.Pi/18), w.legsLength*math.Cos(-math.Pi/18)),
		}}
	Body :=
		group{svg: svg{ID: "body", Fill: w.bodyColor}, Children: []interface{}{
			ellipse{CX: w.bodyCX, CY: w.bodyCY, RX: w.bodyRX, RY: w.bodyRY},
		}}
	Ears :=
		group{svg: svg{ID: "ears", Fill: w.bodyColor, StrokeWidth: "1"}, Children: []interface{}{
			path{svg: svg{ID: "ear1"}}.moveAbs(53, 28).arc(5, 3, 25, 0, 0, -6, 7).close(),
			path{svg: svg{ID: "ear2"}}.moveAbs(75, 28).arc(5, 3, -25, 0, 1, 6, 7).close(),
		}}
	Eyes :=
		group{svg: svg{ID: "eyes", Fill: "#fff", StrokeWidth: "1"}, Children: []interface{}{
			ellipse{svg: svg{ID: "eye1"}, CX: w.eye1CX - w.eyesW/2, CY: w.eye1CY, RX: w.eye1R, RY: w.eye1R},
			ellipse{svg: svg{ID: "eye2"}, CX: w.eye2CX + w.eyesW/2, CY: w.eye2CY, RX: w.eye2R, RY: w.eye2R},
		}}
	Nose :=
		group{svg: svg{ID: "nose", Fill: "pink", StrokeWidth: "1"}, Children: []interface{}{
			path{}.moveAbs(65, 50).line(-w.noseWidth/2, -w.noseLength).horiz(w.noseWidth).close(),
		}}
	Mouth :=
		group{svg: svg{ID: "mouth", StrokeWidth: "1", FillOpacity: "0"}, Children: []interface{}{
			path{svg: svg{ID: "lip1"}}.moveAbs(w.noseX, w.noseY).arc(w.mouthWidth/2, w.mouthWidth/2-2, 0, 0, 1, -w.mouthWidth, 0),
			path{svg: svg{ID: "lip2"}}.moveAbs(w.noseX, w.noseY).arc(w.mouthWidth/2, w.mouthWidth/2-2, 0, 0, 0, w.mouthWidth, 0),
			ellipse{svg: svg{ID: "speaker", Style: "display:none;", Fill: "#000", FillOpacity: "1"}, CX: w.noseX, CY: w.noseY + 3, RX: 5, RY: 3},
		}}
	Cheeks :=
		group{svg: svg{ID: "cheeks", StrokeWidth: "1", FillOpacity: "0"}, Children: []interface{}{
			path{svg: svg{ID: "cheek1"}}.moveAbs(w.noseX-w.cheeksRadius-w.mouthWidth-2, w.noseY-w.cheeksRadius).arc(w.cheeksRadius, w.cheeksRadius, 0, 0, w.cheeksInward, 0, 2*w.cheeksRadius),
			path{svg: svg{ID: "cheek2"}}.moveAbs(w.noseX+w.cheeksRadius+w.mouthWidth+2, w.noseY-w.cheeksRadius).arc(w.cheeksRadius, w.cheeksRadius, 0, 0, 1-w.cheeksInward, 0, 2*w.cheeksRadius),
		}}
	if opt.classes {
		Legs.Class = "walk"
		Ears.Class = "twitch"
		Eyes.Class = "blink"
		Nose.Class = "wiggle"
		Mouth.Class = "talk"
		Cheeks.Class = "swell"
	}
	h.Icon.Children = []interface{}{Legs, Body, Ears, Eyes, Nose, Mouth, Cheeks}
	// optional attachments
	if w.glasses > 0 {
		gr := 2 + math.Max(w.eye1R, w.eye2R)
		gy := (w.eye1CY + w.eye2CY) / 2
		gx1 := w.eye1CX - w.eyesW
		gx2 := w.eye2CX + w.eyesW
		if gr > (gx2-gx1)/2 {
			gx1, gx2 = gx1-gr+(gx2-gx1)/2, gx2+gr-(gx2-gx1)/2
		}
		glasses := group{svg: svg{ID: "glasses", FillOpacity: "0"}, Children: []interface{}{
			path{}.moveAbs(gx1+gr, gy).horiz(gx2 - gx1 - 2*gr),
			path{}.moveAbs(gx1-gr, gy).line(-8, 5),
			path{}.moveAbs(gx2+gr, gy).line(8, 5),
		}}
		if w.glasses == 1 {
			glasses.Children = append(glasses.Children, []interface{}{
				circle{CX: gx1, CY: gy, R: gr},
				circle{CX: gx2, CY: gy, R: gr},
			})
		}
		if w.glasses == 2 {
			glasses.Children = append(glasses.Children, []interface{}{
				path{}.moveAbs(gx1-gr, gy-gr).vert(2 * gr).horiz(2 * gr).vert(-2 * gr).close(),
				path{}.moveAbs(gx2-gr, gy-gr).vert(2 * gr).horiz(2 * gr).vert(-2 * gr).close(),
			})
		}
		h.Icon.Children = append(h.Icon.Children, glasses)
	}
	if w.hat > 0 {
		hat := group{svg: svg{ID: "hat", Class: "tip"}}
		if w.hat == 1 {
			hat.Children = append(hat.Children, []interface{}{
				group{svg: svg{Fill: "#000", Stroke: "#999", StrokeWidth: "1"}, Children: []interface{}{
					rect{X: 53, Y: 3, Height: 11, Width: 18},
					rect{X: 46, Y: 14, Height: 7, Width: 31},
				}},
			})
		}
		h.Icon.Children = append(h.Icon.Children, hat)
	}
	return
}
