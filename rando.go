package main

import (
    "fmt"
    "math"
    "math/rand"
)

var c, python, java bool

//func add(x int, y int) int {
func add(x, y int) int {
    return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {

    fmt.Println("My favorite number from 0 to 100 is ", rand.Intn(100))
    fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
    fmt.Println(math.Pi)
    fmt.Println(add(91, 131))
    fmt.Println(split(170))
    
    var i, j int = 7, 77
    fmt.Println(i, j, c, python, java)
    //i := 7
    //fmt.Println(i, c, python, java)

}
