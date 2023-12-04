package main

import (
	util "aoc2023"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var limits map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func part1() int {
	// content, err := os.ReadFile("2/x1.txt")
	content, err := os.ReadFile("2/real.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	validLines := make([]int, len(lines))

	for i, line := range lines {
		line = strings.Split(line, ":")[1]

		if isValid(line) {
			validLines[i] = i + 1
		}
	}

	return util.Sum(validLines)
}

func isValid(line string) bool {
	sets := strings.Split(line, ";")

	for _, set := range sets {
		cubesInSet := strings.Split(set, ",")

		for _, cubeInSet := range cubesInSet {
			cubeInSet = strings.TrimSpace(cubeInSet)
			splitCubeInSet := strings.Split(cubeInSet, " ")

			cubeCount, err := strconv.Atoi(splitCubeInSet[0])
			cubeColor := splitCubeInSet[1]

			if err != nil {
				fmt.Println("Error:", err)
			}

			if limits[cubeColor] < cubeCount {
				return false
			}
		}
	}

	return true
}
