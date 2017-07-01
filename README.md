# colorcrop

[![Build Status](https://travis-ci.org/nxshock/colorcrop.svg?branch=master)](https://travis-ci.org/nxshock/colorcrop)
[![Coverage Status](https://coveralls.io/repos/github/nxshock/colorcrop/badge.svg)](https://coveralls.io/github/nxshock/colorcrop)
[![GoDoc](https://godoc.org/github.com/nxshock/colorcrop?status.svg)](https://godoc.org/github.com/nxshock/colorcrop)

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

Available comparators are:
- `CmpRGBComponents` - simple RGB components difference: `abs(r1-r2)+abs(g1-g2)+abs(b1-b2)` (default);
- `CmpEuclidean` - [Euclidean difference](https://en.wikipedia.org/wiki/Color_difference#Euclidean);
- `CmpCIE76` - difference of two colors defined in [CIE76 standart](https://en.wikipedia.org/wiki/Color_difference#CIE76).

## Examples

See [here](https://github.com/nxshock/colorcrop/blob/master/example_test.go).
