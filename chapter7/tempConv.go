package chapter7

import (
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/chapter2/tempConv"
)
// tempConv:
//  Satisfies the flag value interface for custom unit temperature, that
//  reads a temp value in form of -173.15°T or -173.15°t or -173.15T or -173.15t
//	where T is C or F or K
//  and converts it to Celsius or Fahrenheit or Kelvin

// celsiusFlag defines value to be converted to Celsius, satisfies flag Set interface
type celsiusFlag struct {
	tempConv.Celsius
}

// fahrenheitFlag defines value to be converted to Fahrenheit, satisfies flag Set interface
type fahrenheitFlag struct {
	tempConv.Fahrenheit
}

// Kelvin defines value to be converted to Kelvin, satisfies flag Set interface
type kelvinFlag struct {
	tempConv.Kelvin
}

// Set satisfy the flag interface for celsiusFlag. Convert to °C
func (c *celsiusFlag) Set (s string) error {
	var unit string
	var value float64
	_, _ = fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "°F", "F", "f", "°f":
		c.Celsius = tempConv.Fahrenheit(value).ToC()
		return nil
	case "°C", "C", "c", "°c":
		c.Celsius = tempConv.Celsius(value)
		return nil
	case "°K", "K", "k", "°k":
		c.Celsius = tempConv.Kelvin(value).ToC()
		return nil
	}

	return fmt.Errorf("invalid temperature %q. Should be %.2f°C or °F or °K", s, value)
}

// CelsiusFlag creates a custom flag
func CelsiusFlag(name string, value tempConv.Celsius, usage string, set *flag.FlagSet) *tempConv.Celsius {
	v := celsiusFlag{value}
	set.Var(&v, name, usage)
	return &v.Celsius
}

// Set satisfy the flag interface for fahrenheitFlag, Convert to °F
func (f *fahrenheitFlag) Set (s string) error {
	var unit string
	var value float64
	_, _ = fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "°F", "F", "f", "°f":
		f.Fahrenheit = tempConv.Fahrenheit(value)
		return nil
	case "°C", "C", "c", "°c":
		f.Fahrenheit = tempConv.Celsius(value).ToF()
		return nil
	case "°K", "K", "k", "°k":
		f.Fahrenheit = tempConv.Kelvin(value).ToF()
		return nil
	}
	return fmt.Errorf("invalid temperature %q. Should be %.2f°C or °F or °K", s, value)
}

// FahrenheitFlag creates a custom flag
func FahrenheitFlag(name string, value tempConv.Fahrenheit, usage string, set *flag.FlagSet) *tempConv.Fahrenheit {
	v := fahrenheitFlag{value}
	set.Var(&v, name, usage)
	return &v.Fahrenheit
}

// Set satisfy the flag interface for kelvinFlag, Convert to °K
func (k *kelvinFlag) Set (s string) error {
	var unit string
	var value float64
	_, _ = fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "°F", "F", "f", "°f":
		k.Kelvin = tempConv.Fahrenheit(value).ToK()
		return nil
	case "°C", "C", "c", "°c":
		k.Kelvin = tempConv.Celsius(value).ToK()
		return nil
	case "°K", "K", "k", "°k":
		k.Kelvin = tempConv.Kelvin(value)
		return nil
	}
	return fmt.Errorf("invalid temperature %q. Should be %.2f°C or °F or °K", s, value)
}

// KelvinFlag creates a custom flag
func KelvinFlag(name string, value tempConv.Kelvin, usage string, set *flag.FlagSet) *tempConv.Kelvin {
	v := kelvinFlag{value}
	set.Var(&v, name, usage)
	return &v.Kelvin
}