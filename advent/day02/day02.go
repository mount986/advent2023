package day02

type Day02 struct {
}

func (a *Day02) Part1(input []string) (any, error) {
	var (
		allGames   []*Game
		validGames []*Game
		maxColors  map[string]int
	)
	maxColors = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, game := range input {
		g, err := ParseGame(game)
		if err != nil {
			return nil, err
		}
		allGames = append(allGames, g)
	}

	for _, game := range allGames {
		if game.IsValid(maxColors) {
			validGames = append(validGames, game)
		}
	}

	sum := 0
	for _, game := range validGames {
		sum += game.Id
	}
	return sum, nil
}

func (a *Day02) Part2(input []string) (any, error) {
	var (
		allGames []*Game
	)

	for _, game := range input {
		g, err := ParseGame(game)
		if err != nil {
			return nil, err
		}
		allGames = append(allGames, g)
	}

	sum := 0
	for _, game := range allGames {
		sum += game.Power()
	}

	return sum, nil
}
