package main

import (
	"slices"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// Simple cache maps
var cache = map[string]bool{}
var cache_part2 = map[string]int{}

// For part2 we need to count all matching variants, doing this recursivly and
// caching on the way
func isMatchingPart2(pattern string, towels []string) int {
	if cached, exists := cache_part2[pattern]; exists {
		return cached
	}

	combinations := 0
	for _, towel := range towels {
		if len(towel) > len(pattern) {
			continue
		}

		if pattern[:len(towel)] == towel {
			if len(pattern) == len(towel) {
				combinations++
			} else if len(pattern) > len(towel) {
				combinations += isMatchingPart2(pattern[len(towel):], towels)
			}
		}
	}

	cache_part2[pattern] = combinations
	return combinations
}

// For part1 we only need to find one matching pattern to be nice, still doing
// this recursivly and caching on the way
func isMatching(pattern string, towels []string) bool {
	if cached, exists := cache[pattern]; exists {
		return cached
	}

	for _, towel := range towels {
		if len(towel) > len(pattern) {
			continue
		}

		if pattern[:len(towel)] == towel {
			if len(pattern) == len(towel) || (len(pattern) > len(towel) && isMatching(pattern[len(towel):], towels)) {
				cache[pattern] = true
				return true
			}
		}
	}

	cache[pattern] = false
	return false
}

func run(part2 bool, input string) any {
	// Parse towels
	lines := strings.Split(input, "\n")
	towels := strings.Split(lines[0], ", ")
	slices.SortFunc(towels, func(a, b string) int {
		return len(b) - len(a)
	})

	matching := 0
	// Loop all patterns
	for _, pattern := range lines[2:] {
		if part2 {
			clear(cache_part2)
			matching += isMatchingPart2(pattern, towels)
		} else {
			matches := isMatching(pattern, towels)
			if matches {
				matching++
			}
		}
	}

	return matching
}
