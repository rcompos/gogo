package main

import(
	"fmt"
	"github.com/fatih/color"
	//"bufio"
	"io/ioutil"
	"regexp"
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

    const directory string = "/Users/rcompos/srv/gopl.io"
	//readmefile := "/Users/rcompos/srv/gopl.io/README.md"
	readmefile := directory + "/README.md"

    rsingle, _ := regexp.Compile("ch[0-9]$")
    rdouble, _ := regexp.Compile("ch[0-9][0-9]$")

	b, err := ioutil.ReadFile(readmefile)
    if err != nil {
		fmt.Println(err)
	}

    //fmt.Println(b)

	str := string(b)
    fmt.Println(str)

    d, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range d {
		//if f.Name() == ".git" { continue }
		//if f.Name() == "README.md" { continue }
		if f.Name() == ".git" || f.Name() == "README.md" { continue }
		//fmt.Println(f.Name())
		//r, _ := regexp.MatchString("ch[0-9]+", f.Name())
		//fmt.Println(r)
		if rsingle.MatchString(f.Name()) {
			//fmt.Println("Single ", f.Name())
			//fmt.Println(f.Name())
			currentdir := directory + "/" + f.Name()
		    d2, err := ioutil.ReadDir(currentdir)
			if err != nil {
				fmt.Println(err)
			}
			for _, f2 := range d2 {
				//fmt.Println("	", f2.Name())
				fmt.Println(currentdir + "/" + f2.Name())
			}
		}

	}
	for _, f := range d {
		if f.Name() == ".git" || f.Name() == "README.md" { continue }
		if rdouble.MatchString(f.Name()) {
			//fmt.Println("Double ", f.Name())
			//fmt.Println(f.Name())
            currentdir := directory + "/" + f.Name()
            d2, err := ioutil.ReadDir(currentdir)
            if err != nil {
                fmt.Println(err)
            }
            for _, f2 := range d2 {
                //fmt.Println("   ", f2.Name())
				fmt.Println(currentdir + "/" + f2.Name())
            }
        }
	}


}
