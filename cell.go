package go_image_merge

import (
	"image/color"
)

type Cell struct {
	ImageFilePath   string
	BackgroundColor color.Color
	OffsetX         int
	OffsetY         int
	Cells           []*Cell
}

