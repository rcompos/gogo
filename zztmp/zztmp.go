package main

import (
	"fmt"
)

func adder(a int, b int) int {
	return a + b
}

func main() {

	fmt.Println("Total: ", adder(6, 9))

	fmt.Println("Fibonacci!")
	j := 0
	prev := 0
	for i := 0; i <= 10; i++ {
		//fmt.Println(i)
		if i == 1 {
			j = 1
		}
		jold := j
		j = j + prev
		prev = jold
		fmt.Println(i, j)
	}

	for k := 0; k < 11; k++ {
		fmt.Println("Hello ", k)
	}

}
