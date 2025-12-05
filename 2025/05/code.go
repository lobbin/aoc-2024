package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	ingredients := false
	fresh := 0

	ranges := make([][2]int, 0)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			ingredients = true
			if part2 {
				break
			} else {
				continue
			}
		}

		if !ingredients {
			numbers := strings.Split(line, "-")
			n0, _ := strconv.Atoi(numbers[0])
			n1, _ := strconv.Atoi(numbers[1])
			ranges = append(ranges, [2]int{n0, n1})
		} else {
			id, _ := strconv.Atoi(line)
			if isFresh(id, ranges) {
				fresh++
			}
		}
	}

	if part2 {
		ranges = mergeRanges(ranges)
		for _, r := range ranges {
			fresh += r[1] - r[0] + 1
		}
	}

	return fresh
}

func mergeRanges(ranges [][2]int) [][2]int {
	slices.SortFunc(ranges, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	merged := make([][2]int, 0, len(ranges))
	merged = append(merged, ranges[0])

	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]

		if r[0] <= last[1]+1 {
			last[1] = max(last[1], r[1])
		} else {
			merged = append(merged, r)
		}
	}

	return merged
}

func isFresh(id int, ranges [][2]int) bool {
	for _, r := range ranges {
		if id >= r[0] && id <= r[1] {
			return true
		}
	}
	return false
}
