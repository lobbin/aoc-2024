package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// Split a stone in two values
func splitStone(stone int) (int, int) {
	str := strconv.Itoa(stone)
	mid := len(str) / 2
	left, _ := strconv.Atoi(str[:mid])
	right, _ := strconv.Atoi(str[mid:])
	return left, right
}

func run(part2 bool, input string) any {
	// Because part2 will break everything, we can't use a list but must use a
	// map instead, and keep track of how many stones of each we have
	stones := make(map[int]int)
	for _, stone := range strings.Split(input, " ") {
		number, _ := strconv.Atoi(strings.TrimSpace(stone))
		stones[number] = 1
	}

	iterations := 25
	if part2 {
		iterations = 75
	}

	for i := 0; i < iterations; i++ {
		// For each loop, we create a new map
		new_map := make(map[int]int, len(stones)*2)

		// Loop all the stones in the current map and applying its rules.
		for stone, count := range stones {
			if stone == 0 {
				new_map[1] += count
			} else if len(strconv.Itoa(stone))%2 == 0 {
				n1, n2 := splitStone(stone)
				new_map[n1] += count
				new_map[n2] += count
			} else {
				new_map[stone*2024] += count
			}
		}

		stones = new_map
	}

	// Loops all the stones to get the final count
	stone_count := 0
	for _, count := range stones {
		stone_count += count
	}

	return stone_count
}
