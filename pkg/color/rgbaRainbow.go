package color

import (
	"fmt"
	"image/color"
)

// RGBARainbow creates a palette of colors in rainbow order. It accepts a step
// parameter which determines the distance between each color shift. The smaller
// the step, the more individual colors will result in the palette. The smallest
// possible step is 7, which results in 223 evenly spaced colors. With larger
// steps, there will be less colors, and the colors will be less evenly spaced.
// TODODOM: make that not so by passing overflow to next color selection.
func RGBARainbow(step int) (color.Palette, error) {
	p := color.Palette{}

	if step < 7 {
		return p, fmt.Errorf(
			"step must be greater than 6, got: %d. Palette cannot hold more than 256 colors",
			step,
		)
	}

	if step > 255 {
		return p, fmt.Errorf(
			"step must be less than 256, got: %d. 255 is the greatest color increment possible",
			step,
		)
	}

	// Start at Red
	rgba := color.RGBA{255, 0, 0, 255}
	p = append(p, rgba)

	// Red -> Yellow
	for rgba.G < 255 {
		nextValue := int(rgba.G) + step

		// TODODOM: don't discard overflow. Use it to make distribution more even.
		if nextValue > 255 {
			nextValue = 255
		}

		rgba.G = uint8(nextValue)
		p = append(p, rgba)
	}

	// Yellow -> Green
	for rgba.R > 0 {
		nextValue := int(rgba.R) - step
		if nextValue < 0 {
			nextValue = 0
		}

		rgba.R = uint8(nextValue)
		p = append(p, rgba)
	}

	// Green -> Cyan
	for rgba.B < 255 {
		nextValue := int(rgba.B) + step
		if nextValue > 255 {
			nextValue = 255
		}

		rgba.B = uint8(nextValue)
		p = append(p, rgba)
	}

	// Cyan -> Blue
	for rgba.G > 0 {
		nextValue := int(rgba.G) - step
		if nextValue < 0 {
			nextValue = 0
		}

		rgba.G = uint8(nextValue)
		p = append(p, rgba)
	}

	// Blue -> Magenta
	for rgba.R < 255 {
		nextValue := int(rgba.R) + step
		if nextValue > 255 {
			nextValue = 255
		}

		rgba.R = uint8(nextValue)
		p = append(p, rgba)
	}

	// Magenta -> Red
	for rgba.B > 0 {
		nextValue := int(rgba.B) - step
		if nextValue < 0 {
			nextValue = 0
		}

		rgba.B = uint8(nextValue)
		p = append(p, rgba)
	}

	return p, nil
}
