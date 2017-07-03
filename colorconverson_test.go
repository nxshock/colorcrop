package colorcrop

import (
	"image/color"
	"math"
	"testing"
)

func TestColorToLAB(t *testing.T) {
	tests := []struct {
		color            color.Color
		expL, expA, expB float64
		gotL, gotA, gotB float64
	}{
		{color: color.RGBA{0, 0, 0, 255}, expL: 0.0, expA: 0.0, expB: 0.0},
		{color: color.RGBA{0, 0, 255, 255}, expL: 32.30258667, expA: 79.19666179, expB: -107.86368104},
		{color: color.RGBA{0, 255, 0, 255}, expL: 87.73703347, expA: -86.18463650, expB: 83.18116475},
		{color: color.RGBA{255, 0, 0, 255}, expL: 53.23288179, expA: 80.10930953, expB: 67.22006831},
		{color: color.RGBA{255, 255, 255, 255}, expL: 100.00000000, expA: 0.00526050, expB: -0.01040818},
	}

	for _, test := range tests {
		test.gotL, test.gotA, test.gotB = colorToLAB(test.color)
		if math.Abs(test.gotL-test.expL) > epsilon || math.Abs(test.gotA-test.expA) > epsilon ||
			math.Abs(test.gotB-test.expB) > epsilon {
			t.Errorf("%v: expected {%.8f, %.8f, %.8f}, got {%.8f, %.8f, %.8f}",
				test.color, test.expL, test.expA, test.expB, test.gotL, test.gotA, test.gotB)
		}
	}
}
