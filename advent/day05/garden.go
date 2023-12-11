package day05

import (
	"regexp"
	"strconv"
)

type Map struct {
	sourceStart, sourceEnd int
	difference             int
}

type Garden struct {
	collections map[string][]int
	maps        map[string][]Map
}

func (garden *Garden) ReadMaps(input []string, seedFunc func(string) []int) {
	var (
		mapName, sourceName, targetName string
	)
	garden.collections = make(map[string][]int)
	garden.maps = make(map[string][]Map)

	for index, line := range input {
		if index == 0 {
			garden.collections["seeds"] = seedFunc(line)
			continue
		}
		switch line {
		case "":
			if mapName != "" {
				garden.FindTargetValues(mapName, sourceName, targetName)
			}
			continue
		case "seed-to-soil map:":
			mapName = "s2s"
			sourceName = "seeds"
			targetName = "soil"
		case "soil-to-fertilizer map:":
			mapName = "s2f"
			sourceName = "soil"
			targetName = "fertilizer"
		case "fertilizer-to-water map:":
			mapName = "f2w"
			sourceName = "fertilizer"
			targetName = "water"
		case "water-to-light map:":
			mapName = "w2l"
			sourceName = "water"
			targetName = "lights"
		case "light-to-temperature map:":
			mapName = "l2t"
			sourceName = "lights"
			targetName = "temperature"
		case "temperature-to-humidity map:":
			mapName = "t2h"
			sourceName = "temperature"
			targetName = "humidity"
		case "humidity-to-location map:":
			mapName = "h2l"
			sourceName = "humidity"
			targetName = "location"
		default:
			garden.ReadMap(mapName, line)
		}
	}
	garden.FindTargetValues(mapName, sourceName, targetName)
}

func (garden *Garden) FindTargetValues(mapName, sourceName, targetName string) {
	found := false
	for _, sourceVal := range garden.collections[sourceName] {
		for _, m := range garden.maps[mapName] {
			if sourceVal >= m.sourceStart && sourceVal <= m.sourceEnd {
				garden.collections[targetName] = append(garden.collections[targetName], sourceVal+m.difference)
				found = true
				break
			}
		}
		if found {
			found = false
			continue
		}
		garden.collections[targetName] = append(garden.collections[targetName], sourceVal)
	}
}

func (garden *Garden) ReadMap(mapName, input string) {

	mapStrings := regexp.MustCompile(`\d+`).FindAllString(input, -1)
	sourceStart, _ := strconv.Atoi(mapStrings[1])
	targetStart, _ := strconv.Atoi(mapStrings[0])
	count, _ := strconv.Atoi(mapStrings[2])
	m := Map{
		sourceStart: sourceStart,
		sourceEnd:   sourceStart + count - 1,
		difference:  targetStart - sourceStart,
	}
	garden.maps[mapName] = append(garden.maps[mapName], m)
}

func BasicSeeds(input string) []int {
	var collection []int
	seedStrings := regexp.MustCompile(`\d+`).FindAllString(input, -1)
	for _, seed := range seedStrings {
		value, _ := strconv.Atoi(seed)
		collection = append(collection, value)
	}

	return collection
}

func ComplexSeeds(input string) []int {
	var collection []int
	seedStrings := regexp.MustCompile(`\d+`).FindAllString(input, -1)
	for i := 0; i < len(seedStrings); i += 2 {
		seedStart, _ := strconv.Atoi(seedStrings[i])
		seedRange, _ := strconv.Atoi(seedStrings[i+1])
		for value := seedStart; value < seedStart+seedRange; value++ {
			collection = append(collection, value)
		}
	}
	return collection
}
