package gridphysics_test

import (
	"testing"

	"github.com/domtriola/automata/internal/gridphysics"
	"github.com/stretchr/testify/assert"
)

// nolint: funlen
func TestCoordinate(t *testing.T) {
	t.Parallel()

	t.Run("Move() returns the correct coordinate", func(t *testing.T) {
		t.Parallel()

		type testCase struct {
			name     string
			coord    gridphysics.Coordinate
			vect     gridphysics.AngleVector
			expected gridphysics.Coordinate
		}

		testCases := []testCase{
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

		for _, tCase := range testCases {
			tc := tCase

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				nextCoord := tc.coord.Move(tc.vect)
				assert.InDelta(t, tc.expected[0], nextCoord[0], 0.00001, "unexpected x coord")
				assert.InDelta(t, tc.expected[1], nextCoord[1], 0.00001, "unexpected y coord")
			})
		}
	})

	t.Run("ToDiscreteCoordinate() returns the correct coordinate", func(t *testing.T) {
		t.Parallel()

		type testCase struct {
			name     string
			coord    gridphysics.Coordinate
			expected gridphysics.DiscreteCoord
		}

		testCases := []testCase{
			{name: "zero", coord: gridphysics.Coordinate{0, 0}, expected: gridphysics.DiscreteCoord{0, 0}},
		}

		for _, tCase := range testCases {
			tc := tCase

			t.Run(tc.name, func(t *testing.T) {
				dCoord := tc.coord.ToDiscreteCoordinate()

				assert.Equal(t, tc.expected[0], dCoord[0], "unexpected x coordinate")
				assert.Equal(t, tc.expected[1], dCoord[1], "unexpected y coordinate")
			})
		}
	})
}

func TestVector(t *testing.T) {
	t.Parallel()

	t.Run("LinearVelocity() returns the correct values", func(t *testing.T) {
		t.Parallel()

		type testCase struct {
			name     string
			vect     gridphysics.AngleVector
			expected [2]float64
		}

		testCases := []testCase{
			{name: "zero", vect: gridphysics.AngleVector{0, 0}, expected: [2]float64{0, 0}},
			{
				name:     "up",
				vect:     gridphysics.AngleVector{gridphysics.DegreeAngle(90).ToRadians(), 1},
				expected: [2]float64{0, 1},
			},
			{
				name:     "over",
				vect:     gridphysics.AngleVector{gridphysics.DegreeAngle(36.87).ToRadians(), 5},
				expected: [2]float64{4, 3},
			},
		}

		for _, tCase := range testCases {
			tc := tCase

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				x, y := tc.vect.LinearVelocity()

				assert.InDelta(t, tc.expected[0], x, 0.0001, "unexpected x coordinate")
				assert.InDelta(t, tc.expected[1], y, 0.0001, "unexpected y coordinate")
			})
		}
	})
}
