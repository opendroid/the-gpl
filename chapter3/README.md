# Chapter 3 — Basic Data Types

Examples from Chapter 3 of *The Go Programming Language*, covering integers, floats,
complex numbers, strings, and constants — made visual through fractal and surface renderers.

## What's Here

| File / Package | Web route | What it does |
|---|---|---|
| `mandelbrot.go` | `/mandel`, `/mandelbw` | Mandelbrot fractal PNG (color and B&W) |
| `surface.go` | `/sinc`, `/egg`, `/valley`, `/sq` | 3-D wire-mesh surface plots as SVG |
| `tstrings/` | — | String utility functions |

## Mandelbrot Fractal

Renders the Mandelbrot set as a PNG image by iterating `z = z² + c` for each pixel
and mapping the iteration count to a colour (or black/white threshold).

```go
// HTTP handlers — registered in serve/web/web.go
chapter3.MBGraphHandler(w, r)   // colour PNG at /mandelimage.png
chapter3.MBGraphBWHandler(w, r) // B&W PNG at /mandelbwimage.png
```

- `SetColor(z complex128, i, maxIter int) color.RGBA` — maps iteration depth to an RGB colour.
- `MandelbrotImage` type: `MBColor` or `MBBW` enum selects the rendering mode.

View live: [the-gpl.com/mandel](https://the-gpl.com/mandel) · [the-gpl.com/mandelbw](https://the-gpl.com/mandelbw)

## 3-D Surface Plots

Plots `z = f(x, y)` as a wire-mesh SVG using an isometric projection.
Each surface function is a `func(float64, float64) float64` passed to `PlotOn3DSurface`.

```go
// Write a sinc SVG to stdout
chapter3.SincSVG(os.Stdout)

// Plot any custom function
chapter3.PlotOn3DSurface(os.Stdout, func(x, y float64) float64 {
    return math.Sin(x) * math.Cos(y)
})
```

| Function | Surface shape | Route |
|---|---|---|
| `Sinc(x, y float64) float64` | sin(r)/r ripple | `/sinc` |
| `Egg(x, y float64) float64` | egg-carton bumps | `/egg` |
| `Valley(x, y float64) float64` | valley / saddle | `/valley` |
| `Squares(x, y float64) float64` | square-wave grid | `/sq` |

SVG handlers (used by the web server):
`EggHandlerSVG`, `SincSVG`, `ValleyHandlerSVG`, `SquaresHandlerSVG` — each takes an `io.Writer`.

View live: [/sinc](https://the-gpl.com/sinc) · [/egg](https://the-gpl.com/egg) · [/valley](https://the-gpl.com/valley) · [/sq](https://the-gpl.com/sq)

## String Utilities (`tstrings/`)

```go
tstrings.Basename("a/b/c.go")         // "c"
tstrings.Comma("1234567")             // "1,234,567"
tstrings.CommaWithBuf("9876543")      // "9,876,543"  (bytes.Buffer version)
tstrings.IntsToString([]int{1, 2, 3}) // "[1, 2, 3]"
```

Also defines two interfaces:

```go
type DayOfWeekName interface {
    DayName(d time.Weekday) string
    Weekend(d time.Weekday) bool
    Weekday(d time.Weekday) bool
}

type MonthOfYearName interface {
    MonthName(m time.Month) string
}
```

## Running Tests

```bash
go test ./chapter3/...
go test -v ./chapter3/tstrings/...
```

## Viewing Locally

```bash
the-gpl server --port=8080
# Then open:
# http://localhost:8080/mandel
# http://localhost:8080/sinc
```

## Go Features Demonstrated

- `complex128` arithmetic for fractal iteration
- `image`, `image/color`, `image/png` standard library
- `math` functions (`Sin`, `Cos`, `Sqrt`, `Abs`, `IsInf`)
- SVG generation via `fmt.Fprintf` to `io.Writer`
- Named constants and `iota`
- `bytes.Buffer` vs string concatenation for performance
- Interface definitions with multiple methods
- `strings` and `strconv` standard library utilities
