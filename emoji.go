package main

import (
	"fmt"
	"strings"
	"gopkg.in/kyokomi/emoji.v1"
)

func adder(a int, b int) int {
	return a + b
}

func main() {

	//emoji.Println(":beer: Beer!!!")
	emoji.Println(strings.Repeat(":beer:", 10))

	pizzaMessage := emoji.Sprint("I like a :pizza: and :sushi:!!")
	fmt.Println(pizzaMessage)

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

	emoji.Println(strings.Repeat(":coffee:", 10))

	for k := 0; k < 11; k++ {
		fmt.Println("Hello ", k)
	}

	emoji.Println(strings.Repeat(":beer:", 10))

}
