package logic

import "github.com/dandynaufaldi/game-of-life/go-1/conway/model"

type Universe struct{}

func NewUniverse() *Universe { return &Universe{} }

func (u *Universe) Tick(currentGrid model.Grid) model.Grid {
	state := make(map[int]map[int]model.Void)

	for row := 0; row < currentGrid.Height(); row++ {
		state[row] = make(map[int]model.Void)

		for column := 0; column < currentGrid.Height(); column++ {
			if currentGrid.IsAlive(row, column) {
				if currentGrid.NeighbourCount(row, column) < 2 {
					continue
				} else {
					state[row][column] = model.Void{}
				}
			} else {
				if currentGrid.NeighbourCount(row, column) == 3 {
					state[row][column] = model.Void{}
				}
			}
		}
	}

	return model.NewGrid(state)
}
