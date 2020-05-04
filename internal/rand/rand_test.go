package rand_test

import (
	"testing"

	"github.com/domtriola/automata/internal/rand"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRand(t *testing.T) {
	t.Parallel()

	implementations := []string{"math", "crypto"}

	t.Run("NewRand() returns an error if implementation not recognized", func(t *testing.T) {
		_, err := rand.NewRand("notreal")
		require.Error(t, err)
	})

	t.Run("Int() generates a number in [0, max)", func(t *testing.T) {
		for _, implementation := range implementations {
			imp := implementation

			t.Run(imp, func(t *testing.T) {
				r, err := rand.NewRand(imp)
				require.NoError(t, err)

				for i := 0; i < 100; i++ {
					num, err := r.Int(2)
					require.NoError(t, err)

					assert.GreaterOrEqual(t, num, int64(0))
					assert.LessOrEqual(t, num, int64(1))
				}
			})
		}
	})
}
