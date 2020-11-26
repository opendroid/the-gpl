// Package chapter3 provides utilities for mandelbrot and surface plots.
package chapter3

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"net/http"
	"sync"
)

var maxAmp float64
var scaleX, scaleY func(float64) float64

// init variables in this file.
func init() {
	maxAmp = cmplx.Abs(complex(MBXMax, MBYMax))
	scaleX = scale(MBXMin, MBXMax, MBWidth)
	scaleY = scale(MBYMin, MBYMax, MBHeight)
}

// scale transforms a value to a corresponding value in range
func scale(min, max, r float64) func(float64) float64 {
	return func(x float64) float64 {
		return x/r*(max-min) + min
	}
}

// mandelbrotImage write a image MBWidth x MBHeight to writer
//	Exercise 8.5:
func mandelbrotImage(w io.Writer, b MandelbrotImage) {
	img := image.NewRGBA(image.Rect(0, 0, MBWidth, MBHeight))
	var wg sync.WaitGroup
	tokens := make(chan struct{}, MaxGoRoutines) // Limit parallelism
	for py := 0; py < MBHeight; py++ {
		y := scaleY(float64(py))
		wg.Add(1)
		go func(y float64, py int) {
			tokens <- struct{}{} // Wait to acquire a token
			for px := 0; px < MBWidth; px++ {
				x := scaleX(float64(px))
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z, b))
			}
			<-tokens  // release a token
			wg.Done() // Mark done
		}(y, py) //
	}
	wg.Wait() // Wait for all  to finish
	err := png.Encode(w, img)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Sorry could not draw Mandelbrot %v\n", err)
	}
}

// mandelbrot returns a color for a z
func mandelbrot(z complex128, b MandelbrotImage) color.Color {
	const n = 255 // Iterations
	const a = 2   // Amplitude
	var v complex128
	for i := uint8(0); i < n; i++ {
		v = v*v + z
		if cmplx.Abs(v) > a {
			switch b {
			case MBBlackAndWhite:
				return color.Gray{Y: 255 - MBContrast*i}
			case MBColor:
				return mbColor(z, i)
			}
		}
	}
	return color.Black
}

// MBGraphHandler
func MBGraphHandler(w http.ResponseWriter, _ *http.Request) { mandelbrotImage(w, MBColor) }

// MBGraphBWHandler in black
func MBGraphBWHandler(w http.ResponseWriter, _ *http.Request) { mandelbrotImage(w, MBBlackAndWhite) }

// SetColor total at subsequent z for ith iteration
func (cp *colorComponents) SetColor(z complex128, iteration float64) {
	f := cmplx.Abs(z) / maxAmp // fraction
	cp.blue += f * MBContrast * iteration
	cp.green += (1 - 0.5*f) * MBContrast * iteration
	cp.red += (1 - f) * MBContrast * iteration
}

// complexAtPixel get complex number at x,y pixels away from z in
// (MBWidth x MBHeight) space.
func complexAtPixel(z complex128) func(float64, float64) complex128 {
	return func(r float64, i float64) complex128 {
		zi := complex((MBXMax-MBXMin)*r/MBWidth, (MBYMax-MBXMin)*i/MBHeight)
		zp := z + zi
		return zp
	}
}

// mbColor makes a mandelbrot color for a complex z at ith iteration.
// it computes color as average of 4 nearest points to a given complex
// number at a pixel. For it to be different than original SetColor need
// to be non-linear so average does not cancel out
func mbColor(z complex128, i uint8) color.Color {
	var cp = new(colorComponents)
	iterations := float64(i)
	zb := complexAtPixel(z)
	zp := zb(-0.5, 0.5)
	cp.SetColor(zp, iterations)
	zp = zb(-0.5, -0.5)
	cp.SetColor(zp, iterations)
	zp = zb(0.5, -0.5)
	cp.SetColor(zp, iterations)
	zp = zb(0.5, 0.5)
	cp.SetColor(zp, iterations)

	// Take average of MBSubPixels points
	return color.RGBA{
		R: uint8(cp.red / MBSubPixels),
		G: uint8(cp.green / MBSubPixels),
		B: uint8(cp.blue / MBSubPixels),
		A: 255,
	}
}
