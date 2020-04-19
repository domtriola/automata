package palette_test

import (
	"testing"

	"github.com/domtriola/automata/pkg/palette"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRainbow(t *testing.T) {
	t.Parallel()

	t.Run("Rainbow() returns the right amount of colors", func(t *testing.T) {
		p, err := palette.Rainbow(7)
		require.NoError(t, err)

		assert.Equal(t, 223, len(p), "unexpected amount of colors")

		p, err = palette.Rainbow(255)
		require.NoError(t, err)

		assert.Equal(t, 7, len(p), "unexpected amount of colors")
	})

	t.Run("Rainbow() returns an error for invalid step inputs", func(t *testing.T) {
		_, err := palette.Rainbow(6)
		require.Error(t, err)

		_, err = palette.Rainbow(256)
		require.Error(t, err)
	})

	t.Skip("Rainbow() returns a palette with less than 256 colors")
	t.Skip("Rainbow() returns evenly spaced colors for any step increments")
}
