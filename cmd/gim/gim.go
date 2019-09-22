package main

import (
	gim "github.com/ozankasikci/go-image-merge"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	paths := []string {
		"cmd/gim/heart.png",
	}

	rgba, err := gim.New(
		paths, 2, 1,
		gim.GridSizeFromNthImageSize(0),
	).Merge()
	if err != nil {
        log.Panicf(err.Error())
	}

	file, err := os.Create("test.jpg")
	if err != nil {
		log.Panicf(err.Error())
	}

	err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 80})
	if err != nil {
		log.Panicf(err.Error())
	}
}
