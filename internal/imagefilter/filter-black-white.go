package imagefilter

import (
	"image"
	"image/color"
)

func init() {
	Register(BlackAndWhite, new(BlackAndWhiteFilter))
}

type BlackAndWhiteFilter struct{}

func (bwf BlackAndWhiteFilter) Apply(img image.Image) image.Image {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(originalColor)
			grayImg.Set(x, y, grayColor)
		}
	}
	return grayImg
}
