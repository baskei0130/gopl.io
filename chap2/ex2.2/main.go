package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"gopl.io/chap2/ex2.2/unitconv"
)

var (
	t = flag.Bool("t", false, "temperature")
	l = flag.Bool("l", false, "length")
	w = flag.Bool("w", false, "weight")
)

func main() {
	flag.Parse()

	if flag.NFlag() > 0 {
		arg := flag.Arg(0)
		val, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fail to ParseFloat: %s", arg)
			os.Exit(1)
		}

		printConv(val, *t, *l, *w)
	} else {
		var opt string
		var val float64

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Printf("input type [t, l, w]: ")
		if scanner.Scan() {
			opt = scanner.Text()
		}

		fmt.Printf("input value: ")
		var tmp string
		if scanner.Scan() {
			tmp = scanner.Text()
		}
		val, err := strconv.ParseFloat(tmp, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fail to ParseFloat: %s", tmp)
			os.Exit(1)
		}

		printConv(val, opt == "t", opt == "l", opt == "w")
	}
}

func printConv(val float64, t, l, w bool) {
	if t {
		f := unitconv.Fahrenheit(val)
		c := unitconv.Celsious(val)
		fmt.Printf("%s = %s, %s = %s\n", f, unitconv.FtoC(f), c, unitconv.CtoF(c))
	}
	if l {
		f := unitconv.Feet(val)
		m := unitconv.Metre(val)
		fmt.Printf("%s = %s, %s = %s\n", f, unitconv.FtoM(f), m, unitconv.MtoF(m))
	}
	if w {
		p := unitconv.Pound(val)
		k := unitconv.Kilogram(val)
		fmt.Printf("%s = %s, %s = %s\n", p, unitconv.PtoK(p), k, unitconv.KtoP(k))
	}
}
