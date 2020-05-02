package models_test

import (
	"testing"

	"github.com/domtriola/automata/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestNewOrgansim(t *testing.T) {
	t.Parallel()

	t.Run("NewOrganism() initializes with the given ID", func(t *testing.T) {
		t.Parallel()

		o := models.NewOrganism(123)
		assert.Equal(t, 123, o.ID, "organism should have ID that was set at initialization")
	})
}
