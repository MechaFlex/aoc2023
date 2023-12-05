package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// content, err := os.ReadFile("3/x1.txt")
	content, err := os.ReadFile("3/real.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")

	numbers := []Number{}

	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
		numbers = append(numbers, findNumbers(i, lines)...)
	}

	fmt.Println("Sum for part 1:", part1(numbers, lines))
	fmt.Println("Sum for part 2:", part2(numbers, lines))
}

type Number struct {
	value int
	row   int
	start int
	end   int
}

func findNumbers(lineIndex int, lines []string) []Number {

	foundNumbers := []Number{}
	workingNumber := Number{0, lineIndex, -1, len(lines[lineIndex]) - 1}

	for i, char := range lines[lineIndex] {

		if unicode.IsDigit(char) {
			if workingNumber.start == -1 {
				workingNumber.start = i
			}
			workingNumber.value = workingNumber.value*10 + int(char-'0')
		} else {
			if workingNumber.start != -1 {
				workingNumber.end = i - 1
				foundNumbers = append(foundNumbers, workingNumber)
				workingNumber = Number{0, lineIndex, -1, 0}
			}
		}
	}

	if workingNumber.start != -1 {
		workingNumber.end = len(lines[lineIndex]) - 1
		foundNumbers = append(foundNumbers, workingNumber)
	}

	return foundNumbers
}

func sumNumbers(numbers []Number) int {

	sum := 0

	for _, number := range numbers {
		sum += number.value
	}

	return sum
}
