package day03

import "unicode"

type Engine struct {
	parts []int
	grid  []string
}

func (engine *Engine) BuildGrid(input []string) {
	engine.grid = []string{}
	for _, line := range input {
		engine.grid = append(engine.grid, line)
	}
}

func (engine *Engine) FindParts() {
	engine.parts = []int{}

	for _, row := range engine.grid {
		for i := 0; i < len(row); i++ {
			r := row[i]
			if unicode.IsDigit(r) {

			}
		}
	}
}
