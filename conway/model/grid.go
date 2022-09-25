package model

type Grid struct {
	state  map[int]map[int]Void
	height int
}

type Void struct{}

func NewGrid(initialState map[int]map[int]Void) Grid {
	height := 0
	for row := range initialState {
		if row+1 > height {
			height = row + 1
		}
	}

	return Grid{
		state:  initialState,
		height: height,
	}
}

func (g Grid) Height() int {
	return g.height
}

func (g Grid) Width() int {
	return 2
}
