package main

import(
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello wurld!")
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
}
