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

type pathDirection struct {
	Moves []interface {
		String() string
	}
}

func (d pathDirection) MarshalXMLAttr(name xml.Name) (a xml.Attr, e error) {
	a.Name = name
	var D string
	for _, m := range d.Moves {
		D += " " + m.String()
	}
	a.Value = D
	return
}

type path struct {
	svg
	XMLName xml.Name      `xml:"ellipse"`
	D       pathDirection `xml:"d,attr"`
}

// path directions

type movePath struct {
	absolute bool
	x, y     int
}

func (m movePath) String() string {
	M := "m"
	if m.absolute {
		M = "M"
	}
	return fmt.Sprintf("%s%d,%d", M, m.x, m.y)
}

type linePath struct {
	absolute bool
	x, y     int
}

func (l linePath) String() string {
	L := "l"
	if l.absolute {
		L = "L"
	}
	return fmt.Sprintf("%s%d,%d", L, l.x, l.y)
}

type arcPath struct {
	absolute                                           bool
	rx, ry, xAxisRotate, largeArcFlag, sweepFlag, x, y int
}

func (a arcPath) String() string {
	A := "a"
	if a.absolute {
		A = "A"
	}
	return fmt.Sprintf("%s%d,%d %d %d,%d %d,%d",
		A, a.rx, a.ry, a.xAxisRotate, a.largeArcFlag, a.sweepFlag, a.x, a.y)
}

type closePath struct{}

func (z closePath) String() string {
	return "z"
}
