package main

func part2(lines []string) int {
	sum := 0

	for i := range lines {
		sum += evalCard(i, lines) + 1 // +1 for the card itself
	}

	return sum
}

func evalCard(currentCard int, lines []string) int {

	matches := getMatches(lines[currentCard])

	sum := 0

	for i := 1; i <= matches; i++ {
		sum += evalCard(currentCard+i, lines)
	}

	return matches + sum
}
