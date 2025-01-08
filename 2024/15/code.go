package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
	"github.com/s0rg/grid"
)

func main() {
	aoc.Harness(run)
}

type Box struct {
	p1, p2 image.Point
}

func hasBox(p image.Point, boxes []*Box) (bool, *Box) {
	for _, box := range boxes {
		if box.p1 == p || box.p2 == p {
			return true, box
		}
	}

	return false, nil
}

// func canMoveBox(g *grid.Map[rune], direction rune, next image.Point, box *Box) bool {

// }

func canMove(g *grid.Map[rune], next image.Point, boxes []*Box) bool {
	if has, _ := hasBox(next, boxes); has {
		return false
	}

	return g.MustGet(next) == '.'
}

func run(part2 bool, input string) any {
	// Setup
	lines := strings.Split(input, "\n")
	w, h := len(lines[0]), len(lines[0])
	if part2 {
		w *= 2
	}
	g := grid.New[rune](image.Rect(0, 0, w, h))

	boxes := make([]*Box, 0)
	var robot image.Point
	commands := make([]rune, 0)

	for y := 0; y < h; y++ {
		line := lines[y]

		for x, r := range line {
			if r == '@' {
				r = '.'
				robot.X, robot.Y = x, y
				if part2 {
					robot.X = x * 2
				}
			}

			if part2 {
				if r == 'O' {
					boxes = append(boxes, &Box{image.Point{x * 2, y}, image.Point{x*2 + 1, y}})
					r = '.'
				}

				g.Set(image.Point{x * 2, y}, r)
				g.Set(image.Point{x*2 + 1, y}, r)
			} else {
				g.Set(image.Point{x, y}, r)
			}
		}
	}

	for i := h + 1; i < len(lines); i++ {
		for _, r := range lines[i] {
			commands = append(commands, r)
		}
	}

	directions := map[rune]image.Point{
		'<': {-1, 0},
		'^': {0, -1},
		'>': {1, 0},
		'v': {0, 1},
	}
	inverted_dirs := map[rune]image.Point{
		'<': {1, 0},
		'^': {0, 1},
		'>': {-1, 0},
		'v': {0, -1},
	}

	i := 0
	for len(commands) > 0 {
		command := commands[0]
		commands = commands[1:]

		fmt.Println("Command", string(command))

		direction := directions[command]
		inverted := inverted_dirs[command]

		if part2 {
			next := robot.Add(direction)

			if has, box := hasBox(next, boxes); has {
				// Try to move the boxes
				moveBox(g, next, command, box)
				fmt.Println("Can't move, box in the way")
			} else {
				fmt.Println("No box in the way", next)
			}

			if canMove(g, next, boxes) {
				robot = next
			}
		} else {
			// Find the next free spot
			next := robot
			for {
				next = next.Add(direction)
				r := g.MustGet(next)
				if r == '#' {
					break
				}

				if r == '.' {
					break
					// if !part2 || !hasBox(next, boxes) {
					// 	break
					// }
				}
			}

			// Backtrack backwards and move all the boxes
			for next != robot {
				next = next.Add(inverted)

				if g.MustGet(next) == 'O' {
					move := next.Add(direction)

					if g.MustGet(move) == '.' {
						g.Set(move, 'O')
						g.Set(next, '.')
					}
				}
			}

			next = robot.Add(direction)
			if g.MustGet(next) == '.' {
				robot = next
			}
		}

		i++
		if i > 5 {
			break
		}
	}

	fmt.Println()

	score := 0
	g.Iter(func(p image.Point, r rune) bool {
		if r == 'O' {
			score += p.X + p.Y*100
		}

		if p == robot {
			r = '@'
		}

		for _, box := range boxes {
			if p == box.p1 {
				r = '['
			} else if p == box.p2 {
				r = ']'
			}
		}

		if p.X == w-1 {
			fmt.Println(string(r))
		} else {
			fmt.Print(string(r))
		}
		return true
	})

	return score
}
