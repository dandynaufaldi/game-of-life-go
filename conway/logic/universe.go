package logic

import "github.com/dandynaufaldi/game-of-life/go-1/conway/model"

type Universe struct{}

func NewUniverse() *Universe { return &Universe{} }

func (u *Universe) Tick(currentGrid model.Grid) model.Grid {
	return currentGrid
}
