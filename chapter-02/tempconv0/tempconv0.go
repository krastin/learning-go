// Package tempconv0 performs Celsius and Fahrenheit temperature computations.
package tempconv0

import "fmt"

// Celsius Degrees
type Celsius float64

// Fahrenheit Degrees
type Fahrenheit float64

const (
	// AbsoluteZeroC Absolute zero in c
	AbsoluteZeroC Celsius = -273.15

	// FreezingC Freezing temp in c
	FreezingC     Celsius = 0

	// BoilingC Boiling temp in c
	BoilingC      Celsius = 100
)

// CToF c to f degrees
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC f to c degrees
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	// fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)          // "true"
	fmt.Println(f >= 0)          // "true"
	//fmt.Println(c == f)          // compile error: type mismatch
	fmt.Println(c == Celsius(f)) // "true"!

	c2 := FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c2)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c2)   // "100°C"
	fmt.Println(c2)          // "100°C"
	fmt.Printf("%g\n", c2)   // "100"; does not call String
	fmt.Println(float64(c2)) // "100"; does not call String
}