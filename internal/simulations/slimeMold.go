package simulations

import (
	"crypto/rand"
	"image"
	"image/color"
	"math"
	"math/big"

	"github.com/domtriola/automata/internal/gridphysics"
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

	// TODODOM: make an option
	// organismMovementSpeed is the rate of movement for all organisms.
	organismMovementSpeed = 1

	// TODODOM: make an option
	// scentDecay is the rate at which a scent decays from a space.
	scentDecay = 0.9

	// TODODOM: make an option
	// scentSpreadFactor is the percentage of scent that dissapates to neighboring
	// spaces.
	scentSpreadFactor = 0.1

	// TODODOM: make options
	sensorDistance = 9
	sensorDegree   = 25
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

	s.setPalette()

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

	for y, row := range g.Rows {
		for x, space := range row {
			space.Features = &models.SpaceFeatures{}

			r, err := rand.Int(rand.Reader, big.NewInt(int64(1/initOrganismChance)))
			if err != nil {
				return err
			}

			shouldPopulate := r.Cmp(big.NewInt(0)) == 0
			if shouldPopulate {
				o := models.NewOrganism(oID)
				o.Features.XPos = float64(x)
				o.Features.YPos = float64(y)
				o.NextFeatures.XPos = float64(x)
				o.NextFeatures.YPos = float64(y)

				r, err := rand.Int(rand.Reader, big.NewInt(360))
				if err != nil {
					return err
				}

				o.Features.Direction = gridphysics.DegreeAngle(r.Int64())

				space.Organism = o

				oID++
			}
		}
	}

	for y, row := range g.Rows {
		for x, space := range row {
			neighbors := g.GetNeighbors(x, y, []string{"nw", "n", "ne", "e", "se", "s", "sw", "w"})
			space.Neighbors = neighbors
		}
	}

	return nil
}

// AdvanceFrame determines and assigns the next state of each organism's
// parameters.
func (s *SlimeMold) AdvanceFrame(g *models.Grid) error {
	if err := s.calculateNextFrame(g); err != nil {
		return err
	}

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

func (s *SlimeMold) setPalette() {
	s.palette = palette.Grey()
}

func (s *SlimeMold) calculateNextFrame(g *models.Grid) error {
	if err := rotateOrganisms(g); err != nil {
		return err
	}

	setNextPositions(g)
	disperseScentTrails(g)

	return nil
}

func (s *SlimeMold) applyNextFrame(g *models.Grid) {
	for _, row := range g.Rows {
		for _, space := range row {
			if space.Organism == nil {
				continue
			}

			space.Organism.Features.XPos = space.Organism.NextFeatures.XPos
			space.Organism.Features.YPos = space.Organism.NextFeatures.YPos

			destCoord := gridphysics.Coordinate{
				space.Organism.Features.XPos,
				space.Organism.Features.YPos,
			}.ToDiscreteCoordinate()

			dest, ok := g.GetSpace(int(destCoord[0]), int(destCoord[1]))
			if !ok {
				continue
			}

			// TODODOM: this is a greedy approach. A better one might be something
			// more like:
			// spaceAtOrganismCoord.Organisms = append(space.Organisms, space.Organism)
			// space.Organisms = pop(space.Organism)
			if dest.Organism == nil {
				dest.Organism = space.Organism
				space.Organism = nil
			}
		}
	}
}

func rotateOrganisms(g *models.Grid) error {
	for _, row := range g.Rows {
		for _, space := range row {
			if space.Organism == nil {
				continue
			}

			if err := rotateOrganism(g, space.Organism); err != nil {
				return err
			}
		}
	}

	return nil
}

func rotateOrganism(g *models.Grid, o *models.Organism) error {
	o.Features.Direction = reduceDirection(o.Features.Direction)

	lSpace, mSpace, rSpace := getScensorSpaces(g, o)

	if lSpace == nil || rSpace == nil || mSpace == nil {
		// Continue in the same direction. Movement phase will turn organism
		// around if it can't proceed further.
		// TODODOM: be more clever - maybe rotate opposite the nil direction?
		// Otherwise organisms could get in a swim-lane loop near the edge.
		return nil
	}

	lScent := lSpace.Features.Scent
	rScent := rSpace.Features.Scent
	mScent := mSpace.Features.Scent

	// If the middle scensor has the greatest scent, proceed straight ahead
	if mScent > lScent && mScent > rScent {
		return nil
	}

	// If scents to the left and right are greater than straight ahead, pick
	// left or right randomly
	if lScent > mScent && rScent > mScent {
		r, err := rand.Int(rand.Reader, big.NewInt(2))
		if err != nil {
			return err
		}

		if r.Cmp(big.NewInt(0)) == 0 {
			o.Features.Direction += sensorDegree
		} else {
			o.Features.Direction -= sensorDegree
		}

		return nil
	}

	// If left scent is greater than right, turn left
	if lScent > rScent {
		o.Features.Direction += sensorDegree
		return nil
	}

	// If right scent is greater than left, turn right
	if rScent > lScent {
		o.Features.Direction -= sensorDegree
		return nil
	}

	// If no scent is greater, then proceed straight ahead
	return nil
}

// reduceDirection ensures that an organism's direction doesn't overflow float64
func reduceDirection(d gridphysics.DegreeAngle) gridphysics.DegreeAngle {
	buffer := 1000000.0
	df := float64(d)

	if df > math.MaxFloat64-buffer || df < math.SmallestNonzeroFloat64+buffer {
		return gridphysics.DegreeAngle(math.Mod(df, 360))
	}

	return d
}

func getScensorSpaces(g *models.Grid, o *models.Organism) (lSpace, mSpace, rSpace *models.Space) {
	ld := o.Features.Direction + sensorDegree
	lv := gridphysics.AngleVector{
		Direction: ld.ToRadians(),
		Magnitude: sensorDistance,
	}

	rd := o.Features.Direction - sensorDegree
	rv := gridphysics.AngleVector{
		Direction: rd.ToRadians(),
		Magnitude: sensorDistance,
	}

	mv := gridphysics.AngleVector{
		Direction: o.Features.Direction.ToRadians(),
		Magnitude: sensorDistance,
	}

	coord := gridphysics.Coordinate{o.Features.XPos, o.Features.YPos}
	lCoord := coord.Move(lv).ToDiscreteCoordinate()
	rCoord := coord.Move(rv).ToDiscreteCoordinate()
	mCoord := coord.Move(mv).ToDiscreteCoordinate()

	lSpace, _ = g.GetSpace(int(lCoord[0]), int(rCoord[1]))
	rSpace, _ = g.GetSpace(int(rCoord[0]), int(rCoord[1]))
	mSpace, _ = g.GetSpace(int(mCoord[0]), int(mCoord[1]))

	return lSpace, mSpace, rSpace
}

func setNextPositions(g *models.Grid) {
	for _, row := range g.Rows {
		for _, space := range row {
			if space.Organism == nil {
				continue
			}

			o := space.Organism

			vect := gridphysics.AngleVector{
				Direction: o.Features.Direction.ToRadians(),
				Magnitude: organismMovementSpeed,
			}

			coord := gridphysics.Coordinate{o.Features.XPos, o.Features.YPos}
			nextCoord := coord.Move(vect)

			if g.HasCoord(int(nextCoord[0]), int(nextCoord[1])) {
				o.NextFeatures.XPos = nextCoord[0]
				o.NextFeatures.YPos = nextCoord[1]
			} else {
				o.Features.Direction += 180
			}
		}
	}
}

func disperseScentTrails(g *models.Grid) {
	for _, row := range g.Rows {
		for _, space := range row {
			space.Features.Scent *= scentDecay

			// for _, n := range space.Neighbors {
			// 	n.Features.Scent += space.Features.Scent * scentSpreadFactor
			// }

			if space.Organism != nil {
				space.Features.Scent++
			}
		}
	}
}
