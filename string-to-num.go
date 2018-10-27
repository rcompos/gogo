package main

import(
	"fmt"
	"strconv"
	"time"
)

func main() {

	fmt.Println("Hello wurld!")
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	fmt.Println(strconv.FormatInt(int64(x), 2))

	s := fmt.Sprintf("x=%b", x)
	fmt.Println(s)

	//xx, errXx := strconv.Atoi("123")
	//fmt.Println(xx,errXx)
	xx, _ := strconv.Atoi("123")
	fmt.Println(xx)
	yy, errYy := strconv.ParseInt("123", 10, 64)
	fmt.Println(yy,errYy)

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

}
