package main

import(
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	for _, filename := range os.Args[1:] {
		fmt.Println(filename)
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(file)
		//fmt.Println(string(file))
		for _, line := range strings.Split(string(file), "\n") {
			fmt.Println("> " + line)
		}
	}
}
