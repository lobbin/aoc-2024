package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// Generate sublists from list by removing a single element for each sublist
func generateSublists(input []int) [][]int {
	var result [][]int
	for i := 0; i < len(input); i++ {
		sublist := append([]int{}, input[:i]...)
		sublist = append(sublist, input[i+1:]...)
		result = append(result, sublist)
	}
	return result
}

// Check whether a specific level is safe by defined constraints
func isLevelSafe(level []int) bool {
	ascending := level[1] > level[0]
	for i := 1; i < len(level); i++ {
		if !((ascending && level[i] > level[i-1] && (level[i]-level[i-1]) <= 3) ||
			(!ascending && level[i] < level[i-1] && (level[i-1]-level[i]) <= 3)) {
			return false
		}
	}

	return true
}

func run(part2 bool, input string) any {
	safe_levels := 0

	// Loop all levels
	for _, level := range strings.Split(input, "\n") {
		// Parse and convert level to integers
		numbers := strings.Split(level, " ")
		level := make([]int, len(numbers))
		for i, number_tmp := range numbers {
			number, _ := strconv.Atoi(number_tmp)
			level[i] = number
		}

		// Check is level is safe
		if isLevelSafe(level) {
			safe_levels++
		} else if part2 {
			// Part2 checks if we can make a level safe by removing a single number
			toleratedLevels := generateSublists(level)
			for _, toleratedLevel := range toleratedLevels {
				if isLevelSafe(toleratedLevel) {
					safe_levels++
					break
				}
			}
		}
	}

	return safe_levels
}
