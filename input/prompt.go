package main

import(
	"fmt"
	"bufio"
	"strings"
	"os"
	"log"
)

func main() {
	
	PromptReturn()
    fmt.Printf(">\n")

	PromptRead()
    fmt.Printf(">\n")

}

func PromptReturn() {
    fmt.Printf("-> Press Return key to delete pods.")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        break
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }
    fmt.Println()
}

func PromptRead() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("Type y to continue: ")
    text, _ := reader.ReadString('\n')
    answer := strings.TrimRight(text, "\n")
    //fmt.Printf("answer: %s \n", answer)
    if answer == "y" || answer == "Y" {
        return
    } else {
        //prompt2() //For recursive prompting
        log.Fatal("Exiting without action.")
    }
}
