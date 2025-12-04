package main

import (
	"image"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
	"github.com/s0rg/grid"
)

func main() {
	aoc.Harness(run)
}

func isRemovable(g *grid.Map[byte], dirs []image.Point, p image.Point) bool {
	rolls := 0
	for _, dir := range dirs {
		np := p.Add(dir)
		nb, ok := g.Get(np)
		if ok && nb == '@' {
			rolls++
		}

		if rolls >= 4 {
			return false
		}
	}

	return true
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	lines := strings.Split(input, "\n")
	w, h := len(lines[0]), len(lines)
	g := grid.New[byte](image.Rect(0, 0, w, h))

	for y := range h {
		for x := range w {
			g.Set(image.Point{x, y}, lines[y][x])
		}
	}

	dirs := grid.Points(grid.DirectionsALL...)
	score := 0
	if part2 {
		for {
			removables := make([]image.Point, 0, w)

			g.Iter(func(p image.Point, b byte) bool {
				if b == '@' && isRemovable(g, dirs, p) {
					removables = append(removables, p)
				}

				return true
			})

			if len(removables) > 0 {
				score += len(removables)
				for _, p := range removables {
					g.Set(p, '.')
				}
			} else {
				break
			}
		}
	} else {
		g.Iter(func(p image.Point, b byte) bool {
			if b == '@' && isRemovable(g, dirs, p) {
				score++
			}

			return true
		})
	}

	return score
}
