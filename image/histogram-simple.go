package main

import (
	"fmt"
)

const N = 10

var n = [N]int{20, 7, 11, 13, 1, 3, 30, 4, 9, 2}

func main() {
	fmt.Printf("%s %s %s\n", "Element", "Value", "Histogram")

	for i := 0; i <= N-1; i++ {
		fmt.Printf("%7d %d        ", i, n[i])

		for j := 1; j <= n[i]; j++ {
			//fmt.Printf("%c", '*')
			fmt.Printf("%c", '∎')
		}
		fmt.Println()
	}
}
