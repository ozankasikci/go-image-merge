package goimagemerge

import (
	"image"
	"image/color"
)

// Grid holds the data for each grid
type Grid struct {
	Image           image.Image
	ImageFilePath   string
	BackgroundColor color.Color
	OffsetX         int
	OffsetY         int
	Grids           []*Grid
	Filters         []FilterType
}

// ApplyFilters applies the selected filters to the image.
func (g Grid) ApplyFilters(img image.Image) image.Image {
	for _, filterType := range g.Filters {
		filter := GetFilter(filterType)
		if filter != nil {
			img = filter.Apply(img)
		}
	}

	return img
}
