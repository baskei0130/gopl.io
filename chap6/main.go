package main

import (
	"fmt"

	"github.com/baskei0130/gopl.io/chap6/geometry"
)

func main() {
	perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(perim.Distance())
}
