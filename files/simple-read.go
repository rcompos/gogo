package main

import(
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Supply file as an argument.")
		os.Exit(1)
	}
	file := os.Args[1]
	//file := "/tmp/dat"
	dat, err := ioutil.ReadFile(file)
	check(err)
	fmt.Print(string(dat))
}

