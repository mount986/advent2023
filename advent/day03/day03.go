package day03

type Day03 struct {
}

func (a *Day03) Part1(input []string) (any, error) {
	var engine Engine

	engine.BuildEngineSchematic(input)
	engine.FindParts()
	sum := engine.SumParts()

	return sum, nil
}

func (a *Day03) Part2(input []string) (any, error) {
	var engine Engine

	engine.BuildEngineSchematic(input)
	engine.FindParts()
	sum := engine.SumGearRatios()

	return sum, nil
}
