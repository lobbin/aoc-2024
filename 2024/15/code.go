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

// Recursive function to move double boxes around. The function returns a list
// of moved boxes so we can backtrack this when we have a split path that doesn't
// work out in the end.
func moveBox(g *grid.Map[rune], command rune, box *Box, boxes []*Box) []*Box {
	var move []*Box
	var diff image.Point

	if command == '<' || command == '>' {
		// Left and right are easier to handle since these only affect one box at a
		// time
		var moveTo image.Point
		if command == '<' {
			diff = image.Point{-1, 0}
			moveTo = box.p1.Add(diff)
		} else {
			diff = image.Point{1, 0}
			moveTo = box.p2.Add(diff)
		}

		if g.MustGet(moveTo) != '.' {
			// It means we can't even move a box this way
			return nil
		}

		// See if there's another box in the way and if so, try to see if we can
		// move it
		hasBlockingBox, blockingBox := hasBox(moveTo, boxes)
		if hasBlockingBox {
			move = moveBox(g, command, blockingBox, boxes)
		} else {
			move = append(move, box)
		}
	} else if command == '^' || command == 'v' {
		// Up and down are a bit different, since one box can affect potentially two
		// other boxes.
		var invert rune
		if command == '^' {
			diff = image.Point{0, -1}
			invert = 'v'
		} else {
			diff = image.Point{0, 1}
			invert = '^'
		}
		mt1, mt2 := box.p1.Add(diff), box.p2.Add(diff)

		if g.MustGet(mt1) != '.' || g.MustGet(mt2) != '.' {
			// It means we can't even move a box this way
			return nil
		}

		hbb1, bb1 := hasBox(mt1, boxes)
		hbb2, bb2 := hasBox(mt2, boxes)
		if hbb1 && hbb2 {
			// If we have two boxes, we must make sure we can move both of them before
			// we can move the dependent box
			if bb1 != bb2 {
				// Try first box
				first := moveBox(g, command, bb1, boxes)
				if first != nil {
					// First box succeeded, let's try the sceond
					move = moveBox(g, command, bb2, boxes)
					if move == nil {
						// If we can't move the second box, we need to the revert the first
						// and all the boxes that was involved in moving the first box
						for _, putBack := range first {
							moveBox(g, invert, putBack, boxes)
						}
					}
				}
			} else {
				// It's the same box so no special care needed
				move = moveBox(g, command, bb1, boxes)
			}
		} else if hbb1 {
			// Only half of our box touches the left box
			move = moveBox(g, command, bb1, boxes)
		} else if hbb2 {
			// Only half of out box touches the right box
			move = moveBox(g, command, bb2, boxes)
		} else {
			// No box in the way, just move!
			move = append(move, box)
		}
	}

	// Make the move if needed
	if move != nil {
		box.p1 = box.p1.Add(diff)
		box.p2 = box.p2.Add(diff)
	}

	return move
}

// Checks whether we can move into the next point
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

	// Extract grid and boxes
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

			// For part2, we need to double to coordinate set. To solve part2 two, I'm
			// breaking out the boxes in its own object list instead of keeping them
			// in the grid
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

	// Extract commands
	for i := h + 1; i < len(lines); i++ {
		for _, r := range lines[i] {
			commands = append(commands, r)
		}
	}

	// Define directions and inverted directions (for part1)
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

	// Loop commands
	for len(commands) > 0 {
		command := commands[0]
		commands = commands[1:]

		direction := directions[command]
		inverted := inverted_dirs[command]

		if part2 {
			next := robot.Add(direction)

			// Checking if we have a box and if so, try to move the box(es)
			if has, box := hasBox(next, boxes); has {
				moveBox(g, command, box, boxes)
			}

			// Checking if we can move (could also be walls, etc)
			if canMove(g, next, boxes) {
				robot = next
			}
		} else {
			// Find the next free spot and move in that direction until we walk into
			// something
			next := robot
			for {
				next = next.Add(direction)
				r := g.MustGet(next)
				if r == '#' {
					break
				}

				if r == '.' {
					break
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

			// Now make the actual move, if possible
			next = robot.Add(direction)
			if g.MustGet(next) == '.' {
				robot = next
			}
		}
	}

	fmt.Println()

	// Print final grid and calculate the score
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
				score += p.X + p.Y*100
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
