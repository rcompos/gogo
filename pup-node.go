package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

func main() {

	var infile = "ai-pup-nodes.txt"
	var show = "true"

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	var nodepath = user.HomeDir + "/srv/control-repo/hieradata/nodes/"

	/*
	   fmt.Println("Hello " + user.Name)
	   fmt.Println("====")
	   fmt.Println("Id: " + user.Uid)
	   fmt.Println("Username: " + user.Username)
	   fmt.Println("Home Dir: " + user.HomeDir)
	*/

	b, err := ioutil.ReadFile(infile) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	//fmt.Println(b) // print the content as 'bytes'

	str := string(b) // convert content to a 'string'

	//fmt.Println(str) // print the content as a 'string'

	scanner := bufio.NewScanner(strings.NewReader(str))
	for scanner.Scan() {
		//fmt.Println(">> ", scanner.Text())
		thisline := scanner.Text()
		//fmt.Println("## ", thisline )
		//nodefile := "./srv/control-repo/hieradata/nodes/" + thisline + ".yaml"
		nodefile := nodepath + thisline + ".yaml"
		//fmt.Println("> ", nodefile)

		if _, err := os.Stat(nodefile); err == nil {
			// path/to/whatever exists
			stringlength := 50
			stringline := strings.Repeat("#", stringlength)
			fmt.Println(stringline + "\n")
			fmt.Println(nodefile + "\n")
			// Read file to byte slice
			data, err := ioutil.ReadFile(nodefile)
			if err != nil {
				log.Fatal(err)
			}
			if show == "true" {
				fmt.Printf("Data read: %s\n", data)
			}
		}

	}

}
