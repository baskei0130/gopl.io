package main

import (
	"fmt"
	"math"
)

func main() {
	ex1()
}

func ex1() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}
