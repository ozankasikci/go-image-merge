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
		{
			ImageFilePath:   "./cmd/gim/input/kitten.jpg",
			BackgroundColor: color.White,
		},
	}

	rows := []*go_image_merge.Row{
		{
			Cells: cells[1:],
		},
		{
			Cells: cells,
		},
	}

	merger := go_image_merge.NewMerger(rows, 400, 400)
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
