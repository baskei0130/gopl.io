package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println(diffCount(c1, c2))
}

func diffCount(c1, c2 [32]byte) int {
	//var pc1 [256]int
	//var pc2 [256]int
	var sum int
	for i := 0; i < 32; i++ {
		sum += int(pc[c1[i]^c2[i]])
	}
	return sum
}
