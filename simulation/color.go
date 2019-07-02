package simulation

import (
	"fmt"
	"image/color"
)

// RGBARainbow creates an array of colors in rainbow order
func RGBARainbow(step int) (rainbow color.Palette, err error) {
	if step < 7 {
		return rainbow, fmt.Errorf(
			"step must be greater than 6, got: %d. Palette cannot hold more than 256 colors",
			step,
		)
	}

	// Start at Red
	rgba := color.RGBA{255, 0, 0, 255}
	rainbow = append(rainbow, rgba)

	// Red -> Yellow
	for rgba.G < 255 {
		nextValue := int(rgba.G) + step
		if nextValue > 255 {
			nextValue = 255
		}

		rgba.G = uint8(nextValue)
		rainbow = append(rainbow, rgba)
	}

	// Yellow -> Green
	for rgba.R > 0 {
		nextValue := int(rgba.R) - step
		if nextValue < 0 {
			nextValue = 0
		}

		rgba.R = uint8(nextValue)
		rainbow = append(rainbow, rgba)
	}

	// Green -> Cyan
	for rgba.B < 255 {
		nextValue := int(rgba.B) + step
		if nextValue > 255 {
			nextValue = 255
		}

		rgba.B = uint8(nextValue)
		rainbow = append(rainbow, rgba)
	}

	// Cyan -> Blue
	for rgba.G > 0 {
		nextValue := int(rgba.G) - step
		if nextValue < 0 {
			nextValue = 0
		}

		rgba.G = uint8(nextValue)
		rainbow = append(rainbow, rgba)
	}

	// Blue -> Magenta
	for rgba.R < 255 {
		nextValue := int(rgba.R) + step
		if nextValue > 255 {
			nextValue = 255
		}

		rgba.R = uint8(nextValue)
		rainbow = append(rainbow, rgba)
	}

	// Magenta -> Red
	for rgba.B > 0 {
		nextValue := int(rgba.B) - step
		if nextValue < 0 {
			nextValue = 0
		}

		rgba.B = uint8(nextValue)
		rainbow = append(rainbow, rgba)
	}

	return rainbow, nil
}
