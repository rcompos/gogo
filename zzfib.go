package main

import "fmt"
import "os"

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
	x, y, n := 0, 1, 50
	for i := 0; i < n; i++ {
		x, y = y, y + x
		fmt.Println(x)
	}
	fmt.Println("2**32: ", 1<<32)
	fmt.Println("2**48: ", 1<<48)
	fmt.Println("2**63: ", 1<<63 - 1)
}
