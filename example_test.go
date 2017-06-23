package colorcrop_test

import (
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/nxshock/colorcrop"
)

func Example() {
	log.SetFlags(0)

	// Read source image
	sourceFile, _ := os.Open("img.png")
	defer sourceFile.Close()

	sourceImage, _ := png.Decode(sourceFile)

	// Crop image white border with 50% thresold
	croppedImage := colorcrop.Crop(sourceImage, color.RGBA{255, 255, 255, 255}, 0.5)

	// Save cropped image
	croppedFile, _ := os.Create("cropped.png")
	defer croppedFile.Close()

	png.Encode(croppedFile, croppedImage)
}
