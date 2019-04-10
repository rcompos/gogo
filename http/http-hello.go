package main

import(
	"fmt"
	"net/http"
)

func howdy(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, "Howdy, my name is net/http!\n")
}

func main() {
	http.HandleFunc("/", howdy)
	http.ListenAndServe("localhost:4000", nil)
}
