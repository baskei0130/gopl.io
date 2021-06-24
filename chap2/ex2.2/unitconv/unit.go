package unitconv

func CtoF(c Celsious) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celsious { return Celsious((f - 32) * 5 / 9) }

func CtoK(c Celsious) Kelvin { return Kelvin(c + 273.15) }
func KtoC(k Kelvin) Celsious { return Celsious(k - 273.15) }

func FtoM(f Feet) Metre { return Metre(f * 0.3048) }
func MtoF(m Metre) Feet { return Feet(m / 0.3048) }

func PtoK(p Pound) Kilogram { return Kilogram(p / 2.20462) }
func KtoP(k Kilogram) Pound { return Pound(k * 2.20462) }
