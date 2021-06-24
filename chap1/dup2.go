package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames, "input")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %vn", err)
			}
			countLines(f, counts, filenames, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s:%d\t%s\n", filenames[line], n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filenames map[string]string, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		filenames[input.Text()] = filename
	}
}
