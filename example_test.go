package colorcrop_test

import (
	"image/color"
	"image/png"
	"os"

	"github.com/nxshock/colorcrop"
)

// Simple remove of white borders.
func ExampleCrop() {
	// Read source image
	sourceFile, _ := os.Open("img.png")
	defer sourceFile.Close()

	sourceImage, _ := png.Decode(sourceFile)

	// Crop white border with 50% thresold
	croppedImage := colorcrop.Crop(
		sourceImage,                    // for source image
		color.RGBA{255, 255, 255, 255}, // crop white border
		0.5)                            // with 50% thresold

	// Save cropped image
	croppedFile, _ := os.Create("cropped.png")
	defer croppedFile.Close()

	png.Encode(croppedFile, croppedImage)
}

// Remove white borders with custom color comparator.
func ExampleCropWithComparator() {
	// Read source image
	sourceFile, _ := os.Open("img.png")
	defer sourceFile.Close()

	sourceImage, _ := png.Decode(sourceFile)

	// Crop white border with 50% thresold
	croppedImage := colorcrop.CropWithComparator(
		sourceImage,                    // for source image
		color.RGBA{255, 255, 255, 255}, // crop white border
		0.5,                            // with 50% thresold
		colorcrop.CmpCIE76)             // using CIE76 standart for defining color difference

	// Save cropped image
	croppedFile, _ := os.Create("cropped.png")
	defer croppedFile.Close()

	png.Encode(croppedFile, croppedImage)
}
