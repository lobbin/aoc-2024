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

type Visited struct {
	point image.Point
	dir   Direction
}

type Direction uint8

const (
	North Direction = iota
	East
	South
	West
)

var coords = []image.Point{
	{X: 0, Y: -1}, // N
	{X: 1, Y: 0},  // E
	{X: 0, Y: 1},  // S
	{X: -1, Y: 0}, // W
}

var turn = map[Direction]Direction{
	North: East,
	East:  South,
	South: West,
	West:  North,
}

// Try to move guard, if obstacle change direction, returns false if off grid
func move(g *grid.Map[byte], dir Direction, guard image.Point) (image.Point, Direction, bool) {
	p := guard.Add(coords[dir])

	b, ok := g.Get(p)
	if ok && b == '#' {
		new_direction := turn[dir]
		return guard, new_direction, true
	}

	return p, dir, ok
}

func run(part2 bool, input string) any {
	// Setup
	lines := strings.Split(input, "\n")
	w, h := len(lines[0]), len(lines)
	g := grid.New[byte](image.Rect(0, 0, w, h))
	var start_point image.Point

	// Create grid
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			point := lines[y][x]
			if lines[y][x] == '^' {
				start_point = image.Point{x, y}
			}

			g.Set(image.Point{x, y}, point)
		}
	}

	guard_dir := North
	guard_point := start_point
	visited := make(map[image.Point]bool)
	visited[guard_point] = true

	// Loop until guard is out of bounds
	for {
		var ok bool
		guard_point, guard_dir, ok = move(g, guard_dir, guard_point)

		if !ok {
			break
		}

		visited[guard_point] = true
	}

	if !part2 {
		return len(visited)
	}

	blocks := 0
	for visit := range visited {
		// For all visited points, we place an obstacle and checks whether that makes
		// the guard go into an infinite loop
		if visit == start_point {
			continue
		}

		g.Set(visit, '#')

		guard_point = start_point
		guard_dir = North
		visited_points := map[Visited]bool{
			{guard_point, guard_dir}: true,
		}

		var ok bool
		for {
			guard_point, guard_dir, ok = move(g, guard_dir, guard_point)
			if !ok {
				break
			}

			// Intifite loops is defined as we already been at this point before with
			// the same direction
			v := Visited{guard_point, guard_dir}
			if _, exists := visited_points[v]; exists {
				blocks++
				break
			}
			visited_points[v] = true
		}

		g.Set(visit, '.')
	}

	return blocks
}
