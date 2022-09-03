// Package lissajous outputs a .png image to a writer.
package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
)

// Config configuration object for Lissajous curve
type Config struct {
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
	color.RGBA{R: 0xff, A: 0xFF},          // rgb(255, 0, 0) Red
	color.RGBA{G: 0xff, A: 0xFF},          // rgb(0, 255, 0) Green
	color.RGBA{R: 0xff, G: 0xff, A: 0xFF}, // rgb(255, 255, 0) Yellow
	color.RGBA{B: 0xff, A: 0xFF},          // rgb(0, 0, 255) Blue
	color.RGBA{R: 0xff, B: 0xff, A: 0xFF}, // rgb(255, 0, 255) Magenta
	color.RGBA{R: 0x7f, B: 0x7f, A: 0xFF}, // rgb(128, 0, 128) Purple
	color.RGBA{R: 0x94, B: 0xD3, A: 0xFF}, // rgb(148, 0, 211) darkviolet
	color.Black,
}

// Lissajous curve: x = A sin(at+d), y = B sin(bt),
//
//	https://en.wikipedia.org/wiki/Lissajous_curve
func Lissajous(w io.Writer, config Config) {
	frequency := rand.Float64() * 3.0 // Freq of y oscillator
	phase := 0.0
	animation := gif.GIF{LoopCount: 0} // Loops forever

	// Get all frames
	paletteIndex := 0
	for i := 0; i < config.NFrames; i++ {
		// Define image and an index in color Palette
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
	_ = gif.EncodeAll(w, &animation) // Ignore error
}

// Lissajous creates a lissajous with a config
func (config Config) Lissajous(out io.Writer) {
	frequency := rand.Float64() * 3.0 // Freq of y oscillator
	phase := 0.0
	animation := gif.GIF{LoopCount: config.NFrames}

	// Get all frames
	paletteIndex := 0
	for i := 0; i < config.NFrames; i++ {
		// Define image and an index in color Palette
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

// Default calls a Lissajous figure with default data
//
//	4 cycles with 512x512 size, 6 frames and 20 ms delay
func Default(w io.Writer) {
	config := Config{
		Cycles:     4,
		Resolution: 0.001,
		Size:       512,
		NFrames:    6,
		DelayMS:    20,
	}
	Lissajous(w, config)
}

// Figure displays public images
func Figure(w http.ResponseWriter, _ *http.Request) {
	Default(w)
}
