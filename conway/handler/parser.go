package handler

import (
	"bufio"
	"io"

	"github.com/dandynaufaldi/game-of-life/go-1/conway/model"
)

const aliveCell = "#"

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(reader io.Reader) map[int]map[int]model.Void {
	scanner := bufio.NewScanner(reader)
	row := 0
	state := make(map[int]map[int]model.Void)

	for scanner.Scan() {
		state[row] = make(map[int]model.Void)
		for i, char := range scanner.Text() {
			if string(char) == aliveCell {
				state[row][i] = model.Void{}
			}
		}
		row++
	}

	return state
}
