package palette

import (
	"image/color"
	"log"
)

// Grey creates a greyscale palette with increments of 1 from 0 to 255.
func Grey() color.Palette {
	p := color.Palette{}
	for i := 0; i < 255; i++ {
		p = append(p, color.RGBA{
			uint8(i),
			uint8(i),
			uint8(i),
			255,
		})
	}

	log.Printf("Grey() length: %d", len(p))

	return p
}
