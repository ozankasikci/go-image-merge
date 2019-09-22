package go_image_merge

import (
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
)

type MergeImage struct {
	ImageFilePaths []string
	ImageCountDX   int
	ImageCountDY   int
	BaseDir        string
	FixedSizeX     int
	FixedSizeY     int
}

func New(paths []string, imageCountDX, imageCountDY int, opts ...func(*MergeImage)) *MergeImage {
	mi := &MergeImage{
		ImageFilePaths: paths,
		ImageCountDX:   imageCountDX,
		ImageCountDY:   imageCountDY,
	}

	for _, option := range opts {
		option(mi)
	}

	return mi
}

func BaseDir(dir string) func(*MergeImage) {
	return func(mi *MergeImage) {
		mi.BaseDir = dir
	}
}

func GridSize(sizeX, sizeY int) func(*MergeImage) {
	return func(mi *MergeImage) {
		mi.FixedSizeX = sizeX
		mi.FixedSizeY = sizeY
	}
}

func (m *MergeImage) readImageFiles(paths []string) ([]image.Image, error) {
	var images []image.Image

	for _, path := range paths {
		if m.BaseDir != "" {
			path = m.BaseDir + path
		}

		img, err := readImageFile(path)
		if err != nil {
			return nil, err
		}

		images = append(images, img)
	}

	return images, nil
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

	img, err := png.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (m *MergeImage) mergeImages(images []image.Image, canvasXUnit, canvasYUnit int) *image.RGBA {
	var rgba *image.RGBA
	imageBoundX := 0
	imageBoundY := 0

	if m.FixedSizeX != 0 && m.FixedSizeY != 0{
		imageBoundX = m.FixedSizeX
		imageBoundY = m.FixedSizeY
	}
	
	canvasX := canvasXUnit * imageBoundX
	canvasY := canvasYUnit * imageBoundY

	canvasDimension := image.Point{canvasX, canvasY}
	canvasRec := image.Rectangle{image.Point{0, 0}, canvasDimension}
	rgba = image.NewRGBA(canvasRec)

	for i, img := range images {
		x := i % canvasXUnit
		y := i / canvasXUnit
		minPoint := image.Point{x * imageBoundX, y * imageBoundY}
		maxPoint := minPoint.Add(image.Point{imageBoundX, imageBoundY})
		rec := image.Rectangle{minPoint, maxPoint}
		draw.Draw(rgba, rec, img, image.Point{0, 0}, draw.Src)
	}

	return rgba
}

func (m *MergeImage) Merge() (*image.RGBA, error) {
	images, err := m.readImageFiles(m.ImageFilePaths)
	if err != nil {
		return nil, err
	}

	return m.mergeImages(images, m.ImageCountDX, m.ImageCountDY), nil
}

