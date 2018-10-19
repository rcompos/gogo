package main

import "fmt"

func main() {
	//n := 10
	//x, y := 0, 1
	x, y, n := 0, 1, 10
	for i := 0; i < n; i++ {
		x, y = y, y + x
		fmt.Println("x: ", x, "   y: ", y)
	}
    fmt.Println(x)
}
