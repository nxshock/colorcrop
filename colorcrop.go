// Package colorcrop is a Go library for cropping images by removing borders
// with specified color.
package colorcrop

import (
	"image"
	"image/color"
)

// Crop returns cropped image with default comparator.
func Crop(img image.Image, color color.Color, thresold float64) image.Image {
	return CropWithComparator(img, color, thresold, CmpRGBComponents)
}

// CropWithComparator returns cropped image with specified comparator.
func CropWithComparator(img image.Image, color color.Color, thresold float64, comparator comparator) image.Image {
	return img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(cropRectanle(img, color, thresold, comparator))
}

// cropRectanle returns rectangle of image without borders.
func cropRectanle(img image.Image, color color.Color, thresold float64, comparator comparator) image.Rectangle {
	rectangle := img.Bounds()

TopLoop:
	for y := rectangle.Min.Y; y < rectangle.Max.Y; y++ {
		rectangle.Min.Y = y
		for x := rectangle.Min.X; x < rectangle.Max.X; x++ {
			if comparator(img.At(x, y), color) > thresold {
				break TopLoop
			}
		}
	}

BottomLoop:
	for y := rectangle.Max.Y - 1; y >= rectangle.Min.Y; y-- {
		rectangle.Max.Y = y + 1
		for x := rectangle.Min.X; x < rectangle.Max.X; x++ {
			if comparator(img.At(x, y), color) > thresold {
				break BottomLoop
			}
		}
	}

LeftLoop:
	for x := rectangle.Min.X; x < rectangle.Max.X; x++ {
		rectangle.Min.X = x
		for y := rectangle.Min.Y; y < rectangle.Max.Y; y++ {
			if comparator(img.At(x, y), color) > thresold {
				break LeftLoop
			}
		}

	}

RightLoop:
	for x := rectangle.Max.X - 1; x >= rectangle.Min.X; x-- {
		rectangle.Max.X = x + 1
		for y := rectangle.Min.Y; y < rectangle.Max.Y; y++ {
			if comparator(img.At(x, y), color) > thresold {
				break RightLoop
			}
		}
	}

	return rectangle
}
