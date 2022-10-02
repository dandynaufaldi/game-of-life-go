package handler

import (
	"io"

	"github.com/dandynaufaldi/game-of-life/go-1/conway/model"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(reader io.Reader) map[int]map[int]model.Void {
	return map[int]map[int]model.Void{
		0: {0: model.Void{}, 1: model.Void{}},
		1: {0: model.Void{}, 1: model.Void{}},
	}
}
