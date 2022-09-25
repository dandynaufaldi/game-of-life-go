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

	t.Run("when initialized with 2x3 returns 3", func(t *testing.T) {
		twoByThree := map[int]map[int]Void{
			0: {0: Void{}, 1: Void{}},
			1: {1: Void{}, 2: Void{}},
		}
		grid := NewGrid(twoByThree)

		assert.Equal(t, 3, grid.Width())
	})
}

func TestGrid_IsAlive(t *testing.T) {
	t.Run("when cell at given row and column is alive returns true", func(t *testing.T) {
		twoByThree := map[int]map[int]Void{
			0: {0: Void{}, 1: Void{}},
			1: {1: Void{}, 2: Void{}},
		}
		grid := NewGrid(twoByThree)

		assert.True(t, grid.IsAlive(0, 1))
	})

	t.Run("when cell at given row and column is dead returns false", func(t *testing.T) {
		twoByThree := map[int]map[int]Void{
			0: {0: Void{}, 1: Void{}},
			1: {1: Void{}, 2: Void{}},
		}
		grid := NewGrid(twoByThree)

		assert.False(t, grid.IsAlive(0, 2))
	})
}

func TestGrid_String(t *testing.T) {
	t.Run("when cell is alive returns hashtag", func(t *testing.T) {
		initialState := map[int]map[int]Void{
			0: {0: Void{}, 1: Void{}},
		}
		grid := NewGrid(initialState)

		assert.Equal(t, "##", grid.String())
	})

	t.Run("when cell is dead returns space", func(t *testing.T) {
		initialState := map[int]map[int]Void{
			0: {1: Void{}},
		}
		grid := NewGrid(initialState)

		assert.Equal(t, " #", grid.String())
	})
}
