package main

import (
	gim "github.com/ozankasikci/go-image-merge"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	paths := []string {
		"cmd/gim/kitten.jpg",
		"cmd/gim/kitten.jpg",
	}

	rgba, err := gim.New(
		paths, 2, 1,
	).Merge()
	if err != nil {
        log.Panicf(err.Error())
	}

	file, err := os.Create("cmd/gim/merged.jpg")
	if err != nil {
		log.Panicf(err.Error())
	}

	err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 80})
	if err != nil {
		log.Panicf(err.Error())
	}
}
