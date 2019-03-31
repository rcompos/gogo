package main

//import "fmt"

func BubbleSort (items []int) {
	L := len(items)
	for i := 0; i < L; i++ {
		for j := 0; j < (L - 1 - i); j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
	return items
}

func main() {
	s := []int{0, 1, 2, 3, 10, 4, 5, 6, 7, 8, 9}
	//fmt.Println(BubbleSort(s))
	BubbleSort(s)
}
