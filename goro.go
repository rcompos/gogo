package main

import "fmt"

func main() {

	fmt.Println("Hello")
	go func() {
		fmt.Println("Hmmm")
	}()
	fmt.Println("Goodbye")
	fmt.Scanln()

}
