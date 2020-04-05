package models

import (
	"fmt"
)

// Organism represents one creature that can be displayed at one space
type Organism struct {
	ID        int
	Direction float64

	// XPos is a float that represents an organisms location within the grid.
	// Positions are floats so that organisms can be at locations that are between
	// spaces. This allows them to travel in any direction (otherwise they would
	// be limited to 45˚ increments). Their visible location will still be one
	// discrete space (whichever space they are closest to).
	XPos float64

	// YPos is a float that represents an organisms location within the grid.
	// Positions are floats so that organisms can be at locations that are between
	// spaces. This allows them to travel in any direction (otherwise they would
	// be limited to 45˚ increments). Their visible location will still be one
	// discrete space (whichever space they are closest to).
	YPos float64

	NextDiscreteXPos int
	NextDiscreteYPos int
}

// OrganismFeatures contains optional features of an organism that apply to
// certain simulations.
type OrganismFeatures struct {
	// direction represents the angular direction an organism is facing. A
	// direction will be between 0 and 360.
	direction float64
}

// GetDirection returns the direction an organism is facing.
func (f OrganismFeatures) GetDirection() float64 {
	return f.direction
}

// SetDirection sets the direction an organism is facing. The direction is an
// angular direction between 0 and 360.
func (f OrganismFeatures) SetDirection(dir float64) error {
	if dir < 0 || dir > 360 {
		return fmt.Errorf(
			"direction must be an angular direction between 0 and 360. Received: %f",
			dir,
		)
	}

	f.direction = dir

	return nil
}
