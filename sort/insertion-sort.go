package main

import(
	"fmt"
	"strings"
	//"os"
	"flag"
)

func main() {
	var u string
	flag.StringVar(&u, "u", "insertionsort", "String to sort")
	flag.Parse()

	//unsorted := "insertionsort"
	//unsorted = strings.ToUpper(os.Args[1])
	unsorted := strings.ToUpper(u)
	fmt.Println(unsorted)

	//sorted    := insertion_sort(strings.ToUpper(unsorted))
	//sorted    := insertion_sort([]byte(strings.ToUpper(os.Args[1])))
	//fmt.Println("Sorted: ", sorted)

	insertion_sort([]byte(unsorted))

}

func insertion_sort(s []byte) string {
	var j int
	for i := 1; i < len(s); i++ {
		j = i
		for j > 0 {
			if string(s[j]) < string(s[j-1]) {
				s[j], s[j-1] = s[j-1], s[j]
			}
			j = j - 1
		}
		fmt.Println(string(s))
	}
	return(string(s))
}
