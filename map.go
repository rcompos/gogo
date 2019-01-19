package main

import(
	"fmt"
	"math"
)

func main() {
	x := 0.0
	y := 9.0
	x = math.Sqrt(y)
	fmt.Println(x)

	a := []string{"hello", "goodbye", "salutations", "greetings", "howdy"}
    for _, val := range a {
		fmt.Println(val)
	}

	m := make(map[string]string)
	m["co"] = "denver"
	m["ca"] = "sacramento"
	m["ut"] = "saltlakecity"
	m["nv"] = "carsoncity"
	m["or"] = "salem"
	m["wa"] = "olympia"
	fmt.Println("mmap: ", m)
	fmt.Println("mmap: ", m["nv"])

	n := make(map[string]bool)
	n["co"] = true
	n["ca"] = true
	n["ut"] = false
	n["nv"] = false
	n["or"] = true
	n["wa"] = true
	fmt.Println("nmap: ", n)
	fmt.Println("nmap: ", n["nv"])

}
