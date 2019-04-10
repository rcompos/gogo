package main

import(
	"fmt"
	"log"
	"os/exec"
)

func main() {
	//err := checkDep("fortune")
	myDep := "go"
	err := checkDep(myDep)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println("Time to get your fortunte")
	fmt.Println("Dependency found: ", myDep)
}

func checkDep(name string) error {
	if _, err := exec.LookPath(name); err != nil {
		// Returns an error when the dependency isn’t found
		es := "Could not find '%s' in PATH: %s" 
		return fmt.Errorf(es, name, err)
	}
	// Returning nil if there was no error
	return nil
}
