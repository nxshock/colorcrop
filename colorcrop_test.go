package colorcrop

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

// epsilon is a maximum permissible error
const epsilon = 0.00000001

func TestCropRectanle(t *testing.T) {
	type test struct {
		filename string
		expected image.Rectangle
		got      image.Rectangle
	}

	comparators := []comparator{CmpRGBComponents, CmpEuclidean, CmpCIE76}
	thresold := 0.5

	tests := []test{
		{filename: "01.png", expected: image.Rectangle{image.Point{0, 0}, image.Point{1, 1}}},
		{filename: "02.png", expected: image.Rectangle{image.Point{1, 0}, image.Point{2, 1}}},
		{filename: "03.png", expected: image.Rectangle{image.Point{2, 0}, image.Point{3, 1}}},
		{filename: "04.png", expected: image.Rectangle{image.Point{3, 0}, image.Point{4, 1}}},
		{filename: "05.png", expected: image.Rectangle{image.Point{0, 1}, image.Point{1, 2}}},
		{filename: "06.png", expected: image.Rectangle{image.Point{1, 1}, image.Point{2, 2}}},
		{filename: "07.png", expected: image.Rectangle{image.Point{2, 1}, image.Point{3, 2}}},
		{filename: "08.png", expected: image.Rectangle{image.Point{3, 1}, image.Point{4, 2}}},
		{filename: "09.png", expected: image.Rectangle{image.Point{0, 2}, image.Point{1, 3}}},
		{filename: "10.png", expected: image.Rectangle{image.Point{1, 2}, image.Point{2, 3}}},
		{filename: "11.png", expected: image.Rectangle{image.Point{2, 2}, image.Point{3, 3}}},
		{filename: "12.png", expected: image.Rectangle{image.Point{3, 2}, image.Point{4, 3}}},
		{filename: "13.png", expected: image.Rectangle{image.Point{0, 3}, image.Point{1, 4}}},
		{filename: "14.png", expected: image.Rectangle{image.Point{1, 3}, image.Point{2, 4}}},
		{filename: "15.png", expected: image.Rectangle{image.Point{2, 3}, image.Point{3, 4}}},
		{filename: "16.png", expected: image.Rectangle{image.Point{3, 3}, image.Point{4, 4}}},
	}

	for _, comparator := range comparators {
		for _, test := range tests {
			file, err := os.Open("testimages/" + test.filename)
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()
			image, err := png.Decode(file)
			if err != nil {
				t.Fatal(err)
			}
			test.got = cropRectanle(image, color.RGBA{255, 255, 255, 255}, thresold, comparator)
			if !reflect.DeepEqual(test.expected, test.got) {
				t.Errorf("expected %v, got %v for comparator: %s, file: %s",
					test.expected, test.got, getFuncName(comparator), test.filename)
			}
		}
	}
}

func TestCropWithComparator(t *testing.T) {
	type test struct {
		filename                      string
		expectedWidth, expectedHeight int
		gotWidth, gotHeight           int
	}

	comparators := []comparator{CmpRGBComponents, CmpEuclidean, CmpCIE76}
	thresold := 0.5
	color := color.RGBA{255, 255, 255, 255}

	tests := []test{
		{filename: "01.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "02.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "03.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "04.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "05.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "06.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "07.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "08.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "09.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "10.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "11.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "12.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "13.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "14.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "15.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "16.png", expectedWidth: 1, expectedHeight: 1},
	}

	for _, comparator := range comparators {
		for _, test := range tests {
			file, err := os.Open("testimages/" + test.filename)
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()
			image, err := png.Decode(file)
			if err != nil {
				t.Fatal(err)
			}
			newImage := CropWithComparator(image, color, thresold, comparator)
			test.gotWidth, test.gotHeight = getImageSize(newImage)
			if test.gotWidth != test.expectedWidth || test.gotHeight != test.expectedWidth {
				t.Errorf("expected {%d %d}, got {%d %d} for comparator: %s, file: %s",
					test.expectedWidth, test.expectedWidth, test.gotWidth, test.gotHeight, getFuncName(comparator), test.filename)
			}
		}
	}
}

func TestCrop(t *testing.T) {
	type test struct {
		filename                      string
		expectedWidth, expectedHeight int
		gotWidth, gotHeight           int
	}

	thresold := 0.5
	color := color.RGBA{255, 255, 255, 255}

	tests := []test{
		{filename: "01.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "02.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "03.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "04.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "05.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "06.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "07.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "08.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "09.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "10.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "11.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "12.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "13.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "14.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "15.png", expectedWidth: 1, expectedHeight: 1},
		{filename: "16.png", expectedWidth: 1, expectedHeight: 1},
	}

	for _, test := range tests {
		file, err := os.Open("testimages/" + test.filename)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()
		image, err := png.Decode(file)
		if err != nil {
			t.Fatal(err)
		}
		newImage := Crop(image, color, thresold)
		test.gotWidth, test.gotHeight = getImageSize(newImage)
		if test.gotWidth != test.expectedWidth || test.gotHeight != test.expectedWidth {
			t.Errorf("expected {%d %d}, got {%d %d} for file: %s",
				test.expectedWidth, test.expectedWidth, test.gotWidth, test.gotHeight, test.filename)
		}
	}
}

func getFuncName(i interface{}) string {
	s := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	p := strings.LastIndex(s, ".")
	if p > 0 {
		return string([]rune(s)[p+1:])
	}
	return s
}

func getImageSize(image image.Image) (width, height int) {
	return image.Bounds().Dx(), image.Bounds().Dy()
}
