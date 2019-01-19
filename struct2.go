package main

import(
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

type Account struct {
	user string
	home string
}


func main() {
	fmt.Println(Vertex{1,2})
	fmt.Println(Account{"rcompos","/home/rcompos"})
	v := Account{"sjobs","/home/sjobs"}
	fmt.Println(v.user)
	fmt.Println(v.home)
}
