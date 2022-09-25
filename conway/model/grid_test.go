package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid_Height(t *testing.T) {
	t.Run("when initialized with 2x3 returns 2", func(t *testing.T) {
		twoByThree := map[int]map[int]Void{
			0: {0: Void{}, 1: Void{}, 2: Void{}},
			1: {0: Void{}, 1: Void{}, 2: Void{}},
		}
		grid := NewGrid(twoByThree)

		assert.Equal(t, 2, grid.Height())
	})

	t.Run("when initialized with 3x3 returns 3", func(t *testing.T) {
		threeByThree := map[int]map[int]Void{
			0: {0: Void{}, 1: Void{}, 2: Void{}},
			1: {0: Void{}, 1: Void{}, 2: Void{}},
			2: {0: Void{}, 1: Void{}, 2: Void{}},
		}
		grid := NewGrid(threeByThree)

		assert.Equal(t, 3, grid.Height())
	})
}

func TestGrid_Width(t *testing.T) {
	t.Run("when initialized with 2x2 returns 2", func(t *testing.T) {
		twoByTwo := map[int]map[int]Void{
			0: {0: Void{}, 1: Void{}},
			1: {1: Void{}},
		}
		grid := NewGrid(twoByTwo)

		assert.Equal(t, 2, grid.Width())
	})
}
