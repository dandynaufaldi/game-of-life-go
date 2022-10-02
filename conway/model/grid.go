package model

import "strings"

type Grid struct {
	state  map[int]map[int]Void
	height int
	width  int
}

type Void struct{}

func NewGrid(initialState map[int]map[int]Void) Grid {
	height := 0
	width := 0
	for row, columns := range initialState {
		if row+1 > height {
			height = row + 1
		}

		for column := range columns {
			if column+1 > width {
				width = column + 1
			}
		}
	}

	return Grid{
		state:  initialState,
		height: height,
		width:  width,
	}
}

func (g Grid) Height() int {
	return g.height
}

func (g Grid) Width() int {
	return g.width
}

func (g Grid) IsAlive(row, column int) bool {
	if columns, ok := g.state[row]; ok {
		if _, ok := columns[column]; ok {
			return true
		}
	}

	return false
}

func (g Grid) String() string {
	var sb strings.Builder
	sb.WriteString("Board\n")
	for row := 0; row < g.Height(); row++ {
		sb.WriteString(g.mapRowToString(row))
		sb.WriteString("\n")
	}

	return strings.TrimSuffix(sb.String(), "\n")
}

func (g Grid) NeighbourCount(row, column int) int {
	neighbourCount := 0
	directions := []struct {
		row, column int
	}{
		{row: -1, column: -1},
		{row: -1, column: 0},
		{row: -1, column: +1},
		{row: 0, column: -1},
		{row: 0, column: +1},
		{row: +1, column: -1},
		{row: +1, column: 0},
		{row: +1, column: +1},
	}

	for _, direction := range directions {
		if g.IsAlive(row+direction.row, column+direction.column) {
			neighbourCount += 1
		}
	}

	return neighbourCount
}

func (g Grid) ExpandLeft() Grid {
	newState := make(map[int]map[int]Void)
	for row, columns := range g.state {
		newState[row] = make(map[int]Void)

		for column := range columns {
			newState[row][column+1] = Void{}
		}
	}

	return NewGrid(newState)
}

func (g Grid) mapRowToString(row int) string {
	var sb strings.Builder
	for column := 0; column < g.Width(); column++ {
		if g.IsAlive(row, column) {
			sb.WriteString("#")
		} else {
			sb.WriteString(" ")
		}
	}

	return strings.TrimRight(sb.String(), " ")
}
