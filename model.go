package main

import "encoding/xml"

/*

<svg width="100" height="100" preserveAspectRatio="xMidYMid meet" xmlns="http://www.w3.org/2000/svg" xmlns:svg="http://www.w3.org/2000/svg">
<style type="text/css" >
<![CDATA[
  ... css ...
]]>
</style>
*/

type hamicon struct {
	XMLName             xml.Name    `xml:"svg"`
	Width               int         `xml:"width,attr"`
	Height              int         `xml:"height,attr"`
	PreserveAspectRatio string      `xml:"preserveAspectRatio,attr"`
	XMLNS               string      `xml:"xmlns,attr"`
	XMLNSSVG            string      `xml:"xmlns:svg,attr"`
	Seed                int64       `xml:"seed,attr"`
	Style               interface{} `xml:"style"`
	Children            []interface{}
}

type options struct {
	seed int64
}

func newIcon(opt options) (h hamicon) {
	/*
	  TODO: build all randomizable variables out of options
	*/
	h.Seed = opt.seed
	h.Width = 100
	h.Height = 100
	h.PreserveAspectRatio = "xMidYMid meet"
	h.XMLNS = "http://www.w3.org/2000/svg"
	h.XMLNSSVG = "http://www.w3.org/2000/svg"
	h.Style = struct {
		Type string `xml:"type,attr"`
		CSS  string `xml:",cdata"`
	}{"text/css", basicCSS}
	Legs :=
		/*
		  <g id="legs" class="walk" style="stroke:#000;stroke-width:2;stroke-linecap:round;">
		    <path id="bleg1" d="M30,75 l0,10"/>
		    <path id="bleg2" d="M35,75 l0,10"/>
		    <path id="fleg1" d="M65,75 l0,10"/>
		    <path id="fleg2" d="M70,75 l0,10"/>
		  </g>
		*/
		group{svg: svg{ID: "legs", Style: "stroke:#000;stroke-width:2;stroke-linecap:round;"}, Children: []interface{}{
			path{svg: svg{ID: "bleg1"}}.moveAbs(30, 75).vert(10),
			path{svg: svg{ID: "bleg2"}}.moveAbs(35, 75).vert(10),
			path{svg: svg{ID: "fleg1"}}.moveAbs(65, 75).vert(10),
			path{svg: svg{ID: "fleg2"}}.moveAbs(70, 75).vert(10),
		}}
	Body :=
		/*
		   <g style="fill-opacity:1;stroke:#000;stroke-width:2;fill:#fff;">
		     <ellipse id="body" cx="50" cy="50" rx="45" ry="30"/>
		   </g>
		*/
		group{svg: svg{ID: "body", Style: "fill-opacity:1;stroke:#000;stroke-width:2;fill:#fff;"}, Children: []interface{}{
			ellipse{CX: 50, CY: 50, RX: 45, RY: 30},
		}}
	Ears :=
		/*
		   <g id="ears" class="twitch" style="stroke:#000;stroke-width:1;fill:#fff;">
		     <path id="ear1" d="M53,28 a5,3 25 0,0 -6,7 z"/>
		     <path id="ear2" d="M75,28 a5,3 -25 0,1 6,7 z"/>
		   </g>
		*/
		group{svg: svg{ID: "ears", Class: "twitch", Style: "stroke:#000;stroke-width:1;fill:#fff;"}, Children: []interface{}{
			path{svg: svg{ID: "ear1"}}.moveAbs(53, 28).arc(5, 3, 25, 0, 0, -6, 7).close(),
			path{svg: svg{ID: "ear2"}}.moveAbs(75, 28).arc(5, 3, -25, 0, 1, 6, 7).close(),
		}}
	Eyes :=
		/*
		   <g id="eyes" style="stroke:#000;stroke-width:2;fill:#fff;">
		     <ellipse class="blink" id="eye1" cx="60" cy="35" rx="3" ry="3"/>
		     <ellipse class="blink" id="eye2" cx="70" cy="35" rx="3" ry="3"/>
		   </g>
		*/
		group{svg: svg{ID: "eyes", Style: "stroke:#000;stroke-width:2;fill:#fff;"}, Children: []interface{}{
			ellipse{svg: svg{ID: "eye1"}, CX: 60, CY: 35, RX: 3, RY: 3},
			ellipse{svg: svg{ID: "eye2"}, CX: 70, CY: 35, RX: 3, RY: 3},
		}}
	Nose :=
		/*
		   <g id="nose" class="wiggle" style="stroke:#000;stroke-width:1;fill:pink;">
		     <path d="M65,50 l-2,-5 l5,0 Z"/>
		   </g>
		*/
		group{svg: svg{ID: "nose", Class: "wiggle", Style: "stroke:#000;stroke-width:1;fill:pink;"}, Children: []interface{}{
			path{}.moveAbs(65, 50).line(-2, -5).line(5, 0).close(),
		}}
	Mouth :=
		/*
		   <g id="mouth" style="stroke:#000;stroke-width:1;fill-opacity:0;">
		     <path id="lip1" d="M65,50 a6,5 0 0,1 -10,0"/>
		     <path id="lip2" d="M65,50 a6,5 0 0,0 10,0"/>
		     <path id="cheek1" class="swell" d="M48,45 a5,5 180 0,0 0,10"/>
		     <path id="cheek2" class="swell" d="M82,45 a5,5 180 0,1 0,10"/>
		     <ellipse id="speaker" class="talk" cx="65" cy="54" rx="5" ry="3" style="fill:#000;fill-opacity:1;"/>
		   </g>
		*/
		group{svg: svg{ID: "mouth", Style: "stroke:#000;stroke-width:1;fill-opacity:0;"}, Children: []interface{}{
			path{svg: svg{ID: "lip1"}}.moveAbs(65, 50).arc(6, 5, 0, 0, 1, -10, 0),
			path{svg: svg{ID: "lip2"}}.moveAbs(65, 50).arc(6, 5, 0, 0, 0, 10, 0),
			path{svg: svg{ID: "cheek1"}}.moveAbs(45, 45).arc(5, 5, 0, 0, 0, 0, 10),
			path{svg: svg{ID: "cheek2"}}.moveAbs(82, 45).arc(5, 5, 0, 0, 1, 0, 10),
			ellipse{svg: svg{ID: "speaker", Class: "talk", Style: "fill:#000;fill-opacity:1;"}, CX: 65, CY: 54, RX: 5, RY: 3},
		}}
	h.Children = []interface{}{Legs, Body, Ears, Eyes, Nose, Mouth}
	return
}
