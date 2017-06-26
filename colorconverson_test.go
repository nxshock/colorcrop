package colorcrop

import (
	"image/color"
	"math"
	"testing"
)

func TestColorToLAB(t *testing.T) {
	tests := []struct {
		color                           color.Color
		expectedL, expectedA, expectedB float64
		gotL, gotA, gotB                float64
	}{
		{color: color.RGBA{0, 0, 0, 255}, expectedL: 0.0, expectedA: 0.0, expectedB: 0.0},
		{color: color.RGBA{0, 0, 255, 255}, expectedL: 32.30258667, expectedA: 79.19666179, expectedB: -107.86368104},
		{color: color.RGBA{0, 255, 0, 255}, expectedL: 87.73703347, expectedA: -86.18463650, expectedB: 83.18116475},
		{color: color.RGBA{255, 0, 0, 255}, expectedL: 53.23288179, expectedA: 80.10930953, expectedB: 67.22006831},
		{color: color.RGBA{255, 255, 255, 255}, expectedL: 100.00000000, expectedA: 0.00526050, expectedB: -0.01040818},
	}

	for _, test := range tests {
		test.gotL, test.gotA, test.gotB = colorToLAB(test.color)
		if math.Abs(test.gotL-test.expectedL) > epsilon || math.Abs(test.gotA-test.expectedA) > epsilon || math.Abs(test.gotB-test.expectedB) > epsilon {
			t.Errorf("%v: expected {%.8f, %.8f, %.8f}, got {%.8f, %.8f, %.8f}", test.color, test.expectedL, test.expectedA, test.expectedB, test.gotL, test.gotA, test.gotB)
		}
	}
}
