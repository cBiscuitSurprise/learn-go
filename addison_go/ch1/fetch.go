// fetch prints the content found at each URL
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, uri := range os.Args[1:] {
		res, err := http.Get(uri)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error during fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading body: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
