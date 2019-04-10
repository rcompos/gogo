package main

import (
	"fmt"

	//"code.google.com/p/gcfg"
    "gopkg.in/gcfg.v1"
)

func main() {
	config := struct {
		Section struct {
			Enabled bool
			Debug   bool
			Path    string
		}
	}{}

	err := gcfg.ReadFileInto(&config, "conf.ini")
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Debug)
	fmt.Println(config.Section.Path)
}
