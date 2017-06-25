package colorcrop

import (
	"image/color"
	"math"
)

// comparator is a function that returns a difference between two colors in
// range 0.0..1.0 (0.0 - same colors, 1.0 - totally different colors).
type comparator func(color.Color, color.Color) float64

// CmpEuclidean returns Euclidean difference of two colors.
// https://en.wikipedia.org/wiki/Color_difference#Euclidean
func CmpEuclidean(color1 color.Color, color2 color.Color) float64 {
	const maxDiff = 113509.94967402637 // Difference between black and white colors

	r1, g1, b1, _ := color1.RGBA()
	r2, g2, b2, _ := color2.RGBA()

	return math.Sqrt(distance(float64(r2), float64(r1))+
		distance(float64(g2), float64(g1))+
		distance(float64(b2), float64(b1))) / maxDiff
}

// CmpRGBComponents returns RGB components difference of two colors.
func CmpRGBComponents(color1 color.Color, color2 color.Color) float64 {
	const maxDiff = 195075.0 // Difference between black and white colors

	r1, g1, b1, _ := color1.RGBA()
	r2, g2, b2, _ := color2.RGBA()

	return (math.Abs(float64(r2)-float64(r1)) +
		math.Abs(float64(g2)-float64(g1)) +
		math.Abs(float64(b2)-float64(b1))) / maxDiff
}

// CmpCIE76 returns difference of two colors defined in CIE76 standart.
// https://en.wikipedia.org/wiki/Color_difference#CIE76
func CmpCIE76(color1 color.Color, color2 color.Color) float64 {
	const maxDiff = 150.8463301377893 // Difference between blue and white colors

	r1, g1, b1, _ := color1.RGBA()
	r2, g2, b2, _ := color2.RGBA()

	cl1, ca1, cb1 := xyztoLAB(rgbtoXYZ(r1/255, g1/255, b1/255))
	cl2, ca2, cb2 := xyztoLAB(rgbtoXYZ(r2/255, g2/255, b2/255))

	return math.Sqrt(distance(cl2, cl1) + distance(ca2, ca1) + distance(cb2, cb1))
}

func distance(x, y float64) float64 {
	return (x - y) * (x - y)
}
