package day01

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Day01 struct {
}

func (a *Day01) Part1(input []string) (any, error) {
	sum := 0

	for _, line := range input {
		num, err := GetCalibration(line)
		if err != nil {
			return nil, err
		}
		sum += num
	}
	return sum, nil
}

func GetCalibration(line string) (int, error) {
	firstIndex := strings.IndexAny(line, "0123456789")
	lastIndex := strings.LastIndexAny(line, "0123456789")
	numString := fmt.Sprintf("%s%s", string(line[firstIndex]), string(line[lastIndex]))
	num, err := strconv.Atoi(numString)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return num, nil
}

func (a *Day01) Part2(input []string) (any, error) {
	sum := 0

	for _, line := range input {
		num, err := GetTrickyCalibration(line)
		if err != nil {
			return nil, err
		}
		sum += num
	}
	return sum, nil
}

func GetTrickyCalibration(line string) (int, error) {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero", `\d`}
	var resultIndexes [][]int

	for _, number := range numbers {
		reg := regexp.MustCompile(number)
		resultIndexes = append(resultIndexes, reg.FindAllStringIndex(line, -1)...)
	}

	sort.Slice(resultIndexes, func(a, b int) bool {
		return resultIndexes[a][0] < resultIndexes[b][0]
	})

	firstResult := resultIndexes[0]
	lastResult := resultIndexes[len(resultIndexes)-1]

	firstNum := parse(line[firstResult[0]:firstResult[1]])
	lastNum := parse(line[lastResult[0]:lastResult[1]])

	numString := fmt.Sprintf("%d%d", firstNum, lastNum)
	num, err := strconv.Atoi(numString)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return num, nil
}

func parse(num string) int {
	conv := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"0":     0,
	}

	return conv[num]
}
