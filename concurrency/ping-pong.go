package main

import(
	"fmt"
	"time"
)

//func ping(c chan string) {
func ping(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

//func printly(c chan string) {
func printly(c <-chan string) {
	for{
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	c := make(chan string)
	go ping(c)
	go pong(c)
	go printly(c)

	var input string
	fmt.Scanln(&input)
}
