# colorcrop

A pure Go library for cropping images by removing borders with specified color.

## Installation

`go get -u github.com/nxshock/colorcrop`

## Usage

Import package with

```go
import "github.com/nxshock/colorcrop"
```

Crop white borders with 50% of thresold:

```go
croppedImage := colorcrop.Crop(
    sourceImage,                    // for source image
    color.RGBA{255, 255, 255, 255}, // crop white border
    0.5)                            // with 50% thresold
```

You may use custom comparator of colors:

```go
croppedImage := colorcrop.CropWithComparator(
    sourceImage,                    // for source image
    color.RGBA{255, 255, 255, 255}, // crop white border
    0.5,                            // with 50% thresold
    colorcrop.CmpCIE76)             // using CIE76 standart for defining color difference
```

## Examples

See in "examples".
