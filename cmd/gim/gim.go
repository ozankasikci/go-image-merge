package main

import (
	"github.com/ozankasikci/go-image-merge"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	cells := []*go_image_merge.Cell{
		{
			ImageFilePath:   "./cmd/gim/input/kitten.jpg",
			BackgroundColor: color.White,
		},
	}

	rows := []*go_image_merge.Row{
		{
			Cells: cells,
		},
	}

	merger := go_image_merge.NewMerger(rows, 100, 100)
	rgba, err := merger.Merge()
	if err != nil {
		log.Panicf(err.Error())
	}

	file, err := os.Create("cmd/gim/output/merged2.jpg")
	if err != nil {
		log.Panicf(err.Error())
	}

	err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 80})
	if err != nil {
		log.Panicf(err.Error())
	}
}

func main2() {
	grids := []*go_image_merge.Grid{
		{
			ImageFilePath:   "./cmd/gim/input/ginger.png",
			BackgroundColor: color.White,
		},
		{
			ImageFilePath:   "./cmd/gim/input/ginger.png",
			BackgroundColor: color.RGBA{R: 0x8b, G: 0xd0, B: 0xc6},
		},
	}
	rgba, err := go_image_merge.New(grids, 2, 1).Merge()
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
