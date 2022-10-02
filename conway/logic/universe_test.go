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
}
