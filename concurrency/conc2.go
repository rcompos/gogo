package main

import(
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Hello")
	wg.Add(1)
	go middle()
	fmt.Println("Goodbye")
	wg.Wait()
}

func middle() {
	fmt.Println("¡Woo-woo!")
	wg.Done()
}
