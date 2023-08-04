// fetch prints the content found at each URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, uri := range os.Args[1:] {
		if !strings.HasPrefix(uri, "http") {
			uri = "https://" + uri
		}
		res, err := http.Get(uri)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error during fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Status: %d\n\n", res.StatusCode)
		_, err = io.Copy(os.Stdout, res.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading body: %v\n", err)
		}
	}
}
