package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
)

func blank(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func headers(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		h.ServeHTTP(w, r)
	})
}

func main() {
	buf := &bytes.Buffer{}
	enc := xml.NewEncoder(buf)
	enc.Indent("", "  ")
	v := newIcon(options{seed: 1, blank: true})
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Fprint(os.Stdout, minimize(string(buf.Bytes())))
	/*
		http.HandleFunc("/", get)
		http.HandleFunc("/blank", blank)
		fmt.Println("listening on :8080")
		http.ListenAndServe(":8080", headers(http.DefaultServeMux))
	*/
}
