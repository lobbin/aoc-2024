package main

import (
	"fmt"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func findBest(bank []int, need int) int {
	result := 0
	start := 0

	for i := range need {
		maxDigit := -1
		maxPos := -1

		windowEnd := len(bank) - (need - i - 1)

		for j := start; j < windowEnd; j++ {
			if bank[j] > maxDigit {
				maxDigit = bank[j]
				maxPos = j
				if maxDigit == 9 {
					break
				}
			}
		}

		result = result*10 + maxDigit
		start = maxPos + 1
	}

	return result
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	banks := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		bank := make([]int, 0)
		batteries := strings.Split(line, "")
		for _, battery := range batteries {
			var num int
			fmt.Sscanf(battery, "%d", &num)
			bank = append(bank, num)
		}
		banks = append(banks, bank)
	}

	score := 0
	for _, bank := range banks {
		var best int
		if part2 {
			best = findBest(bank, 12)
		} else {
			best = findBest(bank, 2)
		}

		score += best
	}

	return score
}
