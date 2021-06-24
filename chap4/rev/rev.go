package main

import (
	"fmt"
)

// reverce slice of int
func reverce(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverce2(s *[8]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, n int) {
	if n < 0 || len(s) == 0 {
		return
	}
	r := len(s) - n%len(s)
	s = append(s[r:], s[:r]...)
	fmt.Println(s)
}

func removeAdj(strings []string) []string {
	if len(strings) == 0 {
		return strings
	}
	prev := strings[0]
	ptr := 1
	for _, s := range strings {
		if s != prev {
			strings[ptr] = s
			ptr++
			prev = s
		}
	}
	return strings[:ptr]
}

func main() {
	s := [8]int{3, 4, 2, 1, 2, 4, 5, 0}
	//reverce(s[:4])
	reverce2(&s)
	fmt.Println(s)
	t := []int{3, 4, 2, 1, 2, 4, 5}
	rotate(t, 4)
	fmt.Println(t)

	strings := []string{"aaa", "jfaj", "jll", "jll", "fjaoadlsj", "aa", "aa"}
	fmt.Println(strings)
	strings = removeAdj(strings)
	fmt.Println(strings)
}
