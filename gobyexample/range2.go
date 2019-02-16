package main

import (
	"fmt"
)

var pow = []int{1,2,4,8,16,32,64,128}

func main() {

	s := []int{0,1,2,3,4,5,6,7,8,9}
	for _, v := range s {
		fmt.Println("v: ", v)
	}
	subset := s[2:6]
	fmt.Println(subset)

	for i := 10; i < 20; i++ {
		s = append(s, i)
	}
	fmt.Println(s)

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

}
