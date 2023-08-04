// fetchall fetches mulitple sites concurrently and reports the length of time
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	fid, err := os.Create(strings.ReplaceAll(url, "https://", "demo."))
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(fid, resp.Body)
	fid.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	elapsed := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", elapsed, nbytes, url)
}
