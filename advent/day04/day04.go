package day04

type Day04 struct {
}

func (a *Day04) Part1(input []string) (any, error) {
	sum := 0
	for _, line := range input {
		card := NewCard(line)
		card.FindMatches()
		sum += card.Value()
	}

	return sum, nil
}

func (a *Day04) Part2(input []string) (any, error) {
	var (
		cardList []*Card
	)
	totalCards := 0

	for _, line := range input {
		card := NewCard(line)
		card.FindMatches()
		cardList = append(cardList, card)
	}

	for index, card := range cardList {
		for i := 0; i < card.Count; i++ {
			for j := 1; j <= len(card.Matches); j++ {
				cardList[index+j].Count++
			}
		}
		totalCards += card.Count
	}

	return totalCards, nil
}
