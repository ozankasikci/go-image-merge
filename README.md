## gim - Grid Based Image Merge Library

`gim` is a image merging library that can accept image paths as input, read the image contents, add background color, draw layers on top of each other, merge them into a grid with the desired size.

[![Go Report Card](https://goreportcard.com/badge/github.com/ozankasikci/go-image-merge)](https://goreportcard.com/report/github.com/ozankasikci/go-image-merge)
[![GoDoc](https://godoc.org/github.com/ozankasikci/go-image-merge?status.svg)](https://godoc.org/github.com/ozankasikci/go-image-merge)

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Getting Started](#getting-started)
- [Examples](#examples)
  * [Grid Unit Count - Vertical & Horizontal](#grid-unit-count---vertical--horizontal)
  * [Grid Background Color](#grid-background-color)
  * [Grid Layers - Draw Grids on top of Grids](#grid-layers---draw-grids-on-top-of-grids)
- [Functional Options](#functional-options)
- [TODO](#todo)

## Overview

`gim` provides an easy and extensible way to merge images into a flexible grid system.

The main purpose of the library is to help creating image collages programmatically.

## Installation

`go get -u github.com/ozankasikci/go-image-merge`

## Getting Started

Import the library and give the image paths and grid size as the minimum required arguments.

Basic usage:

```go
import gim "github.com/ozankasikci/go-image-merge"

// accepts *Grid instances, grid unit count x, grid unit count y
// returns an *image.RGBA object
grids := []*gim.Grid{
	{ImageFilePath: "file.jpg"},
	{ImageFilePath: "file.png"},
}
rgba, err := gim.New(grids, 2, 1).Merge()

// save the output to jpg or png
file, err := os.Create("file/path.jpg|png")
err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 80})
err = png.Encode(file, rgba)
```

See [Examples](#examples) for available options and advanced usage.

## Examples

### Grid Unit Count - Vertical & Horizontal
```go
grids := []*gim.Grid{
    {ImageFilePath: "./cmd/gim/input/kitten.jpg"},
    {ImageFilePath: "./cmd/gim/input/kitten.jpg"},
    {ImageFilePath: "./cmd/gim/input/kitten.jpg"},
    {ImageFilePath: "./cmd/gim/input/kitten.jpg"},
}
rgba, err := gim.New(grids, 2, 2).Merge()
```

#### Output
![](https://raw.githubusercontent.com/ozankasikci/ozankasikci.github.io/master/gim/grid-size-2-2.jpg)

### Grid Background Color
```go
grids := []*gim.Grid{
    {
        ImageFilePath: "./cmd/gim/input/ginger.png",
        BackgroundColor: color.White,
    },
    {
        ImageFilePath: "./cmd/gim/input/ginger.png",
        BackgroundColor: color.RGBA{R: 0x8b, G: 0xd0, B: 0xc6},
    },
}
rgba, err := gim.New(grids, 2, 1).Merge()
```

#### Output
![](https://raw.githubusercontent.com/ozankasikci/ozankasikci.github.io/master/gim/grid-bg-color.jpg)

### Grid Layers - Draw Grids on top of Grids
```go
grids := []*gim.Grid{
    {
        ImageFilePath: "./cmd/gim/input/ginger.png",
        BackgroundColor: color.White,
        // these grids will be drawn on top of the first grid
        Grids: []*gim.Grid{
            {
            	ImageFilePath: "./cmd/gim/input/tick.png",
            	OffsetX: 50, OffsetY: 20,
            },
        },
    },
    {
        ImageFilePath: "./cmd/gim/input/ginger.png",
        BackgroundColor: color.RGBA{R: 0x8b, G: 0xd0, B: 0xc6},
        // these grids will be drawn on top of the second grid
        Grids: []*gim.Grid{
            {
            	ImageFilePath: "./cmd/gim/input/tick.png",
            	OffsetX: 200, OffsetY: 170,
            },
            {
            	ImageFilePath: "./cmd/gim/input/tick.png",
            	OffsetX: 200, OffsetY: 20,
            },
        },
    },
}
rgba, err := gim.New(grids, 2, 1).Merge()
```

#### Output
![](https://raw.githubusercontent.com/ozankasikci/ozankasikci.github.io/master/gim/grid-layers.jpg)

## Functional Options

### OptBaseDir
```go
// you can omit the full path if you set a base dir
grids := []*gim.Grid{
    {ImageFilePath: "kitten.jpg"},
    {ImageFilePath: "kitten.jpg"},
}
rgba, err := gim.New(grids, 1, 2,
	gim.OptBaseDir("./cmd/gim/input"),
).Merge()
```

### OptGridSize
```go
// you can resize the grids in pixels
grids := []*gim.Grid{
    {ImageFilePath: "kitten.jpg"},
    {ImageFilePath: "kitten.jpg"},
}
rgba, err := gim.New( grids, 2, 1,
	gim.OptBaseDir("./cmd/gim"),
	gim.OptGridSize(200,150),
).Merge()
```
#### Output
![](https://raw.githubusercontent.com/ozankasikci/ozankasikci.github.io/master/gim/grid-resize-pixels-200-150.jpg)

## TODO
- [x] Add colored background support
- [ ] Add resize support (stretch, fit etc.)

