package main

import(
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""
    //for _, arg := range os.Args[1:] {
    for ind, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
        fmt.Println(ind, arg)
    }
    fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:3], " "))

    fmt.Println(os.Args[1:])
}
