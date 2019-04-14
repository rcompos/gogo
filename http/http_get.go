package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {

	for _, url := range os.Args[1:] {

		fmt.Println(url)
		resp, err := http.Get(url)

		if err != nil {
			//return nil, err
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			//return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
			fmt.Printf("getting: %s: %s", url, resp.Status)
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			//return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
			fmt.Printf("parsing %s as HTML: %v", url, err)
		}

		//fmt.Println(doc)
		fmt.Printf("%s", doc)

	}
}
