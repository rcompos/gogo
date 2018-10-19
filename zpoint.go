package main

import "fmt"

func main() {
	a, b := 1, 77
	p := &a
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("*p: ", *p)
	*p = 33
	fmt.Println("a: ", a)
	p = &b
	*p = *p / 11
	fmt.Println("b: ", b)
}
