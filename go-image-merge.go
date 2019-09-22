package goimagemerge

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
	ImageFilePaths  []string
	ImageCountDX    int
	ImageCountDY    int
	BaseDir         string
	FixedGridSizeX  int
	FixedGridSizeY  int
	GridSizeMode    gridSizeMode
	GridSizeFromNth int
}

// New returns a new *MergeImage instance
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

func (m *MergeImage) readImageFiles(paths []string) ([]image.Image, error) {
	var images []image.Image

	for _, imgPath := range paths {
		if m.BaseDir != "" {
			imgPath = path.Join(m.BaseDir, imgPath)
		}

		img, err := m.readImageFile(imgPath)
		if err != nil {
			return nil, err
		}

		images = append(images, img)
	}

	return images, nil
}

func (m *MergeImage) readImageFile(path string) (image.Image, error) {
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

func (m *MergeImage) mergeImages(images []image.Image, canvasXUnit, canvasYUnit int) (*image.RGBA, error) {
	var rgba *image.RGBA
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

	return rgba, nil
}

// Merge reads the contents of the given file paths, merges them according to given configuration
func (m *MergeImage) Merge() (*image.RGBA, error) {
	images, err := m.readImageFiles(m.ImageFilePaths)
	if err != nil {
		return nil, err
	}

	return m.mergeImages(images, m.ImageCountDX, m.ImageCountDY)
}
