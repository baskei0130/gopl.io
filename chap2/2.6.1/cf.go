// 数値引数を摂氏と華氏へ変換する
package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/chap2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsious(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FtoC(f), c, tempconv.CtoF(c))
	}
}
