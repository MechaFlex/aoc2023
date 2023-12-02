package main

import (
	"fmt"
	"os"
	"strings"
)

var numberStrings = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func part2() int {
	//content, err := os.ReadFile("1/x2.txt")
	content, err := os.ReadFile("1/real.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	numbers := make([]int, len(lines))

	for i, line := range lines {
		numbers[i] = withStringsGetNumber(line)
	}

	return sum(numbers)
}

func withStringsGetNumber(line string) int {

	indexOfFirstDigit := strings.IndexAny(line, "0123456789")
	indexOfLastDigit := strings.LastIndexAny(line, "0123456789")

	firstDigit := -1
	lastDigit := -1

	stringBeforeFirstDigit := line
	stringAfterLastDigit := line

	if indexOfFirstDigit != -1 {
		firstDigit = int(line[indexOfFirstDigit] - '0')
		stringBeforeFirstDigit = line[:indexOfFirstDigit]
	}

	if indexOfLastDigit != -1 {
		lastDigit = int(line[indexOfLastDigit] - '0')
		stringAfterLastDigit = line[indexOfLastDigit+1:]
	}

findFirstTextDigit:
	for i := range stringBeforeFirstDigit {
		for j, numberString := range numberStrings {
			if strings.Contains(stringBeforeFirstDigit[:i+1], numberString) {
				firstDigit = j
				break findFirstTextDigit
			}
		}
	}

findLastTextDigit:
	for i := range stringAfterLastDigit {
		for j, numberString := range numberStrings {
			if strings.Contains(stringAfterLastDigit[len(stringAfterLastDigit)-i-1:], numberString) {
				lastDigit = j
				break findLastTextDigit
			}
		}
	}

	return firstDigit*10 + lastDigit
}
