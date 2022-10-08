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

		assert.Equal(t, "Board\n##", grid.String())
	})

	t.Run("when cell is dead returns space", func(t *testing.T) {
		initialState := map[int]map[int]Void{
			0: {1: Void{}},
		}
		grid := NewGrid(initialState)

		assert.Equal(t, "Board\n #", grid.String())
	})

	t.Run("when there are multirow cells should add newline", func(t *testing.T) {
		initialState := map[int]map[int]Void{
			0: {0: Void{}, 1: Void{}},
			1: {1: Void{}},
		}
		grid := NewGrid(initialState)

		assert.Equal(t, "Board\n##\n #", grid.String())
	})

	t.Run("when there are dead cells on right edge should be trimmed", func(t *testing.T) {
		initialState := map[int]map[int]Void{
			0: {0: Void{}},
			1: {0: Void{}, 1: Void{}, 2: Void{}},
		}
		grid := NewGrid(initialState)

		assert.Equal(t, "Board\n#\n###", grid.String())
	})
}

func TestGrid_NeighbourCount(t *testing.T) {
	t.Run("when all 8 cells around are dead returns 0", func(t *testing.T) {
		initialState := map[int]map[int]Void{
			1: {1: Void{}},
		}
		grid := NewGrid(initialState)

		assert.Equal(t, 0, grid.NeighbourCount(1, 1))
	})

	t.Run("when top left cell is alive returns 1", func(t *testing.T) {
		initialState := map[int]map[int]Void{
			0: {0: Void{}},
			1: {1: Void{}},
		}
		grid := NewGrid(initialState)

		assert.Equal(t, 1, grid.NeighbourCount(1, 1))
	})

	t.Run("when top center cell is alive returns 1", func(t *testing.T) {
		initialState := map[int]map[int]Void{
			0: {1: Void{}},
			1: {1: Void{}},
		}
		grid := NewGrid(initialState)

		assert.Equal(t, 1, grid.NeighbourCount(1, 1))
	})

	t.Run("when top right cell is alive returns 1", func(t *testing.T) {
		initialState := map[int]map[int]Void{
			0: {2: Void{}},
			1: {1: Void{}},
		}
		grid := NewGrid(initialState)

		assert.Equal(t, 1, grid.NeighbourCount(1, 1))
	})

	t.Run("when center left cell is alive returns 1", func(t *testing.T) {
		initialState := map[int]map[int]Void{
			1: {0: Void{}, 1: Void{}},
		}
		grid := NewGrid(initialState)

		assert.Equal(t, 1, grid.NeighbourCount(1, 1))
	})

	t.Run("when center right cell is alive returns 1", func(t *testing.T) {
		initialState := map[int]map[int]Void{
			1: {1: Void{}, 2: Void{}},
		}
		grid := NewGrid(initialState)

		assert.Equal(t, 1, grid.NeighbourCount(1, 1))
	})

	t.Run("when bottom left cell is alive returns 1", func(t *testing.T) {
		initialState := map[int]map[int]Void{
			1: {1: Void{}},
			2: {0: Void{}},
		}
		grid := NewGrid(initialState)

		assert.Equal(t, 1, grid.NeighbourCount(1, 1))
	})

	t.Run("when bottom center cell is alive returns 1", func(t *testing.T) {
		initialState := map[int]map[int]Void{
			1: {1: Void{}},
			2: {1: Void{}},
		}
		grid := NewGrid(initialState)

		assert.Equal(t, 1, grid.NeighbourCount(1, 1))
	})

	t.Run("when bottom right cell is alive returns 1", func(t *testing.T) {
		initialState := map[int]map[int]Void{
			1: {1: Void{}},
			2: {2: Void{}},
		}
		grid := NewGrid(initialState)

		assert.Equal(t, 1, grid.NeighbourCount(1, 1))
	})
}

func TestGrid_ExpandLeft(t *testing.T) {
	t.Run("add column at left with dead cell", func(t *testing.T) {
		state := map[int]map[int]Void{
			1: {1: Void{}},
			2: {2: Void{}},
		}
		grid := NewGrid(state)
		expandedGrid := grid.ExpandLeft()

		assert.Equal(t, grid.Width()+1, expandedGrid.Width())
		for row := 0; row < expandedGrid.Height(); row++ {
			assert.False(t, expandedGrid.IsAlive(row, 0))
		}
	})
}

func TestGrid_ExpandTop(t *testing.T) {
	t.Run("add row at top with dead cell", func(t *testing.T) {
		state := map[int]map[int]Void{
			1: {1: Void{}},
			2: {2: Void{}},
		}
		grid := NewGrid(state)
		expandedGrid := grid.ExpandTop()

		assert.Equal(t, grid.Height()+1, expandedGrid.Height())
		for column := 0; column < expandedGrid.Width(); column++ {
			assert.False(t, expandedGrid.IsAlive(0, column))
		}
	})
}

func TestGrid_ExpandBottom(t *testing.T) {
	t.Run("add row at bottom with dead cell", func(t *testing.T) {
		state := map[int]map[int]Void{
			1: {1: Void{}},
			2: {2: Void{}},
		}
		grid := NewGrid(state)
		expandedGrid := grid.ExpandBottom()

		assert.Equal(t, grid.Height()+1, expandedGrid.Height())
		for column := 0; column < expandedGrid.Width(); column++ {
			assert.False(t, expandedGrid.IsAlive(expandedGrid.Height()-1, column))
		}
	})
}
