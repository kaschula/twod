package color

import (
	"fmt"
	"image/color"
)

var (
	White  = color.White
	Black  = color.Black
	Red    = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	Purple = color.RGBA{R: 255, G: 0, B: 255, A: 255}
	Blue   = color.RGBA{R: 0, G: 0, B: 255, A: 255}
	Green  = color.RGBA{R: 0, G: 255, B: 0, A: 255}
)

// RGBA  receives values between 0-255
func RGBA(r, g, b, a uint8) color.Color {
	if r < 0 || r > 255 {
		panic(fmt.Sprintf("red value is out of range %v", r))
	}

	if b < 0 || b > 255 {
		panic(fmt.Sprintf("blue value is out of range %v", b))
	}
	if g < 0 || g > 255 {
		panic(fmt.Sprintf("green value is out of range %v", g))
	}

	if a < 0 || a > 255 {
		panic(fmt.Sprintf("alpha value is out of range %v", a))
	}

	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

func WithAlpha(c color.RGBA, alpha uint8) color.Color {
	if alpha < 0 || alpha > 255 {
		panic(fmt.Sprintf("alpha value is out of range %v", alpha))
	}

	return color.RGBA{
		R: c.R,
		G: c.G,
		B: c.B,
		A: alpha,
	}
}
