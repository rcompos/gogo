package main

import(
	"fmt"
	"bufio"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
	. "github.com/logrusorgru/aurora"
	"gopkg.in/kyokomi/emoji.v1"
)

const sleepytime time.Duration = 1

func greet(greeting string) {
    //greeting := "Hello!"
	//fmt.Println(Magenta(greeting))
	emoji.Println(strings.Repeat(greeting,20))
	time.Sleep(sleepytime * time.Second)
}

func main() {

	//fmt.Printf("Time: %v\n", time.Now())
	fmt.Println(Magenta(time.Now()))
    greet(":hamburger:")

    const directory string = "/Users/rcompos/srv/gopl.io"
	const sleepytime time.Duration = 2
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
    //fmt.Println(str)
	scanner := bufio.NewScanner(strings.NewReader(string(str)))
	for scanner.Scan() {
		thisline := scanner.Text()
		fmt.Println(thisline)
		time.Sleep(sleepytime * time.Second)
	}

    d, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range d {
		//if f.Name() == ".git" { continue }
		//if f.Name() == "README.md" { continue }
		if f.Name() == ".git" || f.Name() == "README.md" { continue }
		//fmt.Println(f.Name())
		//fmt.Println(strings.Repeat(f.Name() + " ", 20))
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
				currentdir2 := currentdir + "/" + f2.Name()
				//fmt.Println(currentdir2) 
				filego, err := ioutil.ReadDir(currentdir2)
				if err != nil {
					fmt.Println(err)
				}
				for _, f3 := range filego {
					//fmt.Println(f3.Name())
					//greet(strings.Repeat(f.Name() + " ", 20))
					greet(":hamburger:")
					//fmt.Println(Cyan(strings.Repeat(f.Name() + " ", 20)))
					fmt.Println(currentdir2 + "/" + f3.Name())
                    //fmt.Println(Green(strings.Repeat(f3.Name() + " ", 10)))
					greet(":beer:")
					cf, err := ioutil.ReadFile(currentdir2 + "/" + f3.Name())
					if err != nil {
						fmt.Println(err)
					}
					//fmt.Println(string(cf))
					scanner := bufio.NewScanner(strings.NewReader(string(cf)))
					for scanner.Scan() {
						thisline := scanner.Text()
						fmt.Println(thisline)
						time.Sleep(sleepytime * time.Second)
					}
				}
			}
		}
	}
	for _, f := range d {
		if f.Name() == ".git" || f.Name() == "README.md" { continue }
		if rdouble.MatchString(f.Name()) {
			//fmt.Println("Double ", f.Name())
			//fmt.Println(strings.Repeat(f.Name() + " ", 20))
            currentdir := directory + "/" + f.Name()
            d2, err := ioutil.ReadDir(currentdir)
            if err != nil {
                fmt.Println(err)
            }
            for _, f2 := range d2 {
                //fmt.Println(" ", f2.Name())
                currentdir2 := currentdir + "/" + f2.Name()
                //fmt.Println(currentdir2)
                filego, err := ioutil.ReadDir(currentdir2)
                if err != nil {
                    fmt.Println(err)
                }
                for _, f3 := range filego {
                    //fmt.Println(f3.Name())
					//greet()
					//greet(strings.Repeat(f.Name() + " ", 20))
					greet(":hamburger:")
					fmt.Println(Cyan(strings.Repeat(f.Name() + " ", 20)))
					fmt.Println(currentdir2 + "/" + f3.Name())
                    fmt.Println(Green(strings.Repeat(f3.Name() + " ", 10)))
					greet(":beer:")

                    cf, err := ioutil.ReadFile(currentdir2 + "/" + f3.Name())
                    if err != nil {
                        fmt.Println(err)
                    }
                    //fmt.Println(string(cf))
                    scanner := bufio.NewScanner(strings.NewReader(string(cf)))
                    for scanner.Scan() {
                        thisline := scanner.Text()
                        fmt.Println(thisline)
                        time.Sleep(sleepytime * time.Second)
					}
                }
            }
        }
	}

	fmt.Println(Magenta(time.Now()))

}
