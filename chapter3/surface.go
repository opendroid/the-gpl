package chapter3

import (
	"fmt"
	"io"
	"log/slog"
	"math"
)

var (
	sin30 = math.Sin(SurfaceAngle30)
	cos30 = math.Cos(SurfaceAngle30)
)

// PlotOn3DSurface plats z = f(x,y) as a wire on a 3-D mesh surface using SVG (Scalable Vector Graphics)
//
//	See SVG on https://www.w3schools.com/graphics/svg_intro.asp
//	Example, makes a hexagon
//	<svg width="100" height="100">
//	  <polygon points="25,0 50,0 75,25 50,50 25,50 0,25" stroke="green" stroke-width="4" fill="yellow" />
//	</svg>
func PlotOn3DSurface(w io.Writer, plot func(float64, float64) float64) {
	_, err := fmt.Fprintf(w, SVGPrefixFormat, SurfaceWidth, SurfaceHeight)
	// close SVG tag in all returns
	defer (func() { _, _ = fmt.Fprintf(w, "\n%s\n", SVGSuffixTag) })()
	if err != nil {
		_ = fmt.Errorf("PlotOn3DSurface: Error: %s", err)
		return
	}

	// Make polygons in each cell.
	for i := 0; i < SurfaceGridCells; i++ {
		for j := 0; j < SurfaceGridCells; j++ {
			ax, ay, err := corner(i+1, j, plot)
			if err != nil {
				continue
			}
			bx, by, err := corner(i, j, plot)
			if err != nil {
				continue
			}
			cx, cy, err := corner(i, j+1, plot)
			if err != nil {
				continue
			}
			dx, dy, err := corner(i+1, j+1, plot)
			if err != nil {
				continue
			}
			_, err = fmt.Fprintf(w, "<polygon points=\"%g,%g %g,%g %g,%g %g,%g\"/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			if err != nil {
				slog.Error("PlotOn3DSurface: polygon point Error", "err", err)
				return
			}
		}
	}
}

func corner(i, j int, plot func(float64, float64) float64) (float64, float64, error) {
	// Translate from (SurfaceGridCells x SurfaceGridCells) => (SurfaceXYRange x SurfaceXYRange)
	x := SurfaceXYRange * (float64(i)/SurfaceGridCells - 0.5)
	y := SurfaceXYRange * (float64(j)/SurfaceGridCells - 0.5)
	z := plot(x, y) // compute surface height, z
	if math.IsInf(z, 0) {
		return 0, 0, fmt.Errorf("invalid IsInf polygon at (%d, %d), ignore", i, j)
	}
	if math.IsNaN(z) {
		return 0, 0, fmt.Errorf("invalid IsNaN polygon at (%d, %d), ignore", i, j)
	}
	// Project (x, y, z) to an isometric 2-D SCG canvas
	sx := SurfaceWidth/2 + (x-y)*cos30*SurfaceXYScale
	sy := SurfaceHeight/2 + (x+y)*sin30*SurfaceXYScale - z*SurfaceZScale
	return sx, sy, nil
}

// Sinc sampling function returns 1 or sin r / r
//
//	https://mathworld.wolfram.com/SincFunction.html
func Sinc(x, y float64) float64 {
	r := math.Hypot(x, y) // Distance of (x,y) from (0,0)
	k := math.Sin(r) / r  // recover from a divide by zero error

	return k
}

// Squares of sorts
func Squares(x, y float64) float64 {
	return math.Pow(math.Sin(x/math.Pi)+math.Cos(y/math.Pi), 2) / SquaresDenominator
}

// Valley of sorts
func Valley(x, y float64) float64 {
	r := math.Hypot(x, y) // Distance of (x,y) from (0,0)
	return math.Sin(-y) * math.Pow(2, -r)
}

// Egg sampling function returns squares
func Egg(x, y float64) float64 {
	return math.Pow(2, math.Sin(x)) * math.Pow(2, math.Cos(y)) / EggDenominator
}

// EggHandlerSVG draws an egg on a writer
func EggHandlerSVG(w io.Writer) {
	PlotOn3DSurface(w, Egg)
}

// SincSVG writes raw SVG content
func SincSVG(w io.Writer) {
	PlotOn3DSurface(w, Sinc)
}

// ValleyHandlerSVG draws a Valley on a writer
func ValleyHandlerSVG(w io.Writer) {
	PlotOn3DSurface(w, Valley)
}

// SquaresHandlerSVG draws a sinc on a writer
func SquaresHandlerSVG(w io.Writer) {
	PlotOn3DSurface(w, Squares)
}
