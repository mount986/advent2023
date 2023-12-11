package day05

import (
	"slices"
)

type Day05 struct {
	garden Garden
}

func (a *Day05) Part1(input []string) (any, error) {
	a.garden.ReadMaps(input, BasicSeeds)

	return slices.Min(a.garden.collections["location"]), nil
}

func (a *Day05) Part2(input []string) (any, error) {
	a.garden.ReadMaps(input, ComplexSeeds)

	return slices.Min(a.garden.collections["location"]), nil
}
