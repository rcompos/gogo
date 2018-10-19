package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("When's Saturday?")
	fmt.Println("When's Tuesday?")
	today := time.Now().Weekday()
	fmt.Println("Time: ", time.Thursday)
	fmt.Println("Time: ", time.Friday)
	fmt.Println("Today: ", today)
	fmt.Println("Other: ", today + 1)
	fmt.Println("Other: ", today + 2)
	fmt.Println("Other: ", today + 3)
	fmt.Println("Other: ", today + 4)
	fmt.Println("Other: ", today + 5)
	//switch time.Saturday {
	//switch time.Tuesday {
	//switch time.Thursday {
	switch time.Thursday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("In three days.")
	case today + 4:
		fmt.Println("In four days.")
	case today + 5:
		fmt.Println("In five days.")
	default:
		fmt.Println("Too far away.")
	}
}
