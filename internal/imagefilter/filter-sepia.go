package imagefilter

import (
	"image"
	"image/color"
)

func init() {
	Register(Sepia, new(SepiaFilter))
}

type SepiaFilter struct{}

func (sf *SepiaFilter) Apply(img image.Image) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			r, g, b := originalColor.R, originalColor.G, originalColor.B

			// Calculate sepia tones
			newR := (0.393 * float64(r)) + (0.769 * float64(g)) + (0.189 * float64(b))
			newG := (0.349 * float64(r)) + (0.686 * float64(g)) + (0.168 * float64(b))
			newB := (0.272 * float64(r)) + (0.534 * float64(g)) + (0.131 * float64(b))

			// Clamp values
			newR = min(newR, 255)
			newG = min(newG, 255)
			newB = min(newB, 255)

			// Mix with original color to reduce intensity
			mixFactor := 0.7 // Adjust this value to control the sepia intensity
			mixedR := mixFactor*newR + (1-mixFactor)*float64(r)
			mixedG := mixFactor*newG + (1-mixFactor)*float64(g)
			mixedB := mixFactor*newB + (1-mixFactor)*float64(b)

			dst.Set(x, y, color.RGBA{
				R: uint8(mixedR),
				G: uint8(mixedG),
				B: uint8(mixedB),
				A: originalColor.A,
			})
		}
	}
	return dst
}

// min returns the smaller of x or y.
func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}
