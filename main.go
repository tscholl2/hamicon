package main

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/pkg/errors"
)

func blank(w http.ResponseWriter, r *http.Request) {
	icon := newIcon(options{scale: 4, blank: true})
	if err := writeXML(w, icon); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	icon := newIcon(options{seed: rand.Int63(), scale: 4})
	if err := writeXML(w, icon); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func writeXML(w http.ResponseWriter, v interface{}) error {
	w.Header().Set("Content-Type", "image/svg+xml;charset=utf-8")
	w.Header().Set("content-encoding", "gzip")
	W := gzip.NewWriter(w)
	defer W.Close()
	buf := &bytes.Buffer{}
	encoder := xml.NewEncoder(buf)
	encoder.Indent("", "  ")
	if err := encoder.Encode(v); err != nil {
		return errors.Wrapf(err, "unable to encode")
	}
	_, err := fmt.Fprint(W, minimize(string(buf.Bytes())))
	return errors.Wrapf(err, "unable to write")
}

func main() {
	http.HandleFunc("/", get)
	http.HandleFunc("/blank", blank)
	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
