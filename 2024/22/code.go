package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func mix(secret, mix int) int {
	return secret ^ mix
}

func prune(secret int) int {
	return secret % 16777216
}

func main() {
	aoc.Harness(run)
}

type IntMap [4]int

func run(part2 bool, input string) any {
	lines := strings.Split(input, "\n")
	sequenceMap := make([]map[IntMap]int, len(lines))

	sum := 0
	for l, line := range lines {
		secret, _ := strconv.Atoi(line)
		last_value := secret % 10
		diffs := []int{last_value}

		value_map := map[IntMap]int{}

		// Main loop
		for i := 0; i < 2000; i++ {
			secret = prune(mix(secret, secret*64))
			secret = prune(mix(secret, int(math.Floor(float64(secret)/32))))
			secret = prune(mix(secret, secret*2048))

			if part2 {
				last := secret % 10

				diffs = append(diffs, last-last_value)
				last_value = last

				// We're only interested in positive buys
				if len(diffs) >= 4 && last > 0 {
					key := [4]int{diffs[len(diffs)-4], diffs[len(diffs)-3], diffs[len(diffs)-2], diffs[len(diffs)-1]}
					if _, exists := value_map[key]; !exists {
						value_map[key] = last
					}
				}
			}
		}

		if part2 {
			sequenceMap[l] = value_map
		} else {
			sum += secret
		}
	}

	if part2 {
		best := math.MinInt
		cache := map[IntMap]bool{}

		// Loop all sequences
		for _, sequences := range sequenceMap {
			for k := range sequences {
				// Check whether we looked at this before
				if cache[k] {
					continue
				}
				cache[k] = true

				local_best := 0

				// Loop all sequences once again to find buyers that shares the same
				// sequence and add their value
				for _, sequences_round2 := range sequenceMap {
					if value, exists := sequences_round2[k]; exists {
						local_best += value
					}
				}

				if local_best > best {
					best = local_best
				}
			}
		}

		return best
	}

	return sum
}
