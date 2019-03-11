package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Example struct {
	Int    int       `json:"int,string"`
	Date   time.Time `json:"date"`
	String string    `json:"string"`
}

func main() {
	q, _ := url.ParseQuery("int=12&date=2018-04-25T22:28:56.110Z&string=it+works")
	// q is map[string][]string, but we probably want a map[string]string, so...
	values := make(map[string]string)
	for key, _ := range q {
		values[key] = q.Get(key)
	}
	// marshal it
	js, err := json.Marshal(values)
	if err != nil {
		panic(err)
	}
	// unmarshal it
	e := &Example{}
	json.Unmarshal(js, e)
	fmt.Println(e)

	var qqq string
	qqq = "int=12&date=2018-04-25T22:28:56.110Z&string=it+works"

	fmt.Println("qqq: ", qqq)
}
