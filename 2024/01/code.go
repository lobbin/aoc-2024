package main

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	// Parse file by looping lines, splitting by white space, converting the
	// string to number and adding them to the lists
	list1, list2 := make([]int, 0), make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		numbers := strings.Split(line, " ")
		number1, _ := strconv.Atoi(strings.TrimSpace(numbers[0]))
		number2, _ := strconv.Atoi(strings.TrimSpace(numbers[len(numbers)-1]))
		list1 = append(list1, number1)
		list2 = append(list2, number2)
	}

	result := 0
	if part2 {
		// Pre-count all the numbers in second list
		countMap := make(map[int]int)
		for _, number := range list2 {
			countMap[number]++
		}

		// Iterator first list and calculate using pre-counted values
		for _, number := range list1 {
			result += number * countMap[number]
		}
	} else {
		// Sort lists ascending
		slices.SortFunc(list1, func(i, j int) int { return i - j })
		slices.SortFunc(list2, func(i, j int) int { return i - j })

		// Loop the lists and add the diffrence to the reslut
		for i := 0; i < len(list1); i++ {
			result += int(math.Abs(float64(list1[i]) - float64(list2[i])))
		}
	}

	// Return result
	return result
}
