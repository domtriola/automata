package models

import (
	"fmt"
)

// Organism represents one creature that can be displayed at one Space
type Organism struct {
	ID int

	// XPos is a float that represents an organisms location within the Grid.
	// Positions are floats so that organisms can be at locations that are between
	// Spaces. This allows them to travel in any direction (otherwise they would
	// be limited to 45˚ increments). Their visible location will still be one
	// discrete Space (whichever Space they are closest to).
	XPos float64

	// YPos is a float that represents an organisms location within the Grid.
	// Positions are floats so that organisms can be at locations that are between
	// Spaces. This allows them to travel in any direction (otherwise they would
	// be limited to 45˚ increments). Their visible location will still be one
	// discrete Space (whichever Space they are closest to).
	YPos float64

	NextDiscreteXPos int
	NextDiscreteYPos int

	// Features holds features that may be unique to only certain simulations.
	Features *OrganismFeatures
}

// OrganismFeatures contains optional features of an organism that apply to
// certain simulations.
type OrganismFeatures struct {
	// SpeciesID identifies the species type of a cell
	SpeciesID int

	// direction represents the angular direction an organism is facing. A
	// direction will be between 0 and 360.
	direction float64
}

// NewOrganism instantiates and returns a new Organism
func NewOrganism(id int) *Organism {
	return &Organism{
		ID:       id,
		Features: &OrganismFeatures{},
	}
}

// GetDirection returns the direction an organism is facing.
func (f *OrganismFeatures) GetDirection() float64 {
	return f.direction
}

// SetDirection sets the direction an organism is facing. The direction is an
// angular direction between 0 and 360.
func (f *OrganismFeatures) SetDirection(dir float64) error {
	if dir < 0 || dir > 360 {
		return fmt.Errorf(
			"direction must be an angular direction between 0 and 360. Received: %f",
			dir,
		)
	}

	f.direction = dir

	return nil
}
