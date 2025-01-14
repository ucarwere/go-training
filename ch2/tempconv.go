package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gF", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}

func KToC(k Kelvin) Celsius {
	return Celsius(k) + AbsoluteZeroC
}

func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}

func main() {
	fmt.Printf("0K = %gºC\n", KToC(0))
	fmt.Printf("0ºC = %gK\n", CToK(0))
	fmt.Printf("0ºF = %gK\n", FToK(0))
}
