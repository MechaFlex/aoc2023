package main

import (
	util "aoc2023"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func part1() int {
	//content, err := os.ReadFile("1/x1.txt")
	content, err := os.ReadFile("1/real.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	numbers := make([]int, len(lines))

	for i, line := range lines {
		numbers[i] = getNumber(line)
	}

	return util.Sum(numbers)
}

func getNumber(line string) int {

	firstDigit, lastDigit := -1, -1

	for _, l := range line {
		if !unicode.IsDigit(l) {
			continue
		}

		digit, err := strconv.Atoi(string(l))

		if err != nil {
			fmt.Println("Error:", err)
		}

		if firstDigit == -1 {
			firstDigit = digit
		}

		lastDigit = digit
	}

	return firstDigit*10 + lastDigit
}
