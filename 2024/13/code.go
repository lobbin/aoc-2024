package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

type Play struct {
	x, y int
	a, b int
}

func main() {
	aoc.Harness(run)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func run(part2 bool, input string) any {
	total_tokens := 0

	regex_number := regexp.MustCompile(`\d+`)
	regex_result := regex_number.FindAllStringSubmatch(strings.ReplaceAll(input, "\n", ","), -1)
	for i := 0; i < len(regex_result); i += 6 {
		ax, _ := strconv.Atoi(regex_result[i][0])
		ay, _ := strconv.Atoi(regex_result[i+1][0])
		bx, _ := strconv.Atoi(regex_result[i+2][0])
		by, _ := strconv.Atoi(regex_result[i+3][0])
		rx, _ := strconv.Atoi(regex_result[i+4][0])
		ry, _ := strconv.Atoi(regex_result[i+5][0])
		best := math.MaxInt

		if part2 {
			factor := 10_000_000_000_000
			fx := factor + rx
			fy := factor + ry

			// max_a := min(fx/ax, fy/ay)
			max_b := min(fx/bx, fy/by)
			fmt.Println(ax, ay, bx, by, max_b, rx, ry)
			fmt.Println(max_b*bx, max_b*by, fx, fy)

			//main_loop_part2:
			for b := max_b; b >= 0; b-- {
				need_ax := (fx - b*bx) / ax
				need_ay := (fy - b*by) / ay

				if need_ax == need_ay {
					token := need_ax*3 + b
					fmt.Println("Win", b, need_ax, token)
					break
				}

				// if math.Abs(float64(need_ax)-float64(need_ay)) < 100 {
				// 	min_a := min(need_ax, need_ay)
				// 	fmt.Println("MEH", need_ax, need_ay, min_a)
				// 	for a := min_a - 1; a <= max_b; a++ {
				// 		cx := a*ax + b*bx
				// 		cy := a*ay + b*by
				// 		token := a*3 + b
				// 		if cx == fx && cy == fy {
				// 			fmt.Println("Win", token)
				// 			break main_loop_part2
				// 		} else if cx > fx || cy > fy {
				// 			//fmt.Println("Too big")
				// 			continue main_loop_part2
				// 		}
				// 	}
				// }

			}

		} else {
		main_loop:
			for b := 100; b >= 0; b-- {
				for a := 0; a <= 100; a++ {
					if a*ax+b*bx == rx && a*ay+b*by == ry {
						token := a*3 + b
						if token < best {
							best = token
							break main_loop
						}
					}
				}
			}
		}

		if best != math.MaxInt {
			total_tokens += best
		}
	}

	return total_tokens
}
