package main

import(
	"fmt"
)

func main() {

	mrmap := map[string]float64{
		"pi": 3.141592653,
		"e":  2.718281828,
		"g":  1.618033988,
	}

	for _, val := range mrmap {
		fmt.Println(val)
	}

	i, j := mrmap["pi"]
	fmt.Println(i,j)

	k, l := mrmap["pie"]
	fmt.Println(k,l)

	fmt.Println(mrmap["e"])
	if m, n := mrmap["e"]; n {
		fmt.Println(m)
	}

	delete(mrmap, "e")
	fmt.Println(mrmap["e"])

}
