package imagefilter

// FilterType defines the type of filter to apply to the image.
type FilterType int

const (
	NoFilter FilterType = iota
	// BlackAndWhite applies a black and white filter.
	BlackAndWhite
	Vignette
	Sepia
)
