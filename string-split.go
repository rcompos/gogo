package main

import(
	"fmt"
	"strings"
)

func main() {
	addr := "127.0.0.1:8080"
	fmt.Println(addr)
	s := strings.Split(addr, ":")
	fmt.Println(s)
	ip, port := s[0], s[1]
	fmt.Printf("IP: %s\nPort: %v\n", ip, port)
}
