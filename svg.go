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
	CX      int      `xml:"cx,attr"`
	CY      int      `xml:"cy,attr"`
	RX      int      `xml:"rx,attr"`
	RY      int      `xml:"ry,attr"`
}

type circle struct {
	svg
	XMLName xml.Name `xml:"circle"`
	CX      int      `xml:"cx,attr"`
	CY      int      `xml:"cy,attr"`
	R       int      `xml:"r,attr"`
}

type path struct {
	svg
	XMLName xml.Name `xml:"path"`
	D       string   `xml:"d,attr"`
}

func (p path) moveAbs(x, y int) path {
	p.D += fmt.Sprintf("M%d,%d", x, y)
	return p
}
func (p path) line(x, y int) path {
	p.D += fmt.Sprintf("l%d,%d", x, y)
	return p
}
func (p path) lineAbs(x, y int) path {
	p.D += fmt.Sprintf("L%d,%d", x, y)
	return p
}
func (p path) vert(y int) path {
	p.D += fmt.Sprintf("v%d", y)
	return p
}
func (p path) horiz(x int) path {
	p.D += fmt.Sprintf("h%d", x)
	return p
}
func (p path) arc(rx, ry, xAxisRotate, largeArcFlag, sweepFlag, x, y int) path {
	p.D += fmt.Sprintf("a%d,%d %d %d,%d %d,%d", rx, ry, xAxisRotate, largeArcFlag, sweepFlag, x, y)
	return p
}
func (p path) arcAbs(rx, ry, xAxisRotate, largeArcFlag, sweepFlag, x, y int) path {
	p.D += fmt.Sprintf("A%d,%d %d %d,%d %d,%d", rx, ry, xAxisRotate, largeArcFlag, sweepFlag, x, y)
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
	// TODO: remove unnec whitespace?
	return s
}
