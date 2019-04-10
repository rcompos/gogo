package main

import(
	"fmt"
	"runtime"
)

func main() {
	alpha()
}

func alpha() {
	beta()
}

func beta() {
	gamma()
}

func gamma() {
	buf := make([]byte, 1024)
	//runtime.Stack(buf, false)
	runtime.Stack(buf, true)
	fmt.Printf("Trace: \n%s\n", buf)
}
