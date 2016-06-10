package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type svg struct {
	ID              optionalSVGAttr `xml:"id,attr"`
	Class           optionalSVGAttr `xml:"class,attr"`
	Comment         optionalSVGAttr `xml:",comment"`
	Fill            optionalSVGAttr `xml:"fill,attr"`
	FillOpacity     optionalSVGAttr `xml:"fill-opacity,attr"`
	Stroke          optionalSVGAttr `xml:"stroke,attr"`
	StrokeLinecap   optionalSVGAttr `xml:"stroke-linecap,attr"`
	StrokeWidth     optionalSVGAttr `xml:"stroke-width,attr"`
	Style           optionalSVGAttr `xml:"style,attr"`
	Transform       optionalSVGAttr `xml:"transform,attr"`
	TransformOrigin optionalSVGAttr `xml:"transform-origin,attr"`
}

// TODO: don't use?
type group struct {
	svg
	XMLName  xml.Name `xml:"g"`
	Children []interface{}
}

type ellipse struct {
	svg
	XMLName xml.Name `xml:"ellipse"`
	CX      float64  `xml:"cx,attr"`
	CY      float64  `xml:"cy,attr"`
	RX      float64  `xml:"rx,attr"`
	RY      float64  `xml:"ry,attr"`
}

type rect struct {
	svg
	XMLName xml.Name `xml:"rect"`
	X       float64  `xml:"x,attr"`
	Y       float64  `xml:"y,attr"`
	Height  float64  `xml:"height,attr"`
	Width   float64  `xml:"width,attr"`
}

type circle struct {
	svg
	XMLName xml.Name `xml:"circle"`
	CX      float64  `xml:"cx,attr"`
	CY      float64  `xml:"cy,attr"`
	R       float64  `xml:"r,attr"`
}

type path struct {
	svg
	XMLName xml.Name `xml:"path"`
	D       string   `xml:"d,attr"`
}

func (p path) moveAbs(x, y float64) path {
	p.D += fmt.Sprintf("M%.1f,%.1f", x, y)
	return p
}
func (p path) line(x, y float64) path {
	p.D += fmt.Sprintf("l%.1f,%.1f", x, y)
	return p
}
func (p path) lineAbs(x, y float64) path {
	p.D += fmt.Sprintf("L%.1f,%.1f", x, y)
	return p
}
func (p path) vert(y float64) path {
	p.D += fmt.Sprintf("v%.1f", y)
	return p
}
func (p path) horiz(x float64) path {
	p.D += fmt.Sprintf("h%.1f", x)
	return p
}
func (p path) arc(rx, ry, xAxisRotate, largeArcFlag, sweepFlag, x, y float64) path {
	p.D += fmt.Sprintf("a%.1f,%.1f %.1f %.1f,%.1f %.1f,%.1f", rx, ry, xAxisRotate, largeArcFlag, sweepFlag, x, y)
	return p
}
func (p path) arcAbs(rx, ry, xAxisRotate, largeArcFlag, sweepFlag, x, y float64) path {
	p.D += fmt.Sprintf("A%.1f,%.1f %.1f %.1f,%.1f %.1f,%.1f", rx, ry, xAxisRotate, largeArcFlag, sweepFlag, x, y)
	return p
}
func (p path) close() path {
	p.D += "z"
	return p
}

type optionalSVGAttr string

func (opt optionalSVGAttr) MarshalXMLAttr(name xml.Name) (a xml.Attr, e error) {
	if opt == "" {
		return xml.Attr{}, nil
	}
	a.Name = name
	a.Value = string(opt)
	return
}

func minimize(s string) string {
	s = strings.Replace(s, "></path>", "/>", -1)
	s = strings.Replace(s, "></ellipse>", "/>", -1)
	s = strings.Replace(s, "></circle>", "/>", -1)
	s = strings.Replace(s, "></rect>", "/>", -1)
	s = strings.Replace(s, ".0", "", -1)
	// TODO: remove unnec whitespace?
	return s
}
