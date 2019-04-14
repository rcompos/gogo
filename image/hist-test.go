package main

import (
	"fmt"
)

func main() {
	vals := make(map[string]int)
	vals["red"] = 5
	vals["blue"] = 8
	vals["orange"] = 3
	vals["yellow"] = 12
	vals["purple"] = 34
	vals["orange"] = 22
	vals["magenta"] = 45
	vals["cyan"] = 9
	for i, v := range vals {
		//fmt.Printf("i: %v v: %v\n", i, v)
		//fmt.Printf("%v %v\n", i, v)
		fmt.Printf("%10v %v\n", i, v)
	}
}
