package models

import (
	"github.com/domtriola/automata/internal/gridphysics"
)

// Organism represents one creature that can be displayed at one Space
type Organism struct {
	ID int

	Features     *Features
	NextFeatures *NextFeatures
}

// Features are various parameters that define an organism. Not all features are
// used in every simulation.
type Features struct {
	// Cellular Automata features

	// SpeciesID identifies the species type of a cell
	SpeciesID int
	// Neighbors stores pointers to an organism's neighbors
	Neighbors []*Organism

	// Slime Mold features

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

	// direction represents the angular direction (in degrees) that an organism
	// is facing.
	Direction gridphysics.DegreeAngle
}

// NextFeatures holds the next state of certain features, so we can calculate
// every state transition and apply them all at once.
type NextFeatures struct {
	SpeciesID int
	XPos      float64
	YPos      float64
}

// NewOrganism instantiates and returns a new Organism
func NewOrganism(id int) *Organism {
	return &Organism{
		ID:           id,
		Features:     &Features{},
		NextFeatures: &NextFeatures{},
	}
}
