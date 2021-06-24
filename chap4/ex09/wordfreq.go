package main

import (
	"bufio"
	"fmt"
	"os"
)

const filename = "test.txt"

//count the number of Unicode
func main() {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Open File: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	countsWord := make(map[string]int)
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		countsWord[word]++
	}

	fmt.Printf("word\tcount\n")
	for k, v := range countsWord {
		fmt.Printf("%s\t%d\n", k, v)
	}

}
