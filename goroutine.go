package main

import "fmt"

func main() {
	fmt.Println("Hello!")
	for i := 1; i < 11; i++ {
		fmt.Println("Number ", i, "?")
		go func() {
			fmt.Println("Here we go", i, "!")
		}()
	}
	fmt.Println("Goodbye!")
	fmt.Scanln()
	fmt.Println("Done.")
}
