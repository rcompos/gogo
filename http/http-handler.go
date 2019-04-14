package main

import (
	"fmt"
  "log"
  "net/http"
)

func boozer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there!")
}

func hoozer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")
}

func main() {
	mux := http.NewServeMux()
	rh := http.RedirectHandler("http://example.org", 307)
	mux.Handle("/foo", rh)
	mux.HandleFunc("/boo", boozer)
	mux.HandleFunc("/hoo", hoozer)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
