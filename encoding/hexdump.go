package main

import(
	"fmt"
	"encoding/hex"
)

func main() {
    //fmt.Printf("Hello, world; or ÎšÎ±Î»Î·Î¼Î­ÏÎ± ÎºÏŒÏƒÎ¼Îµ; or ã“ã‚“ã«ã¡ã¯ ä¸–ç•Œ\n")
    a := "Hello, world; or ÎšÎ±Î»Î·Î¼Î­ÏÎ± ÎºÏŒÏƒÎ¼Îµ; or ã“ã‚“ã«ã¡ã¯ ä¸–ç•Œ\n"
    fmt.Printf("HW: %s\n", a)
	b := []byte("Hello, world; or ÎšÎ±Î»Î·Î¼Î­ÏÎ± ÎºÏŒÏƒÎ¼Îµ; or ã“ã‚“ã«ã¡ã¯ ä¸–ç•Œ\n")
	fmt.Println(hex.Dump(b))
}
