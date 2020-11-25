package tempConv

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test_ToF test for ToF, to run
//   cd chapter2
//   go test -run Test_ToF -v
func Test_ToF(t *testing.T) {
	boilingPointF := BoilingPointC.ToF().String()
	absoluteZeroF := AbsoluteZeroC.ToF().String()
	freezingF := FreezingPointC.ToF().String()
	assert.Equal(t, "212.00°F", boilingPointF)
	assert.Equal(t, "32.00°F", freezingF)
	assert.Equal(t, "-459.67°F", absoluteZeroF)
	t.Logf("BP = %s, FP = %s, Abs = %s", boilingPointF, freezingF, absoluteZeroF)
}

// Test_ToC test for ToC, to run
//   cd chapter2
//   go test -run  Test_ToC -v
func Test_ToC(t *testing.T) {
	boilingPointC := BoilingPointF.ToC().String()
	absoluteZeroC := AbsoluteZeroF.ToC().String()
	freezingC := FreezingPointF.ToC().String()
	assert.Equal(t, "100.00°C", boilingPointC)
	assert.Equal(t, "0.00°C", freezingC)
	assert.Equal(t, "-273.15°C", absoluteZeroC)
	t.Logf("BP = %s, FP = %s, Abs = %s", boilingPointC, freezingC, absoluteZeroC)
}

// ExampleFahrenheit_ToC
func ExampleFahrenheit_ToC() {
	boilingPointC := BoilingPointF.ToC().String()
	absoluteZeroC := AbsoluteZeroF.ToC().String()
	freezingC := FreezingPointF.ToC().String()
	fmt.Println(boilingPointC)
	fmt.Println(absoluteZeroC)
	fmt.Println(freezingC)
	// Output:
	// 100.00°C
	// -273.15°C
	// 0.00°C
}
