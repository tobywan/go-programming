// Mandelbrot emits a PNG of the mandlebrot set

package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		//xmin, ymin, xmax, ymax = -2, -2, +2, +2
		// xmin, ymin, xmax, ymax = -2, -0.5, -1, 0.5
		// xmin, ymin, xmax, ymax = -1.5, 0, -1, 0.5
		xmin, ymin, xmax, ymax = -1.25, 0.25, -1, 0.5
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin // The y val -2 <= y < 2
		for px := 0; px < width; px++ {
			x := float64(px)/height*(xmax-xmin) + xmin // The x val -2 <= x < 2
			z := complex(x, y)
			img.Set(px, py, mandlebrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandlebrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}

	}
	return color.Black
}
