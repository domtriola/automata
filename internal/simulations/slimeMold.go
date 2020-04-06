package simulations

import "github.com/domtriola/automata-gen/internal/models"

var _ models.Simulation = (*SlimeMold)(nil)

// SlimeMold simulates a slime mold that leaves behind scent trails and creates
// networks based on where other mold particles have been.
type SlimeMold struct{}

// CalculateNextFrame determines and assigns the next state of each organism's
// parameters.
func (s *SlimeMold) CalculateNextFrame() error {
	return nil
}
