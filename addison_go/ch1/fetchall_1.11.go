// fetchall fetches mulitple sites concurrently and reports the length of time
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	fid, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v", err)
		os.Exit(1)
	}

	r := csv.NewReader(fid)
	var urls []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open file: %v", err)
		}

		newUrl := record[1]
		if !strings.HasPrefix(newUrl, "http") {
			newUrl = "https://" + newUrl
		}
		urls = append(urls, newUrl)
	}
	fmt.Fprintf(os.Stdout, "%d\n", len(urls))

	for _, url := range urls {
		// fmt.Fprintf(os.Stdout, "fetching url, '%s'\n", url)
		go fetch(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	elapsed := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", elapsed, nbytes, url)
}
