package main

import (
	"fmt"
	"image"
	"regexp"
	"slices"
	"strconv"

	"github.com/jpillora/puzzler/harness/aoc"
)

type Robot struct {
	px, py int
	vx, vy int
}

type ImagePointList []image.Point

func main() {
	aoc.Harness(run)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func moveRobot(robot *Robot, iterations int, w int, h int) {
	robot.px = robot.px + robot.vx*iterations
	robot.py = robot.py + robot.vy*iterations

	// Wrap things
	robot.px = (robot.px%w + w) % w
	robot.py = (robot.py%h + h) % h
}

func run(part2 bool, input string) any {
	w, h := 101, 103
	robots := make([]*Robot, 0)

	regex_parser := regexp.MustCompile(`p=([-\d]+),([-\d]+) v=([-\d]+),([-\d]+)`)
	parser := regex_parser.FindAllStringSubmatch(input, -1)
	for _, parse := range parser {
		px, _ := strconv.Atoi(parse[1])
		py, _ := strconv.Atoi(parse[2])
		vx, _ := strconv.Atoi(parse[3])
		vy, _ := strconv.Atoi(parse[4])

		robots = append(robots, &Robot{px, py, vx, vy})
	}

	mw := w / 2
	mh := h / 2

	if part2 && len(robots) > 15 {
		i := 0

		// Loop until we find a pattern
	xmas_loop:
		for {
			i++
			ys := make(map[int]ImagePointList)

			// Move all robots
			for _, robot := range robots {
				moveRobot(robot, 1, w, h)

				// Store the robots Y location to see if we can spot a pattern
				ys[robot.py] = append(ys[robot.py], image.Point{robot.px, robot.py})
			}

			// Check whether we have a pattern
			for _, yc := range ys {
				if len(yc) >= 5 {
					slices.SortFunc(yc, func(a, b image.Point) int {
						return a.X - b.X
					})
					found := true
					for i := 1; i < len(yc); i++ {
						if yc[i].X != yc[i-1].X+1 {
							found = false
							break
						}
					}
					if found {
						break xmas_loop
					}
				}
			}
		}

		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				r := "."
				for _, robot := range robots {
					if robot.px == x && robot.py == y {
						r = "X"
					}
				}

				if x == w-1 {
					fmt.Println(r)
				} else {
					fmt.Print(r)
				}
			}
		}

		return i
	} else {
		quadrants := make([]int, 4)

		iterations := 100
		for _, robot := range robots {
			// Calculate new place
			moveRobot(robot, iterations, w, h)

			// Place them in quadrants
			if robot.px < mw && robot.py < mh {
				// top left
				quadrants[0]++
			} else if robot.px > mw && robot.py < mh {
				// top right
				quadrants[1]++
			} else if robot.px < mw && robot.py > mh {
				// bottom left
				quadrants[2]++
			} else if robot.px > mw && robot.py > mh {
				// bottom right
				quadrants[3]++
			}
		}

		score := 1
		for _, quadrant_score := range quadrants {
			score *= quadrant_score
		}

		return score
	}
}
