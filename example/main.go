package main

import (
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/nxshock/colorcrop"
)

func main() {
	log.SetFlags(0)

	// Read source image
	sourceFile, err := os.Open("img.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer sourceFile.Close()

	sourceImage, err := png.Decode(sourceFile)
	if err != nil {
		log.Fatalln(err)
	}

	// Crop image white border with 50% thresold
	croppedImage := colorcrop.Crop(sourceImage, color.RGBA{255, 255, 255, 255}, 0.5)

	// Save cropped image
	croppedFile, err := os.Create("cropped.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer croppedFile.Close()

	err = png.Encode(croppedFile, croppedImage)
	if err != nil {
		log.Fatalln(err)
	}
}
