package day03

import (
	"regexp"
	"strconv"
)

type Symbol struct {
	symbol   string
	row, col int
	parts    []*Part
}

type Part struct {
	row   int
	col   *Column
	value int
}

type Column struct {
	start, end int
}

type Engine struct {
	parts   []*Part
	symbols []*Symbol
}

func (engine *Engine) BuildEngineSchematic(input []string) {
	engine.parts = []*Part{}
	engine.symbols = []*Symbol{}

	partRegex := regexp.MustCompile(`\d+`)
	symbolRegex := regexp.MustCompile(`[^\.\d]`)

	for row, line := range input {
		for _, part := range partRegex.FindAllStringIndex(line, -1) {
			value, _ := strconv.Atoi(line[part[0]:part[1]])
			engine.parts = append(engine.parts, &Part{
				row: row,
				col: &Column{
					start: part[0],
					end:   part[1],
				},
				value: value,
			})
		}

		for _, symbol := range symbolRegex.FindAllStringIndex(line, -1) {
			engine.symbols = append(engine.symbols, &Symbol{
				symbol: line[symbol[0]:symbol[1]],
				row:    row,
				col:    symbol[0],
				parts:  []*Part{},
			})
		}
	}
}

func (engine *Engine) FindParts() {
	for _, symbol := range engine.symbols {
		for _, part := range engine.parts {
			CheckIfTouching(symbol, part)
		}
	}
}

func CheckIfTouching(symbol *Symbol, part *Part) {
	if part.row >= symbol.row-1 && part.row <= symbol.row+1 &&
		part.col.end-1 >= symbol.col-1 && part.col.start <= symbol.col+1 {
		symbol.parts = append(symbol.parts, part)
	}
}

func (engine *Engine) SumParts() int {
	sum := 0

	for _, symbol := range engine.symbols {
		for _, part := range symbol.parts {
			sum += part.value
		}
	}

	return sum
}

func (engine *Engine) SumGearRatios() int {
	sum := 0

	for _, symbol := range engine.symbols {
		if symbol.symbol == "*" && len(symbol.parts) == 2 {
			sum += symbol.parts[0].value * symbol.parts[1].value
		}
	}

	return sum
}
