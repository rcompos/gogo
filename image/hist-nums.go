package main

import (
	"fmt"
	"sort"
)

func main() {
	vals := make(map[int]int)
	vals[10] = 5
	vals[11] = 8
	vals[13] = 12
	vals[14] = 34
	vals[15] = 22
	vals[16] = 45
	vals[12] = 3
	vals[17] = 9

	fmt.Println("Raw")
	for i, v := range vals {
		//fmt.Printf("i: %v v: %v\n", i, v)
		//fmt.Printf("%v %v\n", i, v)
		fmt.Printf("%10v %v\n", i, v)
	}

	fmt.Println("Sorted")
	keys := make([]int, 0, len(vals))
	for k := range vals {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println(k, vals[k])
	}

}
