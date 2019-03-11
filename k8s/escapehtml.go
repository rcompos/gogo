package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Search struct {
	//Query string `json:"query"`
	Query string `json:"query"`
}

func main() {
	data := &Search{Query: "http://google.com/?q=stackoverflow&ie=UTF-8"}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(data)
	fmt.Println(string(buf.String()))
}
