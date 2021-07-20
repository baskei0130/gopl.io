// tempcong パッケージは摂氏と華氏の温度変換を行う
package tempconv

import (
	"flag"
	"fmt"
)

type Celsious float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsious = -273.15
	FreezingC     Celsious = 0
	BoilingC      Celsious = 100
)

func (c Celsious) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func (k Kelvin) String() string { return fmt.Sprintf("%gK", k) }

type celsiusFlag struct{ Celsious }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C":
		f.Celsious = Celsious(value)
		return nil
	case "F":
		f.Celsious = FtoC(Fahrenheit(value))
		return nil
	case "K":
		f.Celsious = KtoC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsious, usage string) *Celsious {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsious
}
