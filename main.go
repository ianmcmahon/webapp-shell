package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

var templateMap map[string][]string = map[string][]string{
	"sample_form.html": []string{"sample_form.tmpl"},
	"blank.html":       []string{"blank.tmpl"},
}

func main() {
	http.HandleFunc("/sample/form.html", sampleForm)
	http.HandleFunc("/blank.html", blankHandler)

	http.Handle("/", http.FileServer(http.Dir("static")))

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/// handlers

func sampleForm(rw http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{
		"foo": "bar",
	}

	b := new(bytes.Buffer)
	if err := renderTemplate(b, "aircraft_form.html", data); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Header().Set("Content-Type", "text/plain")
		rw.Write([]byte(fmt.Sprintf("Something went wrong: %v", err)))
	}

	rw.Header().Set("Content-Type", "text/html")
	io.Copy(rw, b)
}

func blankHandler(rw http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{}

	b := new(bytes.Buffer)
	if err := renderTemplate(b, "blank.html", data); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Header().Set("Content-Type", "text/plain")
		rw.Write([]byte(fmt.Sprintf("Something went wrong: %v", err)))
	}

	rw.Header().Set("Content-Type", "text/html")
	io.Copy(rw, b)
}
