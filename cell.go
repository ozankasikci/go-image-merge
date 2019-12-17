package go_image_merge

import (
	"image/color"
)

type Cell struct {
	Merge          *Merger
	ImageFilePath   string
	BackgroundColor color.Color
	OffsetX         int
	OffsetY         int
	Cells           []*Cell
}


