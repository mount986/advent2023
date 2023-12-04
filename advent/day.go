package advent

import (
	"bufio"
	"fmt"
	"github.com/mount986/advent2023/advent/day02"
	"github.com/mount986/advent2023/advent/day03"
	"github.com/mount986/advent2023/advent/day04"
	"github.com/mount986/advent2023/advent/day05"
	"github.com/mount986/advent2023/advent/day06"
	"github.com/mount986/advent2023/advent/day07"
	"github.com/mount986/advent2023/advent/day08"
	"github.com/mount986/advent2023/advent/day09"
	"github.com/mount986/advent2023/advent/day10"
	"github.com/mount986/advent2023/advent/day11"
	"github.com/mount986/advent2023/advent/day12"
	"github.com/mount986/advent2023/advent/day13"
	"github.com/mount986/advent2023/advent/day14"
	"github.com/mount986/advent2023/advent/day15"
	"github.com/mount986/advent2023/advent/day16"
	"github.com/mount986/advent2023/advent/day17"
	"github.com/mount986/advent2023/advent/day18"
	"github.com/mount986/advent2023/advent/day19"
	"github.com/mount986/advent2023/advent/day20"
	"github.com/mount986/advent2023/advent/day21"
	"github.com/mount986/advent2023/advent/day22"
	"github.com/mount986/advent2023/advent/day23"
	"github.com/mount986/advent2023/advent/day24"
	"log"
	"os"

	"github.com/mount986/advent2023/advent/day01"
)

var dayTypes map[int]any

type DayRunner interface {
	Part1(input []string) (any, error)
	Part2(input []string) (any, error)
}

func init() {
	dayTypes = map[int]any{}
	dayTypes[1] = &day01.Day01{}
	dayTypes[2] = &day02.Day02{}
	dayTypes[3] = &day03.Day03{}
	dayTypes[4] = &day04.Day04{}
	dayTypes[5] = &day05.Day05{}
	dayTypes[6] = &day06.Day06{}
	dayTypes[7] = &day07.Day07{}
	dayTypes[8] = &day08.Day08{}
	dayTypes[9] = &day09.Day09{}
	dayTypes[10] = &day10.Day10{}
	dayTypes[11] = &day11.Day11{}
	dayTypes[12] = &day12.Day12{}
	dayTypes[13] = &day13.Day13{}
	dayTypes[14] = &day14.Day14{}
	dayTypes[15] = &day15.Day15{}
	dayTypes[16] = &day16.Day16{}
	dayTypes[17] = &day17.Day17{}
	dayTypes[18] = &day18.Day18{}
	dayTypes[19] = &day19.Day19{}
	dayTypes[20] = &day20.Day20{}
	dayTypes[21] = &day21.Day21{}
	dayTypes[22] = &day22.Day22{}
	dayTypes[23] = &day23.Day23{}
	dayTypes[24] = &day24.Day24{}
}

type Day struct {
	Day   int
	Input []string
}

func (day *Day) GetRunner() DayRunner {
	return dayTypes[day.Day].(DayRunner)
}

func (day *Day) ReadInput() error {
	var (
		scanner *bufio.Scanner
		file    *os.File
		err     error
	)
	file, err = os.Open(fmt.Sprintf("advent/day%02d/input", day.Day))
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)

	day.Input = []string{}
	for scanner.Scan() {
		day.Input = append(day.Input, scanner.Text())
	}
	return nil
}

func (day *Day) PrintResult(result *Result) {
	log.Println(fmt.Sprintf("The solution for Day %d Part %d is: %v", day.Day, result.Part, result.Value))
}
