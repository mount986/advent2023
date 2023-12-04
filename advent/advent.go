package advent

import "log"

type Advent struct {
	Days  []int
	Part  int
	Input []string
}

type Result struct {
	Part  int
	Value any
}

func (advent *Advent) Run() {
	var (
		day    Day
		runner DayRunner
		result any
		err    error
	)

	for _, dayNum := range advent.Days {
		day.Day = dayNum
		err = day.ReadInput()
		if err != nil {
			continue
		}
		runner = day.GetRunner()

		switch advent.Part {
		case 0:
			result, err = runner.Part1(day.Input)

			if err != nil {
				log.Println(err)
			}
			day.PrintResult(&Result{Part: 1, Value: result})

			result, err = runner.Part2(day.Input)
			if err != nil {
				log.Println(err)
			}
			day.PrintResult(&Result{Part: 2, Value: result})
		case 1:
			result, err = runner.Part1(day.Input)
			if err != nil {
				log.Println(err)
			}
			day.PrintResult(&Result{Part: 1, Value: result})
		case 2:
			result, err = runner.Part2(day.Input)
			if err != nil {
				log.Println(err)
			}
			day.PrintResult(&Result{Part: 2, Value: result})

		}
	}
}
