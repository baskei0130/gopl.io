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
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

/*
Ans. f(z) = z^4 - 1 = 0
z_n+1 = z_n - f(z_n)/f'(z_n)
      = z_n - z_n^4-1/4z_n^3
	  = z_n - (z_n - 1/z_n^3) / 4
*/
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7

	for n := uint8(0); n < iterations; n++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z) < 1e-6 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
