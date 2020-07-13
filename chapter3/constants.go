package chapter3

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
