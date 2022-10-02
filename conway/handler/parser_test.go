package handler

import (
	"os"
	"testing"

	"github.com/dandynaufaldi/game-of-life/go-1/conway/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser_Parse(t *testing.T) {
	t.Run("when given 2x2 input pattern returns respective pattern", func(t *testing.T) {
		input, err := os.Open("testdata/2x2.txt")
		require.NoError(t, err)
		defer input.Close()

		parser := NewParser()
		expectedState := map[int]map[int]model.Void{
			0: {0: model.Void{}, 1: model.Void{}},
			1: {0: model.Void{}, 1: model.Void{}},
		}
		
		assert.Equal(t, expectedState, parser.Parse(input))
	})
}