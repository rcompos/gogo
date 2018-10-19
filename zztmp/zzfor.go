package main

import( 
	"fmt"
)

func main() {

	tots := 0
	sum := 0
	for i := 0; i < 10; i++ {
		tots += 10
		sum += i
		fmt.Println(i, tots, sum)
	}

}
