// Package tempConv provides conversion among Celsius, Kelvin and Fahrenheit using methods.
package tempConv

import (
	"fmt"
)

// Celsius unit for temperature
type Celsius float64

// Fahrenheit unit for temperature
type Fahrenheit float64

// Fahrenheit unit for temperature
type Kelvin float64

const (
	// AbsoluteZeroC in celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingPointC at which water freezes in celsius
	FreezingPointC Celsius = 0
	// BoilingPointC water boils in celsius
	BoilingPointC Celsius = 100
	// AbsoluteZeroF in fahrenheit
	AbsoluteZeroF Fahrenheit = -459.67
	// FreezingPointF in fahrenheit
	FreezingPointF Fahrenheit = 32
	// BoilingPointF water boils in fahrenheit
	BoilingPointF Fahrenheit = 212
	// AbsoluteZeroK in Kelvin
	AbsoluteZeroK Kelvin = -273.15
	// FreezingPointK in Kelvin
	FreezingPointK Kelvin = 273.15
	// BoilingPointK water boils in Kelvin
	BoilingPointK Kelvin = 373.15
)

// Suppress unused warning errors
var (
	_ = AbsoluteZeroK
	_ = BoilingPointK
)

// ToC converts a Fahrenheit to °C
func (f Fahrenheit) ToC() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// ToK converts a Fahrenheit to °K
func (f Fahrenheit) ToK() Kelvin {
	return Kelvin(f.ToC()) + FreezingPointK
}

// ToF converts a Celsius to °F
func (c Celsius) ToF() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// ToK converts a Celsius to °K
func (c Celsius) ToK() Kelvin {
	return Kelvin(c) + FreezingPointK
}

// ToF converts a Kelvin to °F
func (k Kelvin) ToF() Fahrenheit {
	return Fahrenheit((k-FreezingPointK)*9/5 + 32)
}

// ToC converts a Kelvin to °C
func (k Kelvin) ToC() Celsius {
	return Celsius(k - FreezingPointK)
}

// String prints in format 100°F value
//  Can use to print value in %s fmt.Printf format
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", f)
}

// String prints in format 100°C value
//  Can use in %s fmt.Printf format
func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

// String prints in format 100°K value
//  Can use in %s fmt.Printf format
func (k Kelvin) String() string {
	return fmt.Sprintf("%.2f°K", k)
}
