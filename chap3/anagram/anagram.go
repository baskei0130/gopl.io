package main

import (
	"fmt"
	"reflect"
)

func anagram(s1, s2 string) bool {
	map1 := make(map[string]int, 100)
	map2 := make(map[string]int, 100)
	for _, v := range s1 {
		map1[string(v)]++
	}
	for _, v := range s2 {
		map2[string(v)]++
	}
	return reflect.DeepEqual(map1, map2)
}

func main() {
	fmt.Println(anagram("acafaf", "fafaac"))
	fmt.Println(anagram("acafaf", "fafaab"))
}
