package chapter3

import (
	"math"
)

// HTML
const (
	HTMLBeginEgg     = `<!DOCTYPE html><html lang="en"><body><h1>Egg</h1>`
	HTMLBeginSinc    = `<!DOCTYPE html><html lang="en"><body><h1>Sinc</h1>`
	HTMLBeginValley  = `<!DOCTYPE html><html lang="en"><body><h1>Valley</h1>`
	HTMLBeginSquares = `<!DOCTYPE html><html lang="en"><body><h1>Squares</h1>`
	HTMLEnd          = "</body></html>"
)

// Mandelbrot Set Graph constants
const (
	// MBXMin is lower bound of x value for MB (Mandelbrot Set)
	MBXMin = -2
	// MBXMax is upper bound of x value for MB (Mandelbrot Set)
	MBXMax = 2
	// MBYMin is lower bound of y value for MB (Mandelbrot Set)
	MBYMin = -2
	// MBYMax is upper bound of y value for MB (Mandelbrot Set)
	MBYMax = 2
	// MBWidth pixels wide (along x-axis)
	MBWidth = 1024
	// MBHeight pixels high (along y-axis)
	MBHeight = 1024
	// MBContrast contrast desired
	MBContrast = 100
	// MBSubPixels number of pixels to take average for super aliasing
	MBSubPixels = 4
	// MaxGoRoutines for concurrency
	MaxGoRoutines = 32
)

type MandelbrotImage int

const (
	// MBBlackAndWhite draws MB in Black and White
	MBBlackAndWhite MandelbrotImage = iota
	MBColor
)

type colorComponents struct {
	red, green, blue float64
}

// Surface plot 3D (x,y,z) constants to plot on a 2-D graph
const (
	// SurfaceWidth width of #D surface plot
	SurfaceWidth = 1200
	// SurfaceHeight Height of #D surface plot
	SurfaceHeight = 800
	// GridCells number of cells in a grid
	SurfaceGridCells = 50
	// SurfaceXYRange of x-axis from -SurfaceXYRange .. +SurfaceXYRange
	SurfaceXYRange = 30.0
	// SurfaceXYScale scaling x data to plot on graph
	SurfaceXYScale = SurfaceWidth / 2 / SurfaceXYRange
	// SurfaceZScale scaling x data to plot on graph
	SurfaceZScale = SurfaceHeight * 0.4
	// SurfaceAngle at which x and y are angled
	SurfaceAngle30 = math.Pi / 6
	// SVGPrefix element prefix
	SVGPrefixFormat = `<svg xmlns="http://www.w3/org/2000/svg" style="stroke: grey; fill: white; stroke-width:0.7" width="100%%" height="100%%" viewBox="0 0 %d %d">`
	// SVGSuffixTag closes <svg> ... </svg>
	SVGSuffixTag = "</svg>"
)

const (
	EggDenominator     = 10
	SquaresDenominator = 5
)
