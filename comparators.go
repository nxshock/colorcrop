package colorcrop

import (
	"image/color"
	"math"
)

// comparator is a function that returns a difference between two colors in
// range 0.0..1.0.
type comparator func(color.Color, color.Color) float64

// CmpColorDifference returns difference of two colors
func CmpSquareRGBComponentsDiff(color1 color.Color, color2 color.Color) float64 {
	const maxDiff = 113509.94967402637 // Difference between black and white colors

	r1, g1, b1, _ := color1.RGBA()
	r2, g2, b2, _ := color2.RGBA()
	return math.Sqrt(math.Pow(float64(r2)-float64(r1), 2.0)+
		math.Pow(float64(g2)-float64(g1), 2.0)+
		math.Pow(float64(b2)-float64(b1), 2.0)) / maxDiff
}

// CmpColorDifference returns difference of two colors.
func CmpRGBComponentsDiff(color1 color.Color, color2 color.Color) float64 {
	const maxDiff = 765 // Difference between black and white colors

	r1, g1, b1, _ := color1.RGBA()
	r2, g2, b2, _ := color2.RGBA()
	return math.Sqrt(math.Abs(float64(r2)-float64(r1))+
		math.Abs(float64(g2)-float64(g1))+
		math.Abs(float64(b2)-float64(b1))) / maxDiff
}
