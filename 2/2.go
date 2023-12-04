package main

import (
	util "aoc2023"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part2() int {
	// content, err := os.ReadFile("2/x1.txt")
	content, err := os.ReadFile("2/real.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	powers := make([]int, len(lines))

	for i, line := range lines {
		line = strings.Split(line, ":")[1]

		powers[i] = getPower(line)
	}

	return util.Sum(powers)
}

func getPower(line string) int {
	minCounts := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

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

			if minCounts[cubeColor] < cubeCount {
				minCounts[cubeColor] = cubeCount
			}
		}
	}

	return minCounts["red"] * minCounts["green"] * minCounts["blue"]
}
