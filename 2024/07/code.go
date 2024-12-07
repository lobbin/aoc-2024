package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// Generate permutations based on +, * and for part two, also |
func generatePermutations(n int, part2 bool) [][]rune {
	if n <= 0 {
		return [][]rune{}
	}

	var result [][]rune
	var current []rune
	chars := []rune{'+', '*'}
	if part2 {
		chars = append(chars, '|')
	}

	var backtrack func(int)
	backtrack = func(pos int) {
		if pos == n {
			perm := make([]rune, n)
			copy(perm, current)
			result = append(result, perm)
			return
		}

		for _, char := range chars {
			current = append(current, char)
			backtrack(pos + 1)
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return result
}

// Concat two integers
func concatIntegers(a, b int) (int, error) {
	concatenated := strconv.Itoa(a) + strconv.Itoa(b)
	return strconv.Atoi(concatenated)
}

// Handle concat operator
func concat(i int, perm_result int, test_values []int) int {
	if i > 0 {
		perm_result, _ = concatIntegers(perm_result, test_values[i+1])
	} else {
		perm_result, _ = concatIntegers(test_values[i], test_values[i+1])
	}
	return perm_result
}

// Handle multiply operator
func multiply(i int, perm_result int, test_values []int) int {
	if i > 0 {
		perm_result = perm_result * test_values[i+1]
	} else {
		perm_result += test_values[i] * test_values[i+1]
	}
	return perm_result
}

// Handle add operator
func add(i int, perm_result int, test_values []int) int {
	if i > 0 {
		perm_result = perm_result + test_values[i+1]
	} else {
		perm_result += test_values[i] + test_values[i+1]
	}
	return perm_result
}

func run(part2 bool, input string) any {
	calibration_result := 0
	// Loop each line, using lazy splitting method of each line
	for _, line := range strings.Split(input, "\n") {
		main := strings.Split(line, ":")

		test_value, _ := strconv.Atoi(main[0])
		values := strings.Split(strings.TrimSpace(main[1]), " ")

		// Converting string values to integers
		test_values := make([]int, 0)
		for _, v := range values {
			n, _ := strconv.Atoi(strings.TrimSpace(v))
			test_values = append(test_values, n)
		}

		// Loop each permutation, doing operator left-to-right and then checking
		// whether we have the desired result
		perms := generatePermutations(len(test_values)-1, part2)
		for _, perm := range perms {
			perm_result := 0
			for i := 0; i < len(perm); i++ {
				if perm[i] == '+' {
					perm_result = add(i, perm_result, test_values)
				} else if perm[i] == '*' {
					perm_result = multiply(i, perm_result, test_values)
				} else {
					perm_result = concat(i, perm_result, test_values)
				}
			}
			if perm_result == test_value {
				calibration_result += test_value
				break
			}
		}
	}

	return calibration_result
}
