package main

import (
	"regexp"
	"strconv"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// Convert to in a return multiplication result
func calc(as, bs string) int {
	a, _ := strconv.Atoi(as)
	b, _ := strconv.Atoi(bs)
	return a * b
}

func run(part2 bool, input string) any {
	// Compile regexs
	regex_mul := regexp.MustCompile(`(mul\((\d+),(\d+)\))`)
	regex_do := regexp.MustCompile(`do\(`)
	regex_dont := regexp.MustCompile(`don't\(`)

	result := 0

	// Find all matching strings
	mul := regex_mul.FindAllStringSubmatchIndex(input, -1)
	dos := regex_do.FindAllStringSubmatchIndex(input, -1)
	donts := regex_dont.FindAllStringSubmatchIndex(input, -1)

	do := true
	for i := 0; i < len(input); i++ {
		if len(dos) > 0 && dos[0][0] == i {
			do = true
			dos = dos[1:]
		} else if len(donts) > 0 && donts[0][0] == i {
			do = false
			donts = donts[1:]
		}

		if len(mul) > 0 && mul[0][0] == i {
			if !part2 || do {
				result += calc(input[mul[0][4]:mul[0][5]], input[mul[0][6]:mul[0][7]])
			}

			mul = mul[1:]
		}
	}

	return result
}
