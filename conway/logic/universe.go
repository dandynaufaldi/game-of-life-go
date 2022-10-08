package logic

import (
	"github.com/dandynaufaldi/game-of-life/go-1/conway/model"
)

type Universe struct{}

func NewUniverse() *Universe { return &Universe{} }

func (u *Universe) Tick(currentGrid model.Grid) model.Grid {
	currentGrid = u.expandGrid(currentGrid)
	state := make(map[int]map[int]model.Void)

	for row := 0; row < currentGrid.Height(); row++ {
		state[row] = make(map[int]model.Void)

		for column := 0; column < currentGrid.Width()+1; column++ {
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

	neighbour := grid.NeighbourCount(row, column)
	return neighbour >= 2 && neighbour <= 3
}

func (u *Universe) shouldCellBeRevived(grid model.Grid, row, column int) bool {
	if grid.IsAlive(row, column) {
		return false
	}

	return grid.NeighbourCount(row, column) == 3
}

func (u *Universe) expandGrid(grid model.Grid) model.Grid {
	newGrid := grid

	if u.shouldExpandLeft(grid) {
		newGrid = newGrid.ExpandLeft()
	}

	if u.shouldExpandBottom(grid) {
		newGrid = newGrid.ExpandBottom()
	}

	return newGrid
}

func (u *Universe) shouldExpandLeft(grid model.Grid) bool {
	aliveCounter := 0
	for row := 0; row < grid.Height(); row++ {
		if grid.IsAlive(row, 0) {
			aliveCounter++
		} else {
			aliveCounter = 0
		}

		if aliveCounter == 3 {
			return true
		}
	}

	return false
}

func (u *Universe) shouldExpandBottom(grid model.Grid) bool {
	aliveCounter := 0
	for column := 0; column < grid.Width(); column++ {
		if grid.IsAlive(grid.Height()-1, column) {
			aliveCounter++
		} else {
			aliveCounter = 0
		}

		if aliveCounter == 3 {
			return true
		}
	}

	return false
}
