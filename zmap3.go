package main

import(
	"fmt"
)

func main() {

	mrmap := make(map[string]float64)
	mrmap["pi"] = 3.141592653
	mrmap["e"] = 2.718281828
	mrmap["g"] = 1.618033988

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
