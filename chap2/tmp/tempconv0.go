package tempconv0

type Celsious float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsious = -273.15
	FreezingC     Celsious = 0
	BoilingC      Celsious = 100
)

func CtoF(c Celsious) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsious { return Celsious((f - 32) * 5 / 9) }
