package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
)

//go:generate embd -n basicSVG static/basic.1.svg

// <path id="lip" d="M50,50 a 30,60 30 0,0 30,0"/>
type lip struct {
	Angle int    // 30
	Width int    // 30
	Style string // ""
}

type values struct {
	Lip string
}

var (
	tmpl = template.Must(template.New("basic").Parse(basicSVG))
)

func newHam() (v values) {
	// lip
	angle := rand.Intn(360) - 180
	width := rand.Intn(20) + 15
	frown := rand.Intn(2)
	style := ""
	v.Lip = `<path id = "lip"`
	if style != "" {
		v.Lip += ` style="` + style + `"`
	}
	var y int
	if frown > 0 {
		y = 55
	} else {
		y = 50
	}
	v.Lip += fmt.Sprintf(` d="M%d,%d a %d,60 %d 0,%d %d,0"/>`, 65-width/2, y, width, angle, frown, width)
	return v
}

func get(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, newHam())
}

func headers(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		h.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/", get)
	http.ListenAndServe(":8080", headers(http.DefaultServeMux))
}
