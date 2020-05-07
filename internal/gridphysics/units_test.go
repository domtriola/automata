package gridphysics_test

import (
	"testing"

	"github.com/domtriola/automata/internal/gridphysics"
	"github.com/stretchr/testify/assert"
)

func TestDegreeAngle(t *testing.T) {
	t.Parallel()

	t.Run("DegreeAngle.ToRadians() returns correct values", func(t *testing.T) {
		type testCase struct {
			input    gridphysics.DegreeAngle
			expected gridphysics.RadianAngle
		}

		testCases := []testCase{
			{
				input:    gridphysics.DegreeAngle(0),
				expected: gridphysics.RadianAngle(0),
			},
			{
				input:    gridphysics.DegreeAngle(15),
				expected: gridphysics.RadianAngle(0.2617993877991494),
			},
			{
				input:    gridphysics.DegreeAngle(90),
				expected: gridphysics.RadianAngle(1.5707963267948966),
			},
			{
				input:    gridphysics.DegreeAngle(-15),
				expected: gridphysics.RadianAngle(-0.2617993877991494),
			},
			{
				input:    gridphysics.DegreeAngle(360),
				expected: gridphysics.RadianAngle(6.283185307179586),
			},
			{
				input:    gridphysics.DegreeAngle(1000),
				expected: gridphysics.RadianAngle(17.453292519943293),
			},
			{
				input:    gridphysics.DegreeAngle(1.2345),
				expected: gridphysics.RadianAngle(0.02154608961587),
			},
		}

		for _, tc := range testCases {
			assert.Equal(t, tc.expected, tc.input.ToRadians())
		}
	})
}
