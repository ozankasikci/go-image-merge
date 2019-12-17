package go_image_merge

import "image"

type Row struct {
	Cells []*Cell
}

func (t *Row) ReadImages() ([]image.Image, error) {
	var images []image.Image

	for _, cell := range t.Cells {
		img, err := readImageFile(cell.ImageFilePath)
		if err != nil {
			return nil, err
		}

		images = append(images, img)
	}

	return images, nil
}
