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

type Q struct {
	pos    image.Point
	length int
}

// Find the only available path and return the path itself
func bfs(g *grid.Map[byte], start, end image.Point) map[image.Point]int {
	queue := []Q{{start, 0}}
	visited := map[image.Point]int{}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		visited[path.pos] = path.length
		if path.pos == end {
			break
		}

		g.Neighbours(path.pos, grid.Points(grid.DirectionsCardinal...), func(p image.Point, b byte) bool {
			if _, exists := visited[p]; !exists && b != '#' {
				queue = append(queue, Q{p, path.length + 1})
			}
			return true
		})
	}

	return visited
}

// Since we have a limitation of two steps we can do them directly
func bfsWithCheats(visited map[image.Point]int) []int {
	finished := []int{}

	// From each point in the path, we see if moving two steps would save us any
	// time by looking at the original visited distance
	for point, distance := range visited {
		steps := []image.Point{{-2, 0}, {2, 0}, {0, -2}, {0, 2}}
		for _, step := range steps {
			cheat := point.Add(step)
			cheatDistance, exists := visited[cheat]

			// We're only interested in visited points
			if exists {
				diff := cheatDistance - distance - 2
				if diff > 0 {
					finished = append(finished, diff)
				}
			}
		}
	}

	return finished
}

// We could solve part1 with this solution as well, but it's a bit slower :)
func bfsWithCheats2(visited map[image.Point]int) []int {
	finished := []int{}

	// Reverse the map
	distanceToPoint := map[int]image.Point{}
	for k, v := range visited {
		distanceToPoint[v] = k
	}

	// From the start, we look at all valid points from out starting point and
	// see if we valid a point short track point within the manhattan distance
	// of 20. If so, we check whether going there would actually save us time in
	// the end.
	for i := 0; i < len(distanceToPoint); i++ {
		p1 := distanceToPoint[i]
		for j := i + 1; j < len(distanceToPoint); j++ {
			p2 := distanceToPoint[j]
			distance := grid.DistanceManhattan(p1, p2)

			if distance <= 20 {
				diff := visited[p2] - visited[p1] - int(distance)
				if diff > 0 {
					finished = append(finished, diff)
				}
			}
		}
	}

	return finished
}

func run(part2 bool, input string) any {
	// Setup
	lines := strings.Split(input, "\n")
	w, h := len(lines[0]), len(lines)
	g := grid.New[byte](image.Rect(0, 0, w, h))

	var start, end image.Point

	// Create grid
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			p := image.Point{x, y}
			if lines[y][x] == 'S' {
				start = p
			} else if lines[y][x] == 'E' {
				end = p
			}
			g.Set(p, lines[y][x])
		}
	}

	// Run the working path
	originalPath := bfs(g, start, end)

	saves := make(map[int]int)
	var finishedPaths []int

	// Run the cheated paths
	if part2 {
		finishedPaths = bfsWithCheats2(originalPath)
	} else {
		finishedPaths = bfsWithCheats(originalPath)
	}

	// Reduce
	for _, finished := range finishedPaths {
		saves[finished]++
	}

	// Check how many total steps we can save if we save more than 100 steps
	goodSaves := 0
	for save, count := range saves {
		if save >= 100 {
			goodSaves += count
		}
	}

	return goodSaves
}
