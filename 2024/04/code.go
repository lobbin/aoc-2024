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

func run(part2 bool, input string) any {
	// Setup
	lines := strings.Split(input, "\n")
	w, h := len(lines[0]), len(lines)-1
	g := grid.New[byte](image.Rect(0, 0, w, h))

	// Create grid
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			g.Set(image.Point{x, y}, lines[y][x])
		}
	}

	// Setup and prepare for parts2
	find := "MAS"
	cross := byte('A')
	crossMap := make(map[image.Point]int)

	found := 0
	dirs := grid.Points(grid.DirectionsDiagonal...)

	if !part2 {
		// Make changes for part 1
		find = "XMAS"
		dirs = grid.Points(grid.DirectionsALL...)
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			start := image.Point{x, y}

			// For each grid point, check if we have a potential starting point
			if g.MustGet(start) == find[0] {
				// If we do, loop all possible directions we might find the puzzle word
			dir_loop:
				for _, dir := range dirs {
					var crossPoint image.Point
					point := start

					// Check remaining word letters
					for i := 1; i < len(find); i++ {
						point = point.Add(dir)
						b, exists := g.Get(point)

						// Abort early
						if !exists || b != find[i] {
							continue dir_loop
						}

						// For part 2, we only want those that make up an X
						if part2 && b == cross {
							crossPoint = point
						}

						// If we end up here, we found a valid word
						if i == len(find)-1 {
							crossMap[crossPoint]++
							found++
						}
					}
				}
			}
		}
	}

	if part2 {
		// For part2, were only interested in those words that shares an X with
		// another diagonal word.
		found = 0
		for _, count := range crossMap {
			if count == 2 {
				found++
			}
		}
	}

	return found
}
