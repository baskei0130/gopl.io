package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

//count the number of Unicode
func main() {
	counts := make(map[rune]int)
	var countsLetter, countsNum, countsOther int
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // return rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		if unicode.IsLetter(r) {
			countsLetter++
		} else if unicode.IsNumber(r) {
			countsNum++
		} else {
			countsOther++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	fmt.Print("Letter\tNumber\tOther\n")
	fmt.Print("\nLetter\tNumber\tOther\n")
	for i, n := range utflen {
		fmt.Printf("%d\t%d\n", i, n)
	}
}
