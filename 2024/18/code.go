package main

import (
	"image"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
	"github.com/s0rg/grid"
)

func main() {
	aoc.Harness(run)
}

// Walks the grid to find the shortest path
func checkGridPath(g *grid.Map[byte], w int, h int) ([]image.Point, bool) {
	path, ok := g.Path(
		image.Point{0, 0},
		image.Point{w - 1, h - 1},
		grid.Points(grid.DirectionsCardinal...),
		grid.DistanceManhattan,
		func(p image.Point, f float64, b byte) (float64, bool) {
			if b == '#' {
				return 0, false
			}
			return 1, true
		},
	)
	return path, ok
}

func run(part2 bool, input string) any {
	// Basics
	w, h := 7, 7
	bytes := 12
	lines := strings.Split(input, "\n")
	if len(lines) > 25 {
		w, h = 71, 71
		bytes = 1024
	}

	// Setup grid
	g := grid.New[byte](image.Rect(0, 0, w, h))
	g.Fill(func() byte { return '.' })

	// Path variables
	var path []image.Point
	var path_ok bool

	// Loop all "bytes"
	for i, line := range lines {
		// For the first path, we have a limit number of bytes
		if !part2 && i >= bytes {
			break
		}

		xy := strings.Split(line, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		g.Set(image.Point{x, y}, '#')

		// For part2, one byte will eventually make the grid non-walkable
		if part2 && i >= bytes {
			_, path_ok = checkGridPath(g, w, h)
			if !path_ok {
				return line
			}
		}
	}

	// In first part, we know there's a working path
	if !part2 {
		path, _ = checkGridPath(g, w, h)
	}

	// g.Iter(func(p image.Point, b byte) bool {
	// 	if p.X == w-1 {
	// 		fmt.Println(string(b))
	// 	} else {
	// 		fmt.Print(string(b))
	// 	}
	// 	return true
	// })

	return len(path) - 1 // includes first step
}
