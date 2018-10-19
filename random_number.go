package main

import (
    "fmt"
    "math/rand"
	"time"
)


func main() {

	rand.Seed(time.Now().UnixNano())
    fmt.Println("My favorite number from 0 to 100 is ", rand.Intn(100000000000))

}
