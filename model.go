package main

import "encoding/xml"
import "strconv"

/*

<svg width="100" height="100" preserveAspectRatio="xMidYMid meet" xmlns="http://www.w3.org/2000/svg" xmlns:svg="http://www.w3.org/2000/svg">
<style type="text/css" >
<![CDATA[
  ... css ...
]]>
</style>
<g id="icon">
  <g id="legs" class="walk" style="stroke:#000;stroke-width:2;stroke-linecap:round;">
    <path id="bleg1" d="M30,75 l0,10" transform="rotate(15,30,75)"/>
    <path id="bleg2" d="M35,75 l0,10" transform="rotate(-15,35,75)"/>
    <path id="fleg1" d="M65,75 l0,10" transform="rotate(15,65,75)"/>
    <path id="fleg2" d="M70,75 l0,10" transform="rotate(-15,70,75)"/>
  </g>
  <g style="fill-opacity:1;stroke:#000;stroke-width:2;fill:#fff;">
    <ellipse id="body" cx="50" cy="50" rx="45" ry="30"/>
  </g>
  <g id="ears" class="twitc" style="stroke:#000;stroke-width:1;fill:#fff;">
    <path id="ear1" d="M53,28 a5,3 25 0,0 -6,7 z"/>
    <path id="ear2" d="M75,28 a5,3 -25 0,1 6,7 z"/>
  </g>
  <g id="eyes" style="stroke:#000;stroke-width:2;fill:#fff;">
    <ellipse class="blink" id="eye1" cx="60" cy="35" rx="3" ry="3"/>
    <ellipse class="blink" id="eye2" cx="70" cy="35" rx="3" ry="3"/>
  </g>
  <g id="nose" class="wiggle" style="stroke:#000;stroke-width:1;fill:pink;">
    <path d="M65,50 l-2,-5 l5,0 Z"/>
  </g>
  <g id="mouth" style="stroke:#000;stroke-width:1;fill-opacity:0;">
    <path id="lip1" d="M65,50 a6,5 0 0,1 -10,0"/>
    <path id="lip2" d="M65,50 a6,5 0 0,0 10,0"/>
    <path id="cheek1" class="swell" d="M48,45 a5,5 180 0,0 0,10"/>
    <path id="cheek2" class="swell" d="M82,45 a5,5 180 0,1 0,10"/>
    <ellipse id="speaker" class="talk" cx="65" cy="54" rx="5" ry="3" style="fill:#000;fill-opacity:1;"/>
  </g>
</g>
*/

type hamicon struct {
	XMLName             xml.Name `xml:"svg"`
	Width               int      `xml:"width,attr"`
	Height              int      `xml:"height,attr"`
	PreserveAspectRatio string   `xml:"preserveAspectRatio,attr"`
	XMLNS               string   `xml:"xmlns,attr"`
	XMLNSSVG            string   `xml:"xmlns:svg,attr"`
	Icon                group
}

type options struct {
	seed int64
}

func newIcon(opt options) (h hamicon) {
	/*
	  TODO: build all randomizable variables out of options
	*/
	h.Icon.Comment = optionalSVGAttr(strconv.FormatInt(opt.seed, 16))
	h.Width = 100
	h.Height = 100
	h.PreserveAspectRatio = "xMidYMid meet"
	h.XMLNS = "http://www.w3.org/2000/svg"
	h.XMLNSSVG = "http://www.w3.org/2000/svg"
	h.Icon.Children = []interface{}{
		group{svg: svg{ID: "legs"}, Children: nil},
		group{svg: svg{ID: "body"}, Children: nil},
		group{svg: svg{ID: "ears"}, Children: nil},
		group{svg: svg{ID: "eyes"}, Children: nil},
		group{svg: svg{ID: "nose"}, Children: nil},
		group{svg: svg{ID: "mouth"}, Children: nil},
	}
	return
}
