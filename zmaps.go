package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex {
		40.68433, -74.39967,
	}
	m["TacoBell"] = Vertex {
		55.5555, -55.55555,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["TacoBell"])
}
