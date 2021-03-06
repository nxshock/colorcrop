colorcrop
=========

|bs| |cs| |rc| |gd|

.. |bs| image:: https://travis-ci.org/nxshock/colorcrop.svg?branch=master
   :alt: Build Status
   :target: https://travis-ci.org/nxshock/colorcrop
.. |cs| image:: https://coveralls.io/repos/github/nxshock/colorcrop/badge.svg
   :alt: Coverage Status
   :target: https://coveralls.io/github/nxshock/colorcrop
.. |rc| image:: https://goreportcard.com/badge/github.com/nxshock/colorcrop
   :alt: Go Report Card Status
   :target: https://goreportcard.com/report/github.com/nxshock/colorcrop
.. |gd| image:: https://godoc.org/github.com/nxshock/colorcrop?status.svg
   :alt: GoDoc
   :target: https://godoc.org/github.com/nxshock/colorcrop

A pure Go library for cropping images by removing borders with specified color.

Installation
------------

``go get -u github.com/nxshock/colorcrop``

Usage
-----

Import package with:

.. code:: go

    import "github.com/nxshock/colorcrop"

Crop white borders with 50% of thresold:

.. code:: go

    croppedImage := colorcrop.Crop(
        sourceImage,                    // for source image
        color.RGBA{255, 255, 255, 255}, // crop white border
        0.5)                            // with 50% thresold

You may use custom comparator of colors:

.. code:: go

    croppedImage := colorcrop.CropWithComparator(
        sourceImage,                    // for source image
        color.RGBA{255, 255, 255, 255}, // crop white border
        0.5,                            // with 50% thresold
        colorcrop.CmpCIE76)             // using CIE76 standart for defining color difference

List of available comparators:

================  =============================================================================================================
Comparator        Description
================  =============================================================================================================
CmpRGBComponents  simple RGB components difference: ``abs(r1-r2)+abs(g1-g2)+abs(b1-b2)`` (default);
CmpEuclidean      `Euclidean difference <https://en.wikipedia.org/wiki/Color_difference#Euclidean>`_;
CmpCIE76          difference of two colors defined in `CIE76 standard <https://en.wikipedia.org/wiki/Color_difference#CIE76>`_.
================  =============================================================================================================

Examples
--------

See `here <https://github.com/nxshock/colorcrop/blob/master/example_test.go>`_.
