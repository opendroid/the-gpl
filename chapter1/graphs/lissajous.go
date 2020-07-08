package graphs

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

// LissajousConfig configuration object for Lissajous curve
type LissajousConfig struct {
	// Cycles : Number of oscillations
	Cycles int
	// Resolution Angular resolution
	Resolution float64
	// Size Canvas size [-Size .. +Size]
	Size int
	// NFrames Number of animationFrames
	NFrames int
	// DelayMS Delay in 10 ms units
	DelayMS int
}

var palette = []color.Color{
	color.White,
	color.RGBA{0xff, 0x00, 0x00, 0xFF}, // rgb(255, 0, 0) Red
	color.RGBA{0x00, 0xff, 00, 0xFF},   // rgb(0, 255, 0) Green
	color.RGBA{0x00, 0x00, 0xff, 0xFF}, // rgb(0, 0, 255) Blue
	color.RGBA{0xff, 0xff, 0x00, 0xFF}, // rgb(255, 255, 0) Yellow
	color.Black,
}

// Lissajous curve: x = A sin(at+d), y = B sin(bt),
//     https://en.wikipedia.org/wiki/Lissajous_curve
func Lissajous(config LissajousConfig, out io.Writer) {
	frequency := rand.Float64() * 3.0 // Freq of y oscillator
	phase := 0.0
	animation := gif.GIF{LoopCount: config.NFrames}

	// Get all frames
	paletteIndex := 0
	for i := 0; i < config.NFrames; i++ {
		// Define image and a index in color Palette
		rect := image.Rect(0, 0, 2*config.Size, 2*config.Size)
		img := image.NewPaletted(rect, palette)
		paletteIndex++
		if paletteIndex >= len(palette) {
			paletteIndex = 1
		}

		for angel := 0.0; angel < float64(config.Cycles)*2*math.Pi; angel += config.Resolution {
			x := math.Sin(angel*frequency + phase)
			coordY := config.Size + int(x*float64(config.Size)-0.5)
			y := math.Sin(angel)
			coordX := config.Size + int(y*float64(config.Size)-0.5)
			img.SetColorIndex(coordX, coordY, uint8(paletteIndex))
		}
		phase += 0.1
		animation.Delay = append(animation.Delay, config.DelayMS)
		animation.Image = append(animation.Image, img)
	}
	_ = gif.EncodeAll(out, &animation) // Ignore error
}

// Lissajous creates a graphs with a config
func (config LissajousConfig) Lissajous(out io.Writer) {
	frequency := rand.Float64() * 3.0 // Freq of y oscillator
	phase := 0.0
	animation := gif.GIF{LoopCount: config.NFrames}

	// Get all frames
	paletteIndex := 0
	for i := 0; i < config.NFrames; i++ {
		// Define image and a index in color Palette
		rect := image.Rect(0, 0, 2*config.Size, 2*config.Size)
		img := image.NewPaletted(rect, palette)
		paletteIndex++
		if paletteIndex >= len(palette) {
			paletteIndex = 1
		}

		for angel := 0.0; angel < float64(config.Cycles)*2*math.Pi; angel += config.Resolution {
			x := math.Sin(angel*frequency + phase)
			coordY := config.Size + int(x*float64(config.Size)-0.5)
			y := math.Sin(angel)
			coordX := config.Size + int(y*float64(config.Size)-0.5)
			img.SetColorIndex(coordX, coordY, uint8(paletteIndex))
		}
		phase += 0.1
		animation.Delay = append(animation.Delay, config.DelayMS)
		animation.Image = append(animation.Image, img)
	}
	_ = gif.EncodeAll(out, &animation) // Ignore error
}

// DefaultLissajous calls a Lissajous figure with default data
//    2 cycles with 512x512 size, 12 frames and 10 ms delay
func DefaultLissajous(out io.Writer) {
	config := LissajousConfig{
		Cycles:     2,
		Resolution: 0.000001,
		Size:       512,
		NFrames:    12,
		DelayMS:    10,
	}
	Lissajous(config, out)
}
