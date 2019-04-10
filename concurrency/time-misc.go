package main

import(
	 "fmt"
	 "time"
	 "errors"
)

func main() {

	fmt.Printf("Current time: %v\n", time.Now())

	//panic(nil)
	panic(errors.New("Something bad happened."))

	fmt.Printf("Unix time: %v\n", time.Now().Unix())
	go func() {
		fmt.Println("Hmmm")
	}()
	fmt.Println("Goodbye")
	fmt.Scanln()

	fmt.Println("Divide by zero.")
	divider(1, 0)

}

func divider(a, b int) int {
	return a / b
}
