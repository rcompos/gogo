package main

import(
	"fmt"
	//"math"
)

func main() {
	n := 1000
	//var s []int
	s := []int{}
	printSlice(s)
	//s = append(s, 1999)
	//printSlice(s)

	for i := 0; i < n; i++ {
		count := 0
		for j := i - 1; j > 0; j-- {
			//fmt.Println(i, j)
			if i%j == 0 {
				if j == 1 { continue }
				//fmt.Println("nope")
				count++
			}
		}
		if count == 0 {
			//fmt.Println("Prime: ", i)
			s = append(s, i)
		}
	}
	//printSlice(s)
	fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
