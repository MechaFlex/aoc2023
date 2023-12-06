package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// content, err := os.ReadFile("4/x1.txt")
	content, err := os.ReadFile("4/real.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("File read.")

	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	fmt.Println(part1(lines))
	t := time.Now()
	fmt.Println(part2(lines))
	fmt.Println("Part 2 took", time.Since(t))
}
