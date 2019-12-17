package go_image_merge

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Merger struct {
	Rows []*Row
	SizeX int
	SizeY int
	BaseDir string
}

// New returns a new *MergeImage instance
func NewMerger(rows []*Row, sizeX, sizeY int, opts ...func(*Merger)) *Merger {
	m := &Merger{
		Rows:rows,
		SizeX:sizeX,
		SizeY:sizeY,
	}

	for _, option := range opts {
		option(m)
	}

	return m
}

func (t *Merger) readCellImage(cell *Cell) (image.Image, error) {
	imgPath := cell.ImageFilePath

	if t.BaseDir != "" {
		imgPath = path.Join(t.BaseDir, cell.ImageFilePath)
	}

	return readImageFile(imgPath)
}

func (t *Merger) Merge() (*image.RGBA, error) {
	var canvas *image.RGBA
	canvasBoundX := 400
	canvasBoundY := 400

	canvasMaxPoint := image.Point{canvasBoundX, canvasBoundY}
	canvasRect := image.Rectangle{image.Point{0, 0}, canvasMaxPoint}
	canvas = image.NewRGBA(canvasRect)

	for _, row := range t.Rows {
		images, _ := row.ReadImages()
		minPoint := image.Point{}
		maxPoint := minPoint.Add(image.Point{200, 200})
		gridRect := image.Rectangle{minPoint, maxPoint}
		draw.Draw(canvas, gridRect, images[0], image.Point{}, draw.Src)
		//for i, cell := range row.Cells {
		//	img := images[i]
		//	x := i % m.ImageCountDX
		//	y := i / m.ImageCountDX
		//	minPoint := image.Point{x * imageBoundX, y * imageBoundY}
		//	maxPoint := minPoint.Add(image.Point{imageBoundX, imageBoundY})
		//	nextGridRect := image.Rectangle{minPoint, maxPoint}
		//
		//	if cell.BackgroundColor != nil {
		//		draw.Draw(canvas, nextGridRect, &image.Uniform{cell.BackgroundColor}, image.Point{}, draw.Src)
		//		draw.Draw(canvas, nextGridRect, img, image.Point{}, draw.Over)
		//	} else {
		//		draw.Draw(canvas, nextGridRect, img, image.Point{}, draw.Src)
		//	}
		//
		//	if cell.Grids == nil {
		//		continue
		//	}
		//
		//	// draw top layer grids
		//	for _, grid := range cell.Grids {
		//		img, err := m.readGridImage(grid)
		//		if err != nil {
		//			return nil, err
		//		}
		//
		//		gridRect := nextGridRect.Bounds().Add(image.Point{grid.OffsetX, grid.OffsetY})
		//		draw.Draw(canvas, gridRect, img, image.Point{}, draw.Over)
		//	}
		//}
	}

	// draw grids one by one

	return canvas, nil
}

// OptBaseDir is an functional option to set the BaseDir field
func OptBaseDirMerge(dir string) func(*Merger) {
	return func(mi *Merger) {
		mi.BaseDir = dir
	}
}

func readImageFile(path string) (image.Image, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	imgFile, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}

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
