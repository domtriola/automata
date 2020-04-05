package models

// Space represents one discrete location that may be occupied by an organism.
type Space struct {
	Organism *Organism
	Features *SpaceFeatures
}

// SpaceFeatures contains optional features of a space that apply to certain
// simulations.
type SpaceFeatures struct {
	// Scent represents the strength of a scent that may be left behind by organisms.
	Scent float64
}
