package main

import (
	"fmt"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(vals ...int) int {
	max := -10000
	count := 0
	for _, val := range vals {
		if max < val {
			max = val
		}
		count++
	}
	if count == 0 {
		return -1
	}
	return max
}

func min(vals ...int) int {
	min := 10000000
	count := 0
	for _, val := range vals {
		if min > val {
			min = val
		}
		count++
	}
	if count == 0 {
		return -1
	}
	return min
}
func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))

	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
}
