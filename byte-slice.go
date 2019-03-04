package main

import "fmt"

func main() {
	//fmt.Println([]byte("Hello world!"))
	bs := []byte("♔  Hello world! Ó´¬¬ø ∑ø®¬∂¡ 睧睷 ♔")
	fmt.Println("Byteslice: ", bs)
	fmt.Println("String: ", string(bs))
}
