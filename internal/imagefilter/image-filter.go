package imagefilter

import "image"

// Filter is the interface that wraps the Apply method.
type Filter interface {
	Apply(image.Image) image.Image
}

var Filters = make(map[FilterType]Filter)

func Register(filterType FilterType, filter Filter) {
	Filters[filterType] = filter
}

func Get(filterType FilterType) Filter {
	return Filters[filterType]
}
