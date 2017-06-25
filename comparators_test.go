package colorcrop

import (
	"image/color"
	"reflect"
	"runtime"
	"testing"
)

func TestColorComparators(t *testing.T) {
	type In struct {
		color1 color.Color
		color2 color.Color
	}

	comparators := []comparator{CmpEuclidean, CmpRGBComponents}

	tests := []struct {
		in         In
		out        float64
		commentary string
	}{
		{in: In{color.RGBA{0, 0, 0, 255}, color.RGBA{255, 255, 255, 255}},
			out:        1.00,
			commentary: "Difference between black and white colors"},
		{in: In{color.RGBA{255, 255, 255, 255}, color.RGBA{255, 255, 255, 255}},
			out:        0.00,
			commentary: "Difference between same colors"},
		{in: In{color.RGBA{255, 255, 255, 0}, color.RGBA{255, 255, 255, 255}},
			out:        0.00,
			commentary: "Difference between same colors with different transparency"},
	}

	for _, comparator := range comparators {
		for _, test := range tests {
			if comparator(test.in.color2, test.in.color1) != test.out {
				t.Errorf("%s: %s: expected %.2f, got %.2f", runtime.FuncForPC(reflect.ValueOf(comparator).Pointer()).Name(), test.commentary, test.out, comparator(test.in.color2, test.in.color1))
			}
		}
	}
}
