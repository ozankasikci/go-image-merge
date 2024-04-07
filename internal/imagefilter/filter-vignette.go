package imagefilter

import (
	"image"
	"image/color"
	"math"
)

func init() {
	Register(Vignette, new(VignetteFilter))
}

type VignetteFilter struct{}

func (vf *VignetteFilter) Apply(img image.Image) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)
	centerX, centerY := float64(bounds.Dx())/2, float64(bounds.Dy())/2
	maxDistance := math.Sqrt(centerX*centerX + centerY*centerY)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			distToCenter := math.Sqrt(math.Pow(float64(x)-centerX, 2) + math.Pow(float64(y)-centerY, 2))
			factor := distToCenter / maxDistance

			// Adjusting the curve for a stronger darkening effect towards the edges
			factor = 1 - (factor * factor * 0.95) // Increase the multiplier for a stronger effect

			// Ensure the factor does not make the image too dark, adjust the minimum threshold as needed
			if factor < 0.15 {
				factor = 0.15 // Lowering the threshold allows for more darkening at the edges
			}

			r, g, b, a := originalColor.RGBA()
			dst.Set(x, y, color.RGBA{
				R: uint8(float64(r>>8) * factor),
				G: uint8(float64(g>>8) * factor),
				B: uint8(float64(b>>8) * factor),
				A: uint8(a >> 8),
			})
		}
	}
	return dst
}
