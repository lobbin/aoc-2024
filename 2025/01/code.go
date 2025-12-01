package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

type Pair struct {
	a, b interface{}
}

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
	list := make([]Pair, 0)
	for _, line := range strings.Split(input, "\n") {
		dir := string(line[0])
		value, _ := strconv.Atoi(line[1:])
		list = append(list, Pair{dir, value})
	}

	start := 50
	zeroes := 0
	for _, dial := range list {
		before := start
		steps := dial.b.(int)
		if part2 {
			if steps > 100 {
				z := steps / 100
				steps -= z * 100
				zeroes += z
			}
		}

		if dial.a == "L" {
			start -= steps
		} else {
			start += steps
		}

		if part2 {
			if start > 100 || (before > 0 && start < 0) {
				zeroes++
			}
		}

		start %= 100
		if start < 0 {
			start += 100
		}

		if start == 0 {
			zeroes++
		}
	}

	return zeroes
}
