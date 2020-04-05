package simulations

import "github.com/domtriola/automata-gen/internal/models"

var _ models.Simulation = (*SlimeMold)(nil)

// SlimeMold simulates a slime mold that leaves behind scent trails and creates
// networks based on where other mold particles have been.
type SlimeMold struct{}

// Build creates the simulation
func (s *SlimeMold) Build() error {
	return nil
}
