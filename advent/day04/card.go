package day04

import (
	"math"
	"strconv"
	"strings"
)

type Card struct {
	Winning []int
	Numbers []int
	Matches []int
	Count   int
}

func NewCard(input string) *Card {
	firstSplit := strings.Split(input, ": ")
	gameData := strings.Split(firstSplit[1], " | ")

	card := Card{
		Count: 1,
	}

	for _, x := range strings.Fields(gameData[0]) {
		num, _ := strconv.Atoi(x)
		card.Winning = append(card.Winning, num)
	}
	for _, x := range strings.Fields(gameData[1]) {
		num, _ := strconv.Atoi(x)
		card.Numbers = append(card.Numbers, num)
	}

	return &card
}

func (card *Card) FindMatches() {
	for _, num := range card.Numbers {
		for _, win := range card.Winning {
			if num == win {
				card.Matches = append(card.Matches, num)
			}
		}
	}
}

func (card *Card) Value() int {
	if len(card.Matches) == 0 {
		return 0
	}

	return int(math.Pow(2, float64(len(card.Matches)-1)))
}
