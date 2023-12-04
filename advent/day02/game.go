package day02

import (
	"log"
	"strconv"
	"strings"
)

type Game struct {
	Id     int
	colors map[string]int
}

func (g *Game) IsValid(colors map[string]int) bool {
	for color, amount := range g.colors {
		if amount > colors[color] {
			return false
		}
	}
	return true
}

func (g *Game) Power() int {
	return g.colors["red"] * g.colors["blue"] * g.colors["green"]
}

func ParseGame(input string) (*Game, error) {
	var (
		game *Game
	)

	game = new(Game)
	game.colors = map[string]int{}
	firstSplit := strings.Split(input, ": ")
	id, err := strconv.Atoi(firstSplit[0][strings.Index(firstSplit[0], " ")+1:])
	if err != nil {
		log.Println(err)
		return nil, err
	}
	game.Id = id

	rounds := strings.Split(firstSplit[1], "; ")

	for _, round := range rounds {
		colors := strings.Split(round, ", ")

		for _, colorEntry := range colors {
			amount, err := strconv.Atoi(colorEntry[0:strings.Index(colorEntry, " ")])
			if err != nil {
				log.Println(err)
				return nil, err
			}
			color := colorEntry[strings.Index(colorEntry, " ")+1:]
			if game.colors[color] < amount {
				game.colors[color] = amount
			}

		}
	}

	return game, nil
}
