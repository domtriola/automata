package color_test

import (
	"testing"

	"github.com/domtriola/automata-gen/pkg/color"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRGBARainbow(t *testing.T) {
	t.Parallel()

	t.Run("RGBARainbow() returns the right amount of colors", func(t *testing.T) {
		p, err := color.RGBARainbow(7)
		require.NoError(t, err)

		assert.Equal(t, 223, len(*p), "unexpected amount of colors")

		p, err = color.RGBARainbow(255)
		require.NoError(t, err)

		assert.Equal(t, 7, len(*p), "unexpected amount of colors")
	})

	t.Skip("RGBARainbow() returns evenly spaced colors for any step increments")
}
