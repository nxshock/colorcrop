# colorcrop

Go library for cropping images by removing borders with specified color.

## Installation

`go get -u github.com/nxshock/colorcrop`

## Usage

Crop white borders with 50% of thresold:

`croppedImage := colorcrop.Crop(sourceImage, color.RGBA{255, 255, 255, 255}, 0.5)`

## Examples

See in "examples".
