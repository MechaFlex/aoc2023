package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part1(lines []string) int {

	sum := 0.0

	for _, line := range lines {

		matches := getMatches(line)

		if matches != 0 {
			sum += math.Pow(2, float64(matches)-1)
		}
	}

	return int(sum)
}

func getMatches(line string) int {
	line = strings.Split(line, ":")[1]

	unparsedNumbers := strings.Split(line, "|")
	winning := getNumbers(unparsedNumbers[0])
	having := getNumbers(unparsedNumbers[1])

	matches := 0

	for _, win := range winning {
		for _, have := range having {
			if win == have {
				matches += 1
			}
		}
	}

	return matches
}

func getNumbers(unparsedNumbers string) []int {

	unparsedNumbers = strings.TrimSpace(unparsedNumbers)
	unparsedNumbers = strings.ReplaceAll(unparsedNumbers, "  ", " ")
	numbers := strings.Split(unparsedNumbers, " ")
	ints := []int{}

	for _, number := range numbers {
		i, err := strconv.Atoi(number)

		if err != nil {
			fmt.Println("Error:", err)
		}

		ints = append(ints, i)
	}

	return ints
}
