package chapter2

import (
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
	assert.Equal(t, "212.00℉", boilingPointF)
	assert.Equal(t, "32.00℉", freezingF)
	assert.Equal(t, "-459.67℉", absoluteZeroF)
	t.Logf("BP = %v, FP = %s, Abs = %s", boilingPointF, freezingF, absoluteZeroF)
}

// Test_ToC test for ToC, to run
//   cd chapter2
//   go -run test Test_ToC -v
func Test_ToC(t *testing.T) {
	boilingPointC := BoilingPointF.ToC().String()
	absoluteZeroC := AbsoluteZeroF.ToC().String()
	freezingC := FreezingPointF.ToC().String()
	assert.Equal(t, "100.00℃", boilingPointC)
	assert.Equal(t, "0.00℃", freezingC)
	assert.Equal(t, "-273.15℃", absoluteZeroC)
	t.Logf("BP = %v, FP = %s, Abs = %s", boilingPointC, freezingC, absoluteZeroC)
}
