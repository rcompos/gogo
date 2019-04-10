package main

import (
	"fmt"
)

func aloha() string {
	//fmt.Println("¡Aloha World!")
	return "World"
}

func main() {
	whom := aloha()
	//fmt.Println("¡Aloha ",whom,"!")
	fmt.Printf("¡Aloha %s!\n", whom)
}
