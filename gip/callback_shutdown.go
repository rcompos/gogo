package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func shutdown(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Shutting down immediately.\n")
	os.Exit(0)
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello!\n")
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "Welcome to callback shutdowner.\n")
}
