package colorcrop

import (
	"image/color"
	"math"
	"testing"
)

func TestLinearComparators(t *testing.T) {
	comparators := []comparator{CmpEuclidean, CmpRGBComponents}

	tests := []struct {
		color1 color.Color
		color2 color.Color
		exp    float64
		got    float64
	}{
		{color1: color.RGBA{0, 0, 0, 255}, color2: color.RGBA{0, 0, 0, 255}, exp: 0.00},             // same black colors
		{color1: color.RGBA{255, 255, 255, 255}, color2: color.RGBA{255, 255, 255, 255}, exp: 0.00}, // same white colors
		{color1: color.RGBA{0, 0, 0, 255}, color2: color.RGBA{255, 255, 255, 255}, exp: 1.00},       // different (black and white) colors
		{color1: color.RGBA{255, 255, 255, 255}, color2: color.RGBA{0, 0, 0, 255}, exp: 1.00},       // different (white and black) colors
		{color1: color.RGBA{255, 255, 255, 0}, color2: color.RGBA{255, 255, 255, 255}, exp: 0.00},   // must ignore alpha channel
	}

	for _, comparator := range comparators {
		for _, test := range tests {
			test.got = comparator(test.color1, test.color2)
			if math.Abs(test.got-test.exp) > epsilon {
				t.Errorf("%v %v: expected %.8f, got %.8f",
					test.color1, test.color2, test.exp, test.got)
			}
		}
	}
}

func TestCmpCIE76(t *testing.T) {
	type test struct {
		color1 color.Color
		color2 color.Color
		exp    float64
		got    float64
	}

	tests := []test{
		{color1: color.RGBA{0, 0, 0, 255}, color2: color.RGBA{0, 0, 0, 255}, exp: 0.00000000 / 149.95514755},             // same black colors
		{color1: color.RGBA{255, 255, 255, 255}, color2: color.RGBA{255, 255, 255, 255}, exp: 0.00000000 / 149.95514755}, // same white colors
		{color1: color.RGBA{0, 0, 0, 255}, color2: color.RGBA{255, 255, 255, 255}, exp: 100.00000068 / 149.95514755},     // different (black and white) colors
		{color1: color.RGBA{255, 255, 255, 255}, color2: color.RGBA{0, 0, 0, 255}, exp: 100.00000068 / 149.95514755},     // different (white and black) colors
		{color1: color.RGBA{255, 0, 0, 255}, color2: color.RGBA{255, 255, 255, 255}, exp: 114.55897602 / 149.95514755},   // different (red and white) colors
		{color1: color.RGBA{0, 255, 0, 255}, color2: color.RGBA{255, 255, 255, 255}, exp: 120.41559907 / 149.95514755},   // different (green and white) colors
		{color1: color.RGBA{0, 0, 255, 255}, color2: color.RGBA{255, 255, 255, 255}, exp: 149.95514755 / 149.95514755},   // different (blue and white) colors
	}

	for _, test := range tests {
		test.got = CmpCIE76(test.color1, test.color2)
		if math.Abs(test.got-test.exp) > epsilon {
			t.Errorf("%v %v: expected %.8f, got %.8f",
				test.color1, test.color2, test.exp, test.got)
		}
	}
}

func TestMin(t *testing.T) {
	type test struct {
		x        uint32
		y        uint32
		expected uint32
		got      uint32
	}

	tests := []test{
		{x: 1, y: 2, expected: 1},
		{x: 3, y: 2, expected: 2},
		{x: 4, y: 4, expected: 4},
	}

	for _, test := range tests {
		test.got = min(test.x, test.y)
		if test.got != test.expected {
			t.Errorf("min(%d, %d): expected %d, got %d",
				test.x, test.y, test.expected, test.got)
		}
	}
}

func TestMax(t *testing.T) {
	type test struct {
		x        uint32
		y        uint32
		expected uint32
		got      uint32
	}

	tests := []test{
		{x: 1, y: 2, expected: 2},
		{x: 3, y: 2, expected: 3},
		{x: 4, y: 4, expected: 4},
	}

	for _, test := range tests {
		test.got = max(test.x, test.y)
		if test.got != test.expected {
			t.Errorf("max(%d, %d): expected %d, got %d",
				test.x, test.y, test.expected, test.got)
		}
	}
}
