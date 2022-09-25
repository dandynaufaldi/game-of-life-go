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
	for row := 0; row < g.Height(); row++ {
		sb.WriteString(g.mapRowToString(row))
		sb.WriteString("\n")
	}

	return strings.TrimRight(sb.String(), "\n")
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
