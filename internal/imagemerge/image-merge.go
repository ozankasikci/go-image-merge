package imagemerge

import (
	"errors"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Specifies how the grid pixel size should be calculated
type gridSizeMode int

const (
	// The size in pixels is fixed for all the grids
	fixedGridSize gridSizeMode = iota
	// The size in pixels is set to the nth image size
	gridSizeFromImage
)

// MergeImage is the struct that is responsible for merging the given images
type MergeImage struct {
	Grids           []*Grid
	ImageCountDX    int
	ImageCountDY    int
	BaseDir         string
	FixedGridSizeX  int
	FixedGridSizeY  int
	GridSizeMode    gridSizeMode
	GridSizeFromNth int
}

// New returns a new *MergeImage instance
func New(grids []*Grid, imageCountDX, imageCountDY int, opts ...func(*MergeImage)) *MergeImage {
	mi := &MergeImage{
		Grids:        grids,
		ImageCountDX: imageCountDX,
		ImageCountDY: imageCountDY,
	}

	for _, option := range opts {
		option(mi)
	}

	return mi
}

// OptBaseDir is an functional option to set the BaseDir field
func OptBaseDir(dir string) func(*MergeImage) {
	return func(mi *MergeImage) {
		mi.BaseDir = dir
	}
}

// OptGridSize is an functional option to set the GridSize X & Y
func OptGridSize(sizeX, sizeY int) func(*MergeImage) {
	return func(mi *MergeImage) {
		mi.GridSizeMode = fixedGridSize
		mi.FixedGridSizeX = sizeX
		mi.FixedGridSizeY = sizeY
	}
}

// OptGridSizeFromNthImageSize is an functional option to set the GridSize from the nth image
func OptGridSizeFromNthImageSize(n int) func(*MergeImage) {
	return func(mi *MergeImage) {
		mi.GridSizeMode = gridSizeFromImage
		mi.GridSizeFromNth = n
	}
}

func (m *MergeImage) readGridImage(grid *Grid) (image.Image, error) {
	if grid.Image != nil {
		return grid.ApplyFilters(grid.Image), nil
	}

	imgPath := grid.ImageFilePath

	if m.BaseDir != "" {
		imgPath = path.Join(m.BaseDir, grid.ImageFilePath)
	}

	img, err := m.ReadImageFile(imgPath)
	if err != nil {
		return nil, err
	}

	return grid.ApplyFilters(img), nil
}

func (m *MergeImage) readGridsImages() ([]image.Image, error) {
	var images []image.Image

	for _, grid := range m.Grids {
		img, err := m.readGridImage(grid)
		if err != nil {
			return nil, err
		}

		images = append(images, img)
	}

	return images, nil
}

func (m *MergeImage) ReadImageFile(path string) (image.Image, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	imgFile, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()

	var img image.Image
	splittedPath := strings.Split(path, ".")
	ext := splittedPath[len(splittedPath)-1]

	if ext == "jpg" || ext == "jpeg" {
		img, err = jpeg.Decode(imgFile)
	} else {
		img, err = png.Decode(imgFile)
	}

	if err != nil {
		return nil, err
	}

	return img, nil
}

func (m *MergeImage) mergeGrids(images []image.Image) (*image.RGBA, error) {
	var canvas *image.RGBA
	imageBoundX := 0
	imageBoundY := 0

	if m.GridSizeMode == fixedGridSize && m.FixedGridSizeX != 0 && m.FixedGridSizeY != 0 {
		imageBoundX = m.FixedGridSizeX
		imageBoundY = m.FixedGridSizeY
	} else if m.GridSizeMode == gridSizeFromImage {
		imageBoundX = images[m.GridSizeFromNth].Bounds().Dx()
		imageBoundY = images[m.GridSizeFromNth].Bounds().Dy()
	} else {
		imageBoundX = images[0].Bounds().Dx()
		imageBoundY = images[0].Bounds().Dy()
	}

	canvasBoundX := m.ImageCountDX * imageBoundX
	canvasBoundY := m.ImageCountDY * imageBoundY

	canvasMaxPoint := image.Point{canvasBoundX, canvasBoundY}
	canvasRect := image.Rectangle{image.Point{0, 0}, canvasMaxPoint}
	canvas = image.NewRGBA(canvasRect)

	// draw grids one by one
	for i, grid := range m.Grids {
		img := images[i]
		x := i % m.ImageCountDX
		y := i / m.ImageCountDX
		minPoint := image.Point{x * imageBoundX, y * imageBoundY}
		maxPoint := minPoint.Add(image.Point{imageBoundX, imageBoundY})
		nextGridRect := image.Rectangle{minPoint, maxPoint}

		if grid.BackgroundColor != nil {
			draw.Draw(canvas, nextGridRect, &image.Uniform{grid.BackgroundColor}, image.Point{}, draw.Src)
			draw.Draw(canvas, nextGridRect, img, image.Point{}, draw.Over)
		} else {
			draw.Draw(canvas, nextGridRect, img, image.Point{}, draw.Src)
		}

		if grid.Grids == nil {
			continue
		}

		// draw top layer grids
		for _, grid := range grid.Grids {
			img, err := m.readGridImage(grid)
			if err != nil {
				return nil, err
			}

			gridRect := nextGridRect.Bounds().Add(image.Point{grid.OffsetX, grid.OffsetY})
			draw.Draw(canvas, gridRect, img, image.Point{}, draw.Over)
		}
	}

	return canvas, nil
}

// Merge reads the contents of the given file paths, merges them according to given configuration
func (m *MergeImage) Merge() (*image.RGBA, error) {
	images, err := m.readGridsImages()
	if err != nil {
		return nil, err
	}

	if len(images) == 0 {
		return nil, errors.New("There is no image to merge")
	}

	return m.mergeGrids(images)
}
