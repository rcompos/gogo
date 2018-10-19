package main

import "fmt"

func main() {
	//tf := 32.0
	//tf := 0.0
	tf := 212.0
	tc := (5.0 / 9.0) * (tf - 32.0)
	fmt.Println("Fahrenheit: ", tf)
	fmt.Println("Celsius: ", tc)
}
