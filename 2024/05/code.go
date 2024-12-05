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

type rule struct {
	a, b int
}

// Swap slice entries by their value
func swapByValue(update []int, val1, val2 int) {
	var index1, index2 int

	// Find the indices of the values
	for i, v := range update {
		if v == val1 {
			index1 = i
		}
		if v == val2 {
			index2 = i
		}
	}

	// Swap the values
	update[index1], update[index2] = update[index2], update[index1]
}

// Find the broken rules and swap the broken values until we have a working
// order
func applyAndFixRules(rules []rule, update []int) []int {
full_loop:
	for {
	rule_loop:
		for _, rule := range rules {
			// We only care if the update contains both values
			if slices.Contains(update, rule.a) && slices.Contains(update, rule.b) {
				for _, n := range update {
					if n == rule.a {
						// If a comes before b, everything is great
						continue rule_loop
					} else if n == rule.b {
						swapByValue(update, rule.a, rule.b)
						if areAllRuleApplied(rules, update) {
							break full_loop
						}
						continue rule_loop
					}
				}
			}
		}
	}

	return update
}

// Verify that all rules are already applied
func areAllRuleApplied(rules []rule, update []int) bool {
rule_loop:
	for _, rule := range rules {
		if slices.Contains(update, rule.a) && slices.Contains(update, rule.b) {
			for _, n := range update {
				if n == rule.a {
					continue rule_loop
				} else if n == rule.b {
					return false
				}
			}
		}
	}
	return true
}

// Calculate result
func result(updates [][]int) int {
	result := 0
	for _, update := range updates {
		result += update[len(update)/2]
	}
	return result
}

func run(part2 bool, input string) any {
	rules := make([]rule, 0)
	updates := make([][]int, 0)

	// Parse input to get rules and updates
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "|") {
			split := strings.Split(line, "|")
			a, _ := strconv.Atoi(split[0])
			b, _ := strconv.Atoi(split[1])
			rules = append(rules, rule{a, b})
		} else if len(line) > 0 {
			split := strings.Split(line, ",")
			update := make([]int, 0)
			for _, s := range split {
				n, _ := strconv.Atoi(s)
				update = append(update, n)
			}
			updates = append(updates, update)
		}
	}

	correctly := make([][]int, 0)
	incorrectly := make([][]int, 0)

	// Verify and process all updates
	for _, update := range updates {
		if areAllRuleApplied(rules, update) {
			correctly = append(correctly, update)
		} else if part2 {
			// For part2, we try to fix the broken updates
			incorrectly = append(incorrectly, applyAndFixRules(rules, update))
		}
	}

	if part2 {
		return result(incorrectly)
	} else {
		return result(correctly)
	}
}
