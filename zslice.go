package main

import "fmt"

func main() {
	alpha := [5]string{"alpha", "beta", "gamma", "delta", "epsilon"}
    slice := alpha[3:]
	fmt.Println("slice: ", slice)

	a := [9]int{11, 12, 13, 14, 15, 16, 17, 18, 19}
	fmt.Println("a: ", a)
	s := a[3:6]
	s2 := a[5:9]
	fmt.Println("s: ", s)
	fmt.Println("s2: ", s2)
	s[0] = 99
	s[1] = 0
	s[2] = 2018
	fmt.Println("s: ", s)
	fmt.Println("s2: ", s2)
	fmt.Println("a: ", a)
}
