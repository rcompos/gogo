package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! I'm a HTTP server!")
	})

	fs := http.FileServer(http.Dir("src/"))
	http.Handle("/src/", http.StripPrefix("/src/", fs))

	http.ListenAndServe(":8080", nil)
}
