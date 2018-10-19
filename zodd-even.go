package main

import "fmt"
import "os"
import "strconv"

func isEven(n int) bool {
	if n%2 == 0 {
		//fmt.Println("Even: ", n)
		return true
	} else {
		//fmt.Println("Odd: ", n)
		return false
	}
}

func main() {

	for _, arg := range os.Args[1:] {
		i, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("Is even? ", i, isEven(i))
	}
	
}

