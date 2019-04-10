package main

import(
	"fmt"
	"time"
)

func main() {
	fmt.Println("time.Sleep")
	time.Sleep(2 * time.Second)
	fmt.Println("time.After")
	sleep := time.After(2 * time.Second)
	<-sleep
}
