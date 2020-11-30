package chapter3

import (
	"bytes"
	"testing"
)

// TestPlotOn3DSurface tests the 3D plot function
//   go test -run TestPlotOn3DSurface -v
func TestPlotOn3DSurface(t *testing.T) {
	t.Parallel()
	t.Run("Sinc Plot", func(t *testing.T) {
		t.Skip("Skipping Sinc Plot")
		var buf bytes.Buffer
		PlotOn3DSurface(&buf, Sinc)
		t.Logf("%s", buf.String())
	})
	t.Run("Egg Plot", func(t *testing.T) {
		var buf bytes.Buffer
		PlotOn3DSurface(&buf, Egg)
		t.Logf("%s", buf.String())
	})
}
