// dup1 finds the duplicate lines from stdin
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		handleFromStdin(counts)
	} else {
		handleFromFiles(files, counts)
	}

	printCounts(counts)
}

func handleFromStdin(counts map[string]int) {
	countLines(os.Stdin, counts)
}

func handleFromFiles(files []string, counts map[string]int) {
	for _, file := range files {
		fid, err := os.Open(file)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		}

		countLines(fid, counts)
		fid.Close()
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := f.Name() + "\t" + input.Text()
		counts[key]++
	}
}

func printCounts(counts map[string]int) {
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
