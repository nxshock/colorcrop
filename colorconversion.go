package colorcrop

import (
	"image/color"
	"math"
)

func rgbtoXYZ(r, g, b uint32) (x, y, z float64) {
	varR := float64(r) / 255
	varG := float64(g) / 255
	varB := float64(b) / 255

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

	varR = varR * 100
	varG = varG * 100
	varB = varB * 100

	x = varR*0.4124 + varG*0.3576 + varB*0.1805
	y = varR*0.2126 + varG*0.7152 + varB*0.0722
	z = varR*0.0193 + varG*0.1192 + varB*0.9505
	return x, y, z
}

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

func colorToLAB(color color.Color) (l, a, b float64) {
	cr, cg, cb, _ := color.RGBA()
	return xyztoLAB(rgbtoXYZ(cr, cg, cb))
}
