package main

import(
	"runtime/debug"
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
	debug.PrintStack()
}
