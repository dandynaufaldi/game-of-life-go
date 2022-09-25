package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid_Height(t *testing.T) {
	t.Run("when initialized with 2x3 returns 2", func(t *testing.T) {
		twoByThree := [][]bool{
			{true, true, true},
			{true, true, true},
		}
		grid := NewGrid(twoByThree)

		assert.Equal(t, 2, grid.Height())
	})
}