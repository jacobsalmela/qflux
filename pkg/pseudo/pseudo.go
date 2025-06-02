package pseudo

import "image/color"

// lerp performs linear interpolation between two values
// a. starting value
// b. ending value
// t. interpolation factor (0.0 to 1.0)
func Lerp(a, b uint8, t float64) uint8 {
	return uint8(float64(a) + (float64(b)-float64(a))*t)
}

func LerpF(a, b, t float64) float64 {
	return a + (b-a)*t
}

// lerpColor performs linear interpolation between two colors
// c1. starting color
// c2. ending color
// t. interpolation factor (0.0 to 1.0)
func LerpColor(c1, c2 color.RGBA, t float64) color.RGBA {
	return color.RGBA{
		R: Lerp(c1.R, c2.R, t),
		G: Lerp(c1.G, c2.G, t),
		B: Lerp(c1.B, c2.B, t),
		A: Lerp(c1.A, c2.A, t),
	}
}
