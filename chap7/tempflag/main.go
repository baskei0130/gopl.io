package main

import (
	"flag"
	"fmt"

	"github.com/baskei0130/gopl.io/chap7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
