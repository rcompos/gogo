package main

import(
	"fmt"
	"strings"
	"io/ioutil"
	"gopkg.in/kyokomi/emoji.v1"
)

const emojiFile = "emoji.txt"

func greet(greeting string) {
    //greeting := "Hello!"
	//fmt.Println(Magenta(greeting))
	emoji.Println(strings.Repeat(greeting,20))
}

func main() {

    greet(":hamburger:")

	f, err := ioutil.ReadFile(emojiFile)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(string(f))
	//emoji.Println(string(f))

	for _, val := range strings.Split(string(f), "\n") {
		//fmt.Println(string(val))
		/*
		fmt.Printf("%s ", string(val))
		emoji.Println(string(val))
		*/
		emoji.Printf("%s ", string(val))
	}

	fmt.Println()
}
