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
	canvasBoundX := t.SizeX
	canvasBoundY := t.SizeY

	canvasMaxPoint := image.Point{canvasBoundX, canvasBoundY}
	canvasRect := image.Rectangle{image.Point{0, 0}, canvasMaxPoint}
	canvas = image.NewRGBA(canvasRect)
	err := t.drawCells(canvas)
	if err != nil {
	    return nil, err
	}

	return canvas, nil
}

func (t *Merger) drawCells(canvas *image.RGBA) error {
	for i, row := range t.Rows {
		for j, cell := range row.Cells {
			img, err := t.readCellImage(cell)
			if err != nil {
				return err
			}

			minPoint := image.Point{j * canvas.Bounds().Dx() / len(row.Cells), i * canvas.Bounds().Dy() / len(t.Rows)}
			maxPoint := minPoint.Add(image.Point{canvas.Bounds().Dx() / len(row.Cells), canvas.Bounds().Dy() / len(t.Rows)})
			cellRect := image.Rectangle{minPoint, maxPoint}

			if cell.BackgroundColor != nil {
				draw.Draw(canvas, cellRect, &image.Uniform{cell.BackgroundColor}, image.Point{}, draw.Src)
				draw.Draw(canvas, cellRect, img, image.Point{}, draw.Over)
			} else {
				draw.Draw(canvas, cellRect, img, image.Point{}, draw.Src)
			}

		}
	}

	return nil
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
