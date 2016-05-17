package main

import (
	"net/http"
	"text/template"
)

//go:generate embd -n basicSVG static/basic.1.svg

var (
	tmpl = template.Must(template.New("basic").Parse(basicSVG))
)

func get(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, newDiffs().toMap())
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
