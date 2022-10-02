package logic

import "github.com/dandynaufaldi/game-of-life/go-1/conway/model"

type Universe struct{}

func NewUniverse() *Universe { return &Universe{} }

func (u *Universe) Tick(currentGrid model.Grid) model.Grid {
	state := make(map[int]map[int]model.Void)

	for row := 0; row < currentGrid.Height(); row++ {
		state[row] = make(map[int]model.Void)

		for column := 0; column < currentGrid.Height(); column++ {
			if u.shouldCellStayAlive(currentGrid, row, column) ||
				u.shouldCellBeRevived(currentGrid, row, column) {
				state[row][column] = model.Void{}
			}
		}
	}

	return model.NewGrid(state)
}

func (u *Universe) shouldCellStayAlive(grid model.Grid, row, column int) bool {
	if !grid.IsAlive(row, column) {
		return false
	}

	return grid.NeighbourCount(row, column) >= 2
}

func (u *Universe) shouldCellBeRevived(grid model.Grid, row, column int) bool {
	if grid.IsAlive(row, column) {
		return false
	}

	return grid.NeighbourCount(row, column) == 3
}
