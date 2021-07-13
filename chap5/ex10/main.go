package main

import (
	"fmt"
	"sort"
)

// 情報科学の各講座をそれぞれの事前条件となる講座と対応づける
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal langeages",
		"computer organiization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"dicrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) map[int]string {
	order := make(map[int]string)
	seen := make(map[string]bool)
	var visitAll func(items []string)

	count := 0
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[count] = item
				count++
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}
