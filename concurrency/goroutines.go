package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		r := rand.Intn(1000)
		//time.Sleep(randy * time.Second)
		time.Sleep(time.Duration(r) * time.Millisecond)
		fmt.Println(n, ":", i)
	}
}

func main() {
	n := 10

	//go f(10)
	for i := 0; i <= n; i++ {
		go f(i)
	}

	var input string
	fmt.Scanln(&input)
}
