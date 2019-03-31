// https://tour.golang.org/methods/10
//
package main

import "fmt"

type T struct {
	S string
}

type I interface {
	M()
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello world"}
	i.M()
}
