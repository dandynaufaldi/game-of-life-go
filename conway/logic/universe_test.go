package logic

import (
	"testing"

	"github.com/dandynaufaldi/game-of-life/go-1/conway/model"
	"github.com/stretchr/testify/assert"
)

func TestUniverse_Tick(t *testing.T) {
	t.Run("when given grid with 2x2 pattern returns grid with same 2x2 pattern", func(t *testing.T) {
		twoByTwo := map[int]map[int]model.Void{
			0: {0: model.Void{}, 1: model.Void{}},
			1: {0: model.Void{}, 1: model.Void{}},
		}
		grid := model.NewGrid(twoByTwo)
		universe := NewUniverse()
		expectedPattern := "Board\n##\n##"

		assert.Equal(t, expectedPattern, universe.Tick(grid).String())
	})

	t.Run("when alive cell has less than 2 neighbor should be killed", func(t *testing.T) {
		state := map[int]map[int]model.Void{
			0: {2: model.Void{}},
			2: {0: model.Void{}},
		}
		grid := model.NewGrid(state)
		universe := NewUniverse()
		expectedPattern := "Board\n\n\n"

		assert.Equal(t, expectedPattern, universe.Tick(grid).String())
	})

	t.Run("when dead cell has 3 neighbors should be revived", func(t *testing.T) {
		state := map[int]map[int]model.Void{
			0: {1: model.Void{}},
			1: {1: model.Void{}},
			2: {1: model.Void{}},
		}
		grid := model.NewGrid(state)
		universe := NewUniverse()
		expectedPattern := "Board\n\n###\n"

		assert.Equal(t, expectedPattern, universe.Tick(grid).String())
	})

	t.Run("when alive cell have more than 3 neighbors should be killed", func(t *testing.T) {
		state := map[int]map[int]model.Void{
			0: {0: model.Void{}, 2: model.Void{}},
			1: {1: model.Void{}},
			2: {0: model.Void{}, 2: model.Void{}},
		}
		grid := model.NewGrid(state)
		universe := NewUniverse()
		expectedPattern := "Board\n #\n# #\n #"

		assert.Equal(t, expectedPattern, universe.Tick(grid).String())
	})

	t.Run("when there are 3 alive cells on left edge should expand left edge", func(t *testing.T) {
		state := map[int]map[int]model.Void{
			0: {0: model.Void{}},
			1: {0: model.Void{}},
			2: {0: model.Void{}},
		}
		grid := model.NewGrid(state)
		universe := NewUniverse()
		expectedPattern := "Board\n\n###\n"

		assert.Equal(t, expectedPattern, universe.Tick(grid).String())
	})

	t.Run("when there are 3 alive cells on right edge should expand right edge", func(t *testing.T) {
		state := map[int]map[int]model.Void{
			0: {1: model.Void{}},
			1: {1: model.Void{}},
			2: {1: model.Void{}},
		}
		grid := model.NewGrid(state)
		universe := NewUniverse()
		expectedPattern := "Board\n\n###\n"

		assert.Equal(t, expectedPattern, universe.Tick(grid).String())
	})
}
