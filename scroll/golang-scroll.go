package main

import(
	"fmt"
	"github.com/fatih/color"
	//"bufio"
	"io/ioutil"
)

func greet() {
	fmt.Println("Hello world!")
	color.Red("Hello world!")
	color.Green("Hello world!")
	color.Blue("Hello world!")
	color.Yellow("Hello world!")
}

func main() {

    greet()

	readmefile := "/Users/rcompos/srv/gopl.io/README.md"

	b, err := ioutil.ReadFile(readmefile)
    if err != nil {
		fmt.Println(err)
	}

    #fmt.Println(b)

	str := string(b)
    fmt.Println(str)

}
