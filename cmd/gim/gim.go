package main

import (
	gim "github.com/ozankasikci/go-image-merge"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	grids := []*gim.Grid{
		{
			ImageFilePath:   "./cmd/gim/input/ginger.png",
			BackgroundColor: color.White,
		},
		{
			ImageFilePath:   "./cmd/gim/input/ginger.png",
			BackgroundColor: color.RGBA{R: 0x8b, G: 0xd0, B: 0xc6},
		},
	}
	rgba, err := gim.New(grids, 2, 1).Merge()
	if err != nil {
		log.Panicf(err.Error())
	}

	file, err := os.Create("cmd/gim/output/merged.jpg")
	if err != nil {
		log.Panicf(err.Error())
	}

	err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 80})
	if err != nil {
		log.Panicf(err.Error())
	}
}
