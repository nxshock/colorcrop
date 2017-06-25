# colorcrop

A pure Go library for cropping images by removing borders with specified color.

## Installation

`go get -u github.com/nxshock/colorcrop`

## Usage

Import package with

```go
import "github.com/nxshock/colorcrop"
```

Crop **white** borders with **50%** of thresold:

```go
croppedImage := colorcrop.Crop(sourceImage, color.RGBA{255, 255, 255, 255}, 0.5)
```

You may use custom comparator of colors:

```go
croppedImage := colorcrop.CropWithComparator(sourceImage, color.RGBA{255, 255, 255, 255}, 0.5, colorcrop.CmpCIE76)
```

## Examples

See in "examples".
