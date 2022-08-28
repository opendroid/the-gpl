package chapter3

import (
	"bytes"
	"testing"
)

// TestPlotOn3DSurface tests the 3D plot function
//
//	go test -run TestPlotOn3DSurface -v
func TestPlotOn3DSurface(t *testing.T) {
	t.Parallel()
	t.Run("Sinc Plot", func(t *testing.T) {
		var buf bytes.Buffer
		PlotOn3DSurface(&buf, Sinc)
		if buf.Len() == 0 {
			t.Errorf("EggHandlerSVG failed to return data.")
		}
	})
	t.Run("Egg Plot", func(t *testing.T) {
		var buf bytes.Buffer
		PlotOn3DSurface(&buf, Egg)
		if buf.Len() == 0 {
			t.Errorf("EggHandlerSVG failed to return data.")
		}
	})
}

func TestEggHandlerSVG(t *testing.T) {
	var buf bytes.Buffer
	EggHandlerSVG(&buf)
	if buf.Len() == 0 {
		t.Errorf("EggHandlerSVG failed to return data.")
	}
}

func TestSincSVG(t *testing.T) {
	var buf bytes.Buffer
	SincSVG(&buf)
	if buf.Len() == 0 {
		t.Errorf("SincSVG failed to return data.")
	}
}

func TestValleyHandlerSVG(t *testing.T) {
	var buf bytes.Buffer
	ValleyHandlerSVG(&buf)
	if buf.Len() == 0 {
		t.Errorf("ValleyHandlerSVG failed to return data.")
	}
}

func TestSquaresHandlerSVG(t *testing.T) {
	var buf bytes.Buffer
	SquaresHandlerSVG(&buf)
	if buf.Len() == 0 {
		t.Errorf("SquaresHandlerSVG failed to return data.")
	}
}
