package main

import(
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello")
	go middle()
	fmt.Println("Goodbye")
	time.Sleep(3 * time.Second)
}

func middle() {
	fmt.Println("¡Woo-woo!")
}
