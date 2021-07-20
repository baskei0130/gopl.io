package tempconv

func CtoF(c Celsious) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FtoC(f Fahrenheit) Celsious { return Celsious((f - 32) * 5 / 9) }

func CtoK(c Celsious) Kelvin { return Kelvin(c + 273.15) }

func KtoC(k Kelvin) Celsious { return Celsious(k - 273.15) }
