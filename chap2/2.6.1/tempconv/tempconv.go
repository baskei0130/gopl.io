// tempcong パッケージは摂氏と華氏の温度変換を行う
package tempconv

import "fmt"

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
