package main

import "fmt"

func main() {
    a := make([]int, 3, 13)
    b := [10]int{}
    fmt.Println("a: ", a)
    fmt.Println("b: ", b)

    for i, j := range b {
        fmt.Println(i,j)
    }
    for i, _ := range b {
        fmt.Println(i)
    }
    for _, j := range b {
        fmt.Println(j)
    }
}
