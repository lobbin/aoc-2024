package main

import (
	"image"
	"math"
	"slices"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
	"github.com/s0rg/grid"
)

func main() {
	aoc.Harness(run)
}

type Direction int8

const (
	North Direction = iota
	East
	South
	West
)

type Path struct {
	pos    image.Point
	path   []image.Point
	dir    Direction
	points int
}

var moves = map[Direction]image.Point{
	North: {0, -1},
	East:  {1, 0},
	South: {0, 1},
	West:  {-1, 0},
}

func turnDirs(dir Direction) []Direction {
	if dir == North || dir == South {
		return []Direction{West, East}
	} else /* if dir == Right || dir == Left */ {
		return []Direction{North, South}
	}
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

	score := math.MaxInt
	paths := []Path{{start, []image.Point{start}, East, 0}}

	visited := make(map[image.Point]int)
	best_seats := make(map[image.Point]bool)

	// i := 0
	for len(paths) > 0 {
		path := paths[0]
		paths = paths[1:]

		// We got to the end, checking if it was the best way
		if path.pos == end {
			if path.points <= score {
				if path.points < score {
					clear(best_seats)
				}

				score = path.points

				for _, point := range path.path {
					best_seats[point] = true
				}
			}
			continue
		}

		// No need to keep looking at higher scores
		if path.points >= score {
			continue
		}

		// Try moving
		move := path.pos.Add(moves[path.dir])
		if points, exists := visited[move]; (!exists || path.points+1 <= points) && g.MustGet(move) != '#' {
			visited[move] = path.points + 1
			if part2 {
				visited_path := make([]image.Point, len(path.path))
				copy(visited_path, path.path)
				visited_path = append(visited_path, move)
				paths = append(paths, Path{move, visited_path, path.dir, path.points + 1})
			} else {
				paths = append(paths, Path{move, path.path, path.dir, path.points + 1})
			}
		}

		// Try turning
		turns := turnDirs(path.dir)
		for _, turn := range turns {
			move = path.pos.Add(moves[turn])
			if points, exists := visited[move]; (!exists || path.points+1000 <= points+1) && g.MustGet(move) != '#' {
				if part2 {
					visited[path.pos] += 1000
				}
				paths = append(paths, Path{path.pos, path.path, turn, path.points + 1000})
			}
		}

		// Sort to make sure we always process the most likely to win first
		slices.SortFunc(paths, func(a, b Path) int {
			return a.points - b.points
		})
	}

	if part2 {
		return len(best_seats)
	} else {
		return score
	}
}
