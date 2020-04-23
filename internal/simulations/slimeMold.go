package simulations

import (
	"crypto/rand"
	"image"
	"image/color"
	"math/big"

	"github.com/domtriola/automata/internal/models"
	"github.com/domtriola/automata/internal/palette"
)

const (
	// A multiplier of 100 means scents of at least 0.01 will be barely visible,
	// since they will be shown at the first index of the grey scale. It also
	// means that scents over 2.55 will have maximum visibility.
	// e.g. 100 x 0.01 == 1 (color index) == RGB(1, 1, 1)
	// e.g. 100 x 2.55 == 255 (color index) == RGB(255, 255, 255)
	scentVisibilityMultiplier = 100

	// initOrganismChance is the chance that an organism will exist at any given
	// space when the simulation is initialized.
	initOrganismChance = 1.0 / 20
)

var _ models.Simulation = &SlimeMold{}

// SlimeMold simulates a slime mold that leaves behind scent trails and creates
// networks based on where other mold particles have been.
type SlimeMold struct {
	cfg     SlimeMoldConfig
	palette color.Palette
}

// SlimeMoldConfig holds the configurations for the slime mold simulation
type SlimeMoldConfig struct {
}

// NewSlimeMold initializes and returns a new slime mold simulation
func NewSlimeMold(cfg models.SimulationConfig) (*SlimeMold, error) {
	s := &SlimeMold{cfg: SlimeMoldConfig{}}

	err := s.setPalette()
	if err != nil {
		return &SlimeMold{}, nil
	}

	return s, nil
}

// OutputName creates an output file path based on parameters of the
// simulation
func (s *SlimeMold) OutputName() (string, error) {
	return "slime_mold.gif", nil
}

// InitializeGrid instantiates a grid
func (s *SlimeMold) InitializeGrid(g *models.Grid) error {
	oID := 0

	for _, row := range g.Rows {
		for _, space := range row {
			space.Features = &models.SpaceFeatures{}

			r, err := rand.Int(rand.Reader, big.NewInt(int64(1/initOrganismChance)))
			if err != nil {
				return err
			}

			shouldPopulate := r.Cmp(big.NewInt(0)) == 0
			if shouldPopulate {
				o := models.NewOrganism(oID)
				space.Organism = o
				oID++
			}
		}
	}

	return nil
}

// AdvanceFrame determines and assigns the next state of each organism's
// parameters.
func (s *SlimeMold) AdvanceFrame(g *models.Grid) error {
	s.calculateNextFrame(g)
	s.applyNextFrame(g)

	return nil
}

// DrawSpace colors the image at the specified location according to the
// properties of the Space.
func (s *SlimeMold) DrawSpace(
	sp *models.Space,
	img *image.Paletted,
	x int,
	y int,
) error {
	paletteMax := uint8(len(img.Palette)) - 1

	var colorIndex uint8

	if sp.Organism != nil {
		colorIndex = paletteMax
	} else {
		colorIndex = uint8(sp.Features.Scent * scentVisibilityMultiplier)
	}

	if colorIndex >= paletteMax {
		colorIndex = paletteMax
	}

	img.SetColorIndex(x, y, colorIndex)

	return nil
}

// GetPalette returns the simulation's color palette
func (s *SlimeMold) GetPalette() color.Palette {
	return s.palette
}

func (s *SlimeMold) setPalette() error {
	s.palette = palette.Grey()

	return nil
}

func (s *SlimeMold) calculateNextFrame(g *models.Grid) {}

func (s *SlimeMold) applyNextFrame(g *models.Grid) {
	for _, row := range g.Rows {
		for _, space := range row {
			if space.Organism == nil {
				continue
			}

			// vel := Velocity{
			// 	direction: organism.direction,
			// 	speed:     1,
			// }
			// nextX, nextY := NextPos(organism.xPos, organism.yPos, vel)
			// nextDiscreteX, nextDiscreteY := FloatPosToGridPos(nextX, nextY)

			// organism.xPos = nextX
			// organism.yPos = nextY

			// if grid.hasCoord(nextDiscreteX, nextDiscreteY) {
			// 	organism.nextDiscreteXPos = nextDiscreteX
			// 	organism.nextDiscreteYPos = nextDiscreteY
			// } else {
			// 	space.organism = nil
			// }
		}
	}
}

// TODODOM: incorporate
// func moveOrganisms(grid Grid) {
// 	for _, row := range grid.rows {
// 		for _, space := range row {
// 			space.scent *= scentDecay
// 			// TODO: scent spread

// 			if space.organism != nil {
// 				organism := space.organism
// 				destinationSpace := grid.rows[organism.nextDiscreteXPos][organism.nextDiscreteYPos]

// 				// Movement Step
// 				// Move if possible, otherwise change directions
// 				moveOrganism(organism, space, destinationSpace)

// 				// Sensory Step
// 				// Find the direction with the largest scent trail and turn in that direction
// 				rotateOrganism(organism, grid)
// 			}
// 		}
// 	}
// }

// TODODOM: incorporate
// func moveOrganism(organism *Organism, currentSpace *Space, destinationSpace *Space) {
// 	if destinationSpace.organism == nil {
// 		// NOTE: This is a greedy approach
// 		// Setting this to nil before calculating the rest of the movements
// 		// means the top-left of the screen becomes more dense than the
// 		// bottom-right because it opens up movement options to the upper-left
// 		// zones before the lower right zones
// 		currentSpace.organism = nil
// 		destinationSpace.organism = organism
// 		destinationSpace.scent++
// 	} else if destinationSpace.organism.id != organism.id {
// 		organism.direction = float64(rand.Intn(360))
// 	}
// }

// TODODOM: incorporate
// func rotateOrganism(organism *Organism, grid Grid) {
// 	sensorDistance := float64(options["sensorDistance"])
// 	sensorDegree := float64(options["sensorDegree"])
// 	leftSensorDirection := organism.direction + sensorDegree
// 	if leftSensorDirection > 360 {
// 		leftSensorDirection -= 360
// 	}
// 	rightSensorDirection := organism.direction - sensorDegree
// 	if rightSensorDirection < 0 {
// 		rightSensorDirection += 360
// 	}
// 	leftSensorVelocity := Velocity{
// 		speed:     sensorDistance,
// 		direction: leftSensorDirection,
// 	}
// 	rightSensorVelocity := Velocity{
// 		speed:     sensorDistance,
// 		direction: rightSensorDirection,
// 	}
// 	frontSensorVelocity := Velocity{
// 		speed:     sensorDistance,
// 		direction: organism.direction,
// 	}
// 	leftSensorX, leftSensorY := NextPos(organism.xPos, organism.yPos, leftSensorVelocity)
// 	rightSensorX, rightSensorY := NextPos(organism.xPos, organism.yPos, rightSensorVelocity)
// 	frontSensorX, frontSensorY := NextPos(organism.xPos, organism.yPos, frontSensorVelocity)
// 	leftSensorDiscreteX, leftSensorDiscreteY := FloatPosToGridPos(leftSensorX, leftSensorY)
// 	rightSensorDiscreteX, rightSensorDiscreteY := FloatPosToGridPos(rightSensorX, rightSensorY)
// 	frontSensorDiscreteX, frontSensorDiscreteY := FloatPosToGridPos(frontSensorX, frontSensorY)

// 	var leftSensorSpace *Space
// 	var rightSensorSpace *Space
// 	var frontSensorSpace *Space
// 	if grid.hasCoord(leftSensorDiscreteX, leftSensorDiscreteY) {
// 		leftSensorSpace = grid.rows[leftSensorDiscreteY][leftSensorDiscreteX]
// 	}
// 	if grid.hasCoord(rightSensorDiscreteX, rightSensorDiscreteY) {
// 		rightSensorSpace = grid.rows[rightSensorDiscreteY][rightSensorDiscreteX]
// 	}
// 	if grid.hasCoord(frontSensorDiscreteX, frontSensorDiscreteY) {
// 		frontSensorSpace = grid.rows[frontSensorDiscreteY][frontSensorDiscreteX]
// 	}

// 	if leftSensorSpace == nil || rightSensorSpace == nil || frontSensorSpace == nil {
// 		organism.direction = float64(rand.Intn(360))
// 	} else {
// 		leftScent := leftSensorSpace.scent
// 		rightScent := rightSensorSpace.scent
// 		frontScent := frontSensorSpace.scent

// 		if frontScent > leftScent && frontScent > rightScent {
// 			// Continue in same direction
// 		} else if leftScent > frontScent && rightScent > frontScent {
// 			// Rotate randomly left or right
// 			toss := rand.Intn(2)
// 			if toss == 0 {
// 				organism.direction += sensorDegree
// 				if organism.direction > 360 {
// 					organism.direction -= 360
// 				}
// 			} else {
// 				organism.direction -= sensorDegree
// 				if organism.direction < 0 {
// 					organism.direction += 360
// 				}
// 			}
// 		} else if leftScent > rightScent {
// 			// Rotate left
// 			organism.direction += sensorDegree
// 			if organism.direction > 360 {
// 				organism.direction -= 360
// 			}
// 		} else if rightScent > leftScent {
// 			// Rotate right
// 			organism.direction -= sensorDegree
// 			if organism.direction < 0 {
// 				organism.direction += 360
// 			}
// 		}
// 		// Else continue in same direction
// 	}
// }
