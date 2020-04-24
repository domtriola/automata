package models

import (
	"github.com/domtriola/automata/internal/gridphysics"
)

// Organism represents one creature that can be displayed at one Space
type Organism struct {
	ID int

	CAFeatures *CAFeatures
	SMFeatures *SMFeatures
}

// CAFeatures holds features that are unique to the cellular automata
// simulation.
type CAFeatures struct {
	// SpeciesID identifies the species type of a cell
	SpeciesID     int
	NextSpeciesID int

	// Neighbors stores pointers to an organism's neighbors
	Neighbors []*Organism
}

// SMFeatures holds features that are unique to the slime mold simulation.
type SMFeatures struct {
	// XPos is a float that represents an organisms location within the Grid.
	// Positions are floats so that organisms can be at locations that are between
	// Spaces. This allows them to travel in any direction (otherwise they would
	// be limited to 45˚ increments). Their visible location will still be one
	// discrete Space (whichever Space they are closest to).
	XPos     float64
	NextXPos float64

	// YPos is a float that represents an organisms location within the Grid.
	// Positions are floats so that organisms can be at locations that are between
	// Spaces. This allows them to travel in any direction (otherwise they would
	// be limited to 45˚ increments). Their visible location will still be one
	// discrete Space (whichever Space they are closest to).
	YPos     float64
	NextYPos float64

	// direction represents the angular direction (in degrees) that an organism
	// is facing.
	Direction gridphysics.DegreeAngle
}

// NewOrganism instantiates and returns a new Organism
func NewOrganism(id int) *Organism {
	return &Organism{
		ID:         id,
		CAFeatures: &CAFeatures{},
		SMFeatures: &SMFeatures{},
	}
}
