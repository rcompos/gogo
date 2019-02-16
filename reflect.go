package main

import(
	"fmt"
	"reflect"
)

func main() {
	var c int32
	fmt.Println("Value of c: ", reflect.ValueOf(c))
	fmt.Println("Type of c: ", reflect.TypeOf(c))
}
