package day05

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	var (
		file    *os.File
		input   []string
		scanner *bufio.Scanner
		err     error
	)
	file, err = os.Open("test.dat")
	if err != nil {
		t.Fatal("Failed to open test data file")
	}
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	expected := 35
	day05 := Day05{}
	got, err := day05.Part1(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, got.(int))
}

func TestPart2(t *testing.T) {
	var (
		file    *os.File
		input   []string
		scanner *bufio.Scanner
		err     error
	)
	file, err = os.Open("test.dat")
	if err != nil {
		t.Fatal("Failed to open test data file")
	}
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	expected := 46
	day05 := Day05{}
	got, err := day05.Part2(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, got.(int))
}
