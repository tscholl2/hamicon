package main

import (
	"encoding/xml"
	"fmt"
)

type svg struct {
	ID      optionalSVGAttr `xml:"id,attr"`
	Class   optionalSVGAttr `xml:"class,attr"`
	Comment optionalSVGAttr `xml:",comment"`
	Style   optionalSVGAttr `xml:"style,attr"`
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

// groups

// TODO: don't use?
type group struct {
	svg
	XMLName  xml.Name `xml:"g"`
	Children []interface{}
}

// ellipses

type ellipse struct {
	svg
	XMLName xml.Name `xml:"ellipse"`
	CX      int      `xml:"cx,attr"`
	CY      int      `xml:"cy,attr"`
	RX      int      `xml:"rx,attr"`
	RY      int      `xml:"ry,attr"`
}

// paths

type path struct {
	svg
	XMLName xml.Name `xml:"path"`
	D       string   `xml:"d,attr"`
}

// path directions

func (p *path) moveAbs(x, y int) *path {
	p.D += fmt.Sprintf("M%d,%d", x, y)
	return p
}
func (p *path) line(x, y int) *path {
	p.D += fmt.Sprintf("l%d,%d", x, y)
	return p
}
func (p *path) lineAbs(x, y int) *path {
	p.D += fmt.Sprintf("L%d,%d", x, y)
	return p
}
func (p *path) arc(rx, ry, xAxisRotate, largeArcFlag, sweepFlag, x, y int) *path {
	p.D += fmt.Sprintf("a%d,%d %d %d,%d %d,%d", rx, ry, xAxisRotate, largeArcFlag, sweepFlag, x, y)
	return p
}
func (p *path) arcAbs(rx, ry, xAxisRotate, largeArcFlag, sweepFlag, x, y int) *path {
	p.D += fmt.Sprintf("A%d,%d %d %d,%d %d,%d", rx, ry, xAxisRotate, largeArcFlag, sweepFlag, x, y)
	return p
}
func (p *path) close() *path {
	p.D += "z"
	return p
}
