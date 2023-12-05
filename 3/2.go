package main

type Coord struct {
	x         int
	y         int
	isGear    bool
	gearRatio int
}

func part2(numbers []Number, lines []string) int {
	starCoords := []Coord{}

	for i, line := range lines {
		for j, char := range line {
			if char == '*' {
				starCoords = append(starCoords, Coord{j, i, false, 0})
			}
		}
	}

	for i, starCoord := range starCoords {
		adjacentNumbers := []Number{}

		for _, number := range numbers {
			// Left or right
			if number.row == starCoord.y && (number.end == starCoord.x-1 || number.start == starCoord.x+1) {
				adjacentNumbers = append(adjacentNumbers, number)
				continue
			}
			// Above or below
			if number.row == starCoord.y-1 || number.row == starCoord.y+1 {
				if number.start <= starCoord.x+1 && number.end >= starCoord.x-1 {
					adjacentNumbers = append(adjacentNumbers, number)
					continue
				}
			}
		}

		if len(adjacentNumbers) == 2 {
			starCoords[i].isGear = true
			starCoords[i].gearRatio = adjacentNumbers[0].value * adjacentNumbers[1].value
		}
	}

	return sumGearRatios(starCoords)
}

func sumGearRatios(starCoords []Coord) int {
	sum := 0

	for _, starCoord := range starCoords {
		if starCoord.isGear {
			sum += starCoord.gearRatio
		}
	}

	return sum
}
