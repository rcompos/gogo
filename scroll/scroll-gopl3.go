package main

import(
	"fmt"
	"github.com/fatih/color"
	"bufio"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
)

func greet(greeting string) {
    //greeting := "zzscroll!"
	const sleepytime time.Duration = 2
	color.Cyan(greeting)
	time.Sleep(sleepytime * time.Second)
	color.Red(greeting)
	time.Sleep(sleepytime * time.Second)
	color.Green(greeting)
	time.Sleep(sleepytime * time.Second)
	color.Blue(greeting)
	time.Sleep(sleepytime * time.Second)
	color.Yellow(greeting)
	time.Sleep(sleepytime * time.Second)
	color.Magenta(greeting)
	time.Sleep(sleepytime * time.Second)
}

func main() {

	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
    greet("Hello!!")

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
		//color.Red(strings.Repeat(f.Name() + " ", 20))
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
				//color.Green(currentdir2) 
				filego, err := ioutil.ReadDir(currentdir2)
				if err != nil {
					fmt.Println(err)
				}
				for _, f3 := range filego {
					//fmt.Println(f3.Name())
					greet(strings.Repeat(f.Name() + " ", 20))
					color.Red(strings.Repeat(f.Name() + " ", 20))
					color.Green(currentdir2)
                    color.Yellow(strings.Repeat(f3.Name() + " ", 10))
					color.Cyan("Current Unix Time: %v\n", time.Now().Unix())
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
			//color.Red(strings.Repeat(f.Name() + " ", 20))
            currentdir := directory + "/" + f.Name()
            d2, err := ioutil.ReadDir(currentdir)
            if err != nil {
                fmt.Println(err)
            }
            for _, f2 := range d2 {
                //fmt.Println(" ", f2.Name())
                currentdir2 := currentdir + "/" + f2.Name()
                //color.Green(currentdir2)
                filego, err := ioutil.ReadDir(currentdir2)
                if err != nil {
                    fmt.Println(err)
                }
                for _, f3 := range filego {
                    //fmt.Println(f3.Name())
					//greet()
					greet(strings.Repeat(f.Name() + " ", 20))
					color.Red(strings.Repeat(f.Name() + " ", 20))
					color.Green(currentdir2)
                    color.Yellow(strings.Repeat(f3.Name() + " ", 10))
					color.Cyan("Current Unix Time: %v\n", time.Now().Unix())
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

	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
}
