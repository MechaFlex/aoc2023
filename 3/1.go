package main

import (
	"unicode"
)

func part1(numbers []Number, lines []string) int {

	numbersWithSymbol := findNumbersWithSymbol(numbers, lines)

	sum := sumNumbers(numbersWithSymbol)

	return sum
}

func findNumbersWithSymbol(numbers []Number, lines []string) []Number {

	numbersWithSymbol := []Number{}

	for _, number := range numbers {
		foundSymbol := false

		touchesLeftEdge := number.start == 0
		touchesRightEdge := number.end == len(lines[number.row])-1

		if !touchesLeftEdge {
			if isSymbol(rune(lines[number.row][number.start-1])) {
				foundSymbol = true
			}
		}

		if !touchesRightEdge {
			if isSymbol(rune(lines[number.row][number.end+1])) {
				foundSymbol = true
			}
		}

		leftOffset := 0
		if !touchesLeftEdge {
			leftOffset = 1
		}
		rightOffset := 1
		if !touchesRightEdge {
			rightOffset = 2
		}

		if number.row > 0 {
			for _, char := range lines[number.row-1][number.start-leftOffset : number.end+rightOffset] {
				if isSymbol(char) {
					foundSymbol = true
				}
			}
		}

		if number.row < len(lines)-1 {
			for _, char := range lines[number.row+1][number.start-leftOffset : number.end+rightOffset] {
				if isSymbol(char) {
					foundSymbol = true
				}
			}
		}

		if foundSymbol {
			numbersWithSymbol = append(numbersWithSymbol, number)
		}
	}

	return numbersWithSymbol
}

func isSymbol(char rune) bool {
	return !(char == '.' || unicode.IsDigit(char))
}
