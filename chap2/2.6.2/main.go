package main

import (
	"fmt"

	"gopl.io/chap2/2.6.2/popcount"
)

func main() {
	fmt.Print("Hello World\n")
	popcount.PopCount4(uint64(63))
	popcount.PopCount1(uint64(63))
}
