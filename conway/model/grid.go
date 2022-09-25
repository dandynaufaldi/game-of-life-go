package model

type Grid struct {

}

func NewGrid(initialState [][]bool) Grid {
	return Grid{}
}

func (g Grid) Height() int {
	return 2
}
