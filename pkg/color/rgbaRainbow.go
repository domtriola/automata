package color

import (
	"fmt"
	"image/color"
)

// CustomPalette contains an image/color.Palette and its size
type CustomPalette struct {
	color.Palette
	Size int
}

// RGBARainbow creates a palette of colors in rainbow order. It accepts a step
// parameter which determines the distance between each color shift. The smaller
// the step, the more individual colors will result in the palette.
func RGBARainbow(step int) (rainbow CustomPalette, err error) {
	if step < 7 {
		return rainbow, fmt.Errorf(
			"step must be greater than 6, got: %d. Palette cannot hold more than 256 colors",
			step,
		)
	}

	p := color.Palette{}

	// Start at Red
	rgba := color.RGBA{255, 0, 0, 255}
	p = append(p, rgba)
	size := 1

	// Red -> Yellow
	for rgba.G < 255 {
		nextValue := int(rgba.G) + step
		if nextValue > 255 {
			nextValue = 255
		}

		rgba.G = uint8(nextValue)
		p = append(p, rgba)
		size++
	}

	// Yellow -> Green
	for rgba.R > 0 {
		nextValue := int(rgba.R) - step
		if nextValue < 0 {
			nextValue = 0
		}

		rgba.R = uint8(nextValue)
		p = append(p, rgba)
		size++
	}

	// Green -> Cyan
	for rgba.B < 255 {
		nextValue := int(rgba.B) + step
		if nextValue > 255 {
			nextValue = 255
		}

		rgba.B = uint8(nextValue)
		p = append(p, rgba)
		size++
	}

	// Cyan -> Blue
	for rgba.G > 0 {
		nextValue := int(rgba.G) - step
		if nextValue < 0 {
			nextValue = 0
		}

		rgba.G = uint8(nextValue)
		p = append(p, rgba)
		size++
	}

	// Blue -> Magenta
	for rgba.R < 255 {
		nextValue := int(rgba.R) + step
		if nextValue > 255 {
			nextValue = 255
		}

		rgba.R = uint8(nextValue)
		p = append(p, rgba)
		size++
	}

	// Magenta -> Red
	for rgba.B > 0 {
		nextValue := int(rgba.B) - step
		if nextValue < 0 {
			nextValue = 0
		}

		rgba.B = uint8(nextValue)
		p = append(p, rgba)
		size++
	}

	rainbow.Palette = p
	rainbow.Size = size

	return rainbow, nil
}
