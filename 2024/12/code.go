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

type Plot struct {
	area, perimeter int
	label           byte
	points          []image.Point
}

type OuterCorner struct {
	p1, p2 int
}
type InnerCorner struct {
	p1, p2, p3 int
}

// This is ugly, but somehow I can't get it into the main loop
func countSides(points []image.Point) int {
	pointSet := make(map[image.Point]bool)
	for _, p := range points {
		pointSet[p] = true
	}

	dirs := grid.Points(grid.DirectionsALL...)
	// DirectionsALL = []dir{
	// 	NorthWest,
	// 	North,
	// 	NorthEast,
	// 	East,
	// 	SouthEast,
	// 	South,
	// 	SouthWest,
	// 	West,
	// }
	outerCorners := []OuterCorner{
		{1, 3}, // north and east
		{1, 7}, // north and west
		{3, 5}, // east and south
		{5, 7}, // south and west
	}
	innerCorners := []InnerCorner{
		{1, 3, 2}, // north and east, not northeast
		{3, 5, 4}, // east and south, not southest
		{5, 7, 6}, // south and west, not southwest
		{7, 1, 0}, // west and north, not northwest
	}

	sides := 0
	didPoints := make([]image.Point, len(dirs))
	for _, p := range points {
		for i, dir := range dirs {
			didPoints[i] = p.Add(dir)
		}

		// Count all the outer corners
		for _, oc := range outerCorners {
			if !pointSet[didPoints[oc.p1]] && !pointSet[didPoints[oc.p2]] {
				sides++
			}
		}

		// Count all the inner corners
		for _, in := range innerCorners {
			if pointSet[didPoints[in.p1]] && pointSet[didPoints[in.p2]] && !pointSet[didPoints[in.p3]] {
				sides++
			}
		}
	}

	return sides
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

	plots := make([]*Plot, 0)

	// Iterate grid
	dirPoints := grid.Points(grid.DirectionsCardinal...)
	processed := make(map[image.Point]bool)

	g.Iter(func(p image.Point, b byte) bool {
		if _, exists := processed[p]; !exists {
			plot := &Plot{0, 0, b, make([]image.Point, 0)}

			// Some sort of flood-fill
			bfs := []image.Point{p}
			for len(bfs) > 0 {
				entry := bfs[0]
				bfs = bfs[1:]

				if _, exists := processed[entry]; !exists {
					plot.points = append(plot.points, entry)
					processed[entry] = true

					g.Neighbours(entry, dirPoints, func(p2 image.Point, b2 byte) bool {
						if b == b2 {
							bfs = append(bfs, p2)
						}
						return true
					})
				}
			}

			plots = append(plots, plot)
		}

		return true
	})

	// Calculate perimeter for part1 and price for part2. This is ugly as well as
	// I'd like to get everything in the same loop, but here we are :)
	price := 0
	for _, plot := range plots {
		if part2 {
			sides := countSides(plot.points)
			price += sides * len(plot.points)
		} else {
			for _, point := range plot.points {
				for _, dir := range dirPoints {
					neighbour := point.Add(dir)
					b, ok := g.Get(neighbour)

					if !ok || b != plot.label {
						plot.perimeter++
					}
				}
			}
		}

	}

	if part2 {
		return price
	} else {
		checksum := 0
		for _, plot := range plots {
			checksum += plot.perimeter * len(plot.points)
		}
		return checksum
	}
}
