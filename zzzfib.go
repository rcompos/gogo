package main

import "fmt"

func main() {
	x, y, n := 0, 1, 100
	for i:= 0; i < n; i++ {
		x, y = y, y + x
		fmt.Println(x)
	}
}
