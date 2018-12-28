package main

import(
	"fmt"
	"strings"
	"os"
)

func main() {
	//unsorted := "insertionsort"
	//sorted    := insertion_sort(strings.ToUpper(unsorted))
	sorted    := insertion_sort([]byte(strings.ToUpper(os.Args[1])))
	fmt.Println("Sorted: ", sorted)
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
