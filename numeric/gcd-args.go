package main

import(
	"fmt"
	"flag"
)

func main() {
	var x, y int
	flag.IntVar(&x, "x", 1, "X value")
	flag.IntVar(&y, "y", 1, "Y value")
	flag.Parse()
	//x := 133
	//y := 25
	fmt.Println("x0: ", x)
	fmt.Println("y0: ", y)
	for y != 0 {
		x, y = y, x%y
		fmt.Println("x ", x)
		fmt.Println("y ", y)
	}
	fmt.Println(x)
}
