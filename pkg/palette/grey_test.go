package palette_test

import (
	"testing"

	"github.com/domtriola/automata/pkg/palette"
	"github.com/stretchr/testify/assert"
)

func TestGrey(t *testing.T) {
	t.Parallel()

	t.Run("Grey() returns a grey palette", func(t *testing.T) {
		p := palette.Grey()

		for _, c := range p {
			r, g, b, _ := c.RGBA()
			assert.Equal(t, r, g)
			assert.Equal(t, g, b)
		}
	})

	t.Run("Grey() returns a palette that has 255 characters or less", func(t *testing.T) {
		p := palette.Grey()

		assert.Less(t, len(p), 256)
	})
}
