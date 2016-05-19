package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
)

//go:generate embd -n basicSVG static/basic.1.svg

var (
	tmpl = template.Must(template.New("basic").Parse(basicSVG))
)

func get(w http.ResponseWriter, r *http.Request) {
	rnd := rand.New(rand.NewSource(rand.Int63()))
	tmpl.Execute(w, newDiffs(rnd).toMap())
}

func blank(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, diffs{}.toMap())
}

func headers(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		h.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/", get)
	http.HandleFunc("/blank", blank)
	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", headers(http.DefaultServeMux))
}
