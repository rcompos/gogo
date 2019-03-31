package main

import "fmt"

func main() {
	x := 133
	y := 25
	for y != 0 {
		x, y = y, x%y
		fmt.Println("x ", x)
		fmt.Println("y ", y)
	}
	fmt.Println(x)
}
