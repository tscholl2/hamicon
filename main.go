package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
)

func blank(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	enc := xml.NewEncoder(buf)
	enc.Indent("", "  ")
	v := newIcon(options{seed: 1, scale: 4, blank: true})
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Fprint(w, minimize(string(buf.Bytes())))
}

func get(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	enc := xml.NewEncoder(buf)
	enc.Indent("", "  ")
	v := newIcon(options{seed: 1, scale: 4})
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Fprint(w, minimize(string(buf.Bytes())))
	/*
		enc := xml.NewEncoder(w)
		enc.Indent("", "")
		v := newIcon(options{seed: rand.Int63()})
		if err := enc.Encode(v); err != nil {
			fmt.Fprintf(w, "error: %v\n", err)
		}
	*/
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
