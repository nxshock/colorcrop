package colorcrop

import (
	"image/color"
	"math"
)

// colorToXYZ returns CIE XYZ representation of color.
// https://en.wikipedia.org/wiki/Color_model#CIE_XYZ_color_space
func colorToXYZ(color color.Color) (x, y, z float64) {
	r, g, b, _ := color.RGBA()
	varR := float64(r>>8) / 255
	varG := float64(g>>8) / 255
	varB := float64(b>>8) / 255

	if varR > 0.04045 {
		varR = math.Pow((varR+0.055)/1.055, 2.4)
	} else {
		varR = varR / 12.92
	}

	if varG > 0.04045 {
		varG = math.Pow((varG+0.055)/1.055, 2.4)
	} else {
		varG = varG / 12.92
	}

	if varB > 0.04045 {
		varB = math.Pow((varB+0.055)/1.055, 2.4)
	} else {
		varB = varB / 12.92
	}

	x = varR*41.24 + varG*35.76 + varB*18.05
	y = varR*21.26 + varG*71.52 + varB*7.22
	z = varR*1.93 + varG*11.92 + varB*95.05

	return x, y, z
}

// xyztoLAB converts CIE XYZ color space to CIE LAB color space
// https://en.wikipedia.org/wiki/Lab_color_space#CIELAB-CIEXYZ_conversions
func xyztoLAB(x, y, z float64) (l, a, b float64) {
	refX, refY, refZ := 95.047, 100.000, 108.883 // Daylight, sRGB, Adobe-RGB, Observer D65, 2°

	varX := x / refX
	varY := y / refY
	varZ := z / refZ

	if varX > 0.008856 {
		varX = math.Pow(varX, (1.0 / 3.0))
	} else {
		varX = (7.787 * varX) + (16.0 / 116.0)
	}
	if varY > 0.008856 {
		varY = math.Pow(varY, (1.0 / 3.0))
	} else {
		varY = (7.787 * varY) + (16.0 / 116.0)
	}
	if varZ > 0.008856 {
		varZ = math.Pow(varZ, (1.0 / 3.0))
	} else {
		varZ = (7.787 * varZ) + (16.0 / 116.0)
	}

	l = (116 * varY) - 16
	a = 500 * (varX - varY)
	b = 200 * (varY - varZ)

	return l, a, b
}

// colorToLAB returns LAB representation of any color (without aplha)
// https://en.wikipedia.org/wiki/Lab_color_space
func colorToLAB(color color.Color) (l, a, b float64) {
	return xyztoLAB(colorToXYZ(color))
}
