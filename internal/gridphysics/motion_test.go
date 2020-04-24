package gridphysics_test

import (
	"testing"

	"github.com/domtriola/automata/internal/gridphysics"
	"github.com/stretchr/testify/assert"
)

func TestCoordinate(t *testing.T) {
	t.Parallel()

	t.Run("Move() returns the correct coordinate", func(t *testing.T) {
		t.Parallel()

		type testCases struct {
			name     string
			coord    gridphysics.Coordinate
			vect     gridphysics.AngleVector
			expected gridphysics.Coordinate
		}

		tCases := []testCases{
			{
				name:     "no movement",
				coord:    gridphysics.Coordinate{0, 0},
				vect:     gridphysics.AngleVector{gridphysics.DegreeAngle(0).ToRadians(), 0},
				expected: gridphysics.Coordinate{0, 0},
			},
			{
				name:     "just magnitude",
				coord:    gridphysics.Coordinate{0, 0},
				vect:     gridphysics.AngleVector{gridphysics.DegreeAngle(0).ToRadians(), 50},
				expected: gridphysics.Coordinate{50, 0},
			},
			{
				name:     "up",
				coord:    gridphysics.Coordinate{20, 15},
				vect:     gridphysics.AngleVector{gridphysics.DegreeAngle(90).ToRadians(), 15},
				expected: gridphysics.Coordinate{20, 30},
			},
			{
				name:     "to the left",
				coord:    gridphysics.Coordinate{20, 50},
				vect:     gridphysics.AngleVector{gridphysics.DegreeAngle(180).ToRadians(), 10},
				expected: gridphysics.Coordinate{10, 50},
			},
			{
				name:     "full circle",
				coord:    gridphysics.Coordinate{0, 0},
				vect:     gridphysics.AngleVector{gridphysics.DegreeAngle(360).ToRadians(), 10},
				expected: gridphysics.Coordinate{10, 0},
			},
			{
				name:     "more than full circle",
				coord:    gridphysics.Coordinate{0, 0},
				vect:     gridphysics.AngleVector{gridphysics.DegreeAngle(450).ToRadians(), 10},
				expected: gridphysics.Coordinate{0, 10},
			},
			{
				name:     "negative magnitude",
				coord:    gridphysics.Coordinate{20, 40},
				vect:     gridphysics.AngleVector{gridphysics.DegreeAngle(270).ToRadians(), -30},
				expected: gridphysics.Coordinate{20, 70},
			},
			{
				name:     "negative coord",
				coord:    gridphysics.Coordinate{10, 10},
				vect:     gridphysics.AngleVector{gridphysics.DegreeAngle(225).ToRadians(), 30},
				expected: gridphysics.Coordinate{-11.213203, -11.213203},
			},
		}

		for _, tc := range tCases {
			t.Run(tc.name, func(t *testing.T) {
				nextCoord := tc.coord.Move(tc.vect)
				assert.InDelta(t, tc.expected[0], nextCoord[0], 0.00001, "unexpected x coord")
				assert.InDelta(t, tc.expected[1], nextCoord[1], 0.00001, "unexpected y coord")
			})
		}
	})
}
