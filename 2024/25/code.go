package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	locks := [][5]int{}
	keys := [][5]int{}

	// Load keys and locks
	lines := strings.Split(input, "\n")
	for l := 0; l < len(lines); l += 8 {
		// For each, count the number of # between the first and last line.
		data := [5]int{}
		for y := 1; y < 6; y++ {
			for x := 0; x < 5; x++ {
				if lines[l+y][x] == '#' {
					data[x]++
				}
			}
		}

		// This prefix tells us this is a lock, otherwise it's a key
		lock := strings.HasPrefix(lines[l], "#####")
		if lock {
			locks = append(locks, data)
		} else {
			keys = append(keys, data)
		}
	}

	fits := 0
	// Loop the keys with all the locks to see if they could be a possible match
	for _, key := range keys {
	lock_loop:
		for _, lock := range locks {
			for i := 0; i < 5; i++ {
				if key[i]+lock[i] > 5 {
					continue lock_loop
				}
			}
			fits++
		}
	}

	return fits
}
