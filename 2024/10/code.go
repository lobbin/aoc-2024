package main

import (
	"image"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
	"github.com/s0rg/grid"
)

type PathList []image.Point

type Path struct {
	path  PathList
	pos   image.Point
	value byte
}

func main() {
	aoc.Harness(run)
}

// Find all possible paths.
func findPaths(g *grid.Map[byte], start image.Point, trail_heads map[image.Point]int) int {
	queue := []Path{
		{PathList{start}, start, '0'},
	}

	dirs := grid.Points(grid.DirectionsCardinal...)
	heads := make(map[image.Point]bool)
	paths := make([]PathList, 0)

	// Working the grid, BFS wise
	for len(queue) > 0 {
		// Take first entry
		entry := queue[0]
		queue = queue[1:]

		// Work the neighbours, only interested in those leading to an end or the
		// ones that are one value higher than our current position
		g.Neighbours(entry.pos, dirs, func(p image.Point, b byte) bool {
			if b-1 == entry.value {
				if b == '9' {
					paths = append(paths, append(entry.path, p))
					heads[p] = true
				} else {
					n := Path{append(entry.path, p), p, b}
					queue = append(queue, n)
				}
			}
			return true
		})
	}

	// For part1, we want the number of unique trail heads
	for head := range heads {
		trail_heads[head]++
	}

	// For part2, we want the number of possible paths
	return len(paths)
}

func run(part2 bool, input string) any {
	// Setup
	lines := strings.Split(input, "\n")
	w, h := len(lines[0]), len(lines)
	g := grid.New[byte](image.Rect(0, 0, w, h))

	// Create grid
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			g.Set(image.Point{x, y}, lines[y][x])
		}
	}

	ratings := 0
	trail_heads := make(map[image.Point]int)

	// Loop all grid, looking for starting points
	g.Iter(func(p image.Point, b byte) bool {
		if b == '0' {
			ratings += findPaths(g, p, trail_heads)
		}
		return true
	})

	if part2 {
		return ratings
	} else {
		score := 0
		for _, count := range trail_heads {
			score += count
		}

		return score
	}
}
