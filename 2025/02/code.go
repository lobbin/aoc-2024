package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// My original implementation, that solved the problem. Asked Claude 4.5 if it
// could be improved for clarity and efficiency, and it produced the second code.

// func run(part2 bool, input string) any {
// 	ranges := make([][2]int, 0)
// 	for _, line := range strings.Split(input, ",") {
// 		range_parts := strings.Split(line, "-")
// 		start, _ := strconv.Atoi(range_parts[0])
// 		end, _ := strconv.Atoi(range_parts[1])
// 		ranges = append(ranges, [2]int{start, end})
// 	}

// 	score := 0
// 	for _, r := range ranges {
// 		for i := r[0]; i <= r[1]; i++ {
// 			s := strconv.Itoa(i)
// 			if part2 {
// 				mid := (len(s) + 1) / 2
// 				for j := mid; j > 0; j-- {
// 					left, _ := strconv.Atoi(s[:j])

// 					count := strings.Count(s, strconv.Itoa(left))
// 					if count >= 2 && count*j == len(s) {
// 						score += i
// 						break
// 					}
// 				}
// 			} else {
// 				if len(s)%2 == 0 {
// 					mid := len(s) / 2
// 					left, _ := strconv.Atoi(s[:mid])
// 					right, _ := strconv.Atoi(s[mid:])

// 					if left == right {
// 						score += i
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return score
// }

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	ranges := make([][2]int, 0)
	for _, line := range strings.Split(input, ",") {
		range_parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(range_parts[0])
		end, _ := strconv.Atoi(range_parts[1])
		ranges = append(ranges, [2]int{start, end})
	}

	score := 0
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			if part2 {
				if isRepeatedPattern(i) {
					score += i
				}
			} else {
				if isDoubledPattern(i) {
					score += i
				}
			}
		}
	}

	return score
}

// isDoubledPattern checks if a number is a pattern repeated exactly twice (e.g., 11, 1010, 123123)
func isDoubledPattern(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	// Must have even length
	if length%2 != 0 {
		return false
	}

	mid := length / 2
	return s[:mid] == s[mid:]
}

// isRepeatedPattern checks if a number is made of a pattern repeated 2+ times
func isRepeatedPattern(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	// Try all possible pattern lengths (divisors of total length)
	for patternLen := 1; patternLen <= length/2; patternLen++ {
		if length%patternLen == 0 {
			pattern := s[:patternLen]
			isMatch := true
			for pos := patternLen; pos < length; pos += patternLen {
				if s[pos:pos+patternLen] != pattern {
					isMatch = false
					break
				}
			}
			if isMatch {
				return true
			}
		}
	}
	return false
}
