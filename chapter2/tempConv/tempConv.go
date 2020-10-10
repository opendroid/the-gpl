package tempConv

import (
	"fmt"
)

// Celsius unit for temperature
type Celsius float64

// Fahrenheit unit for temperature
type Fahrenheit float64

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
)

// ToC converts a Fahrenheit to °C
func (f Fahrenheit) ToC() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// ToF converts a Celsius to °F
func (c Celsius) ToF() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
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
