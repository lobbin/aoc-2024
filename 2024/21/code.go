package main

import (
	"hash/fnv"
	"image"
	"math"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
	"github.com/s0rg/grid"
	"github.com/x1m3/priorityQueue"
)

func main() {
	aoc.Harness(run)
}

type RuneCombo struct {
	r1, r2 rune
}

type PathItem []image.Point

func (p PathItem) HigherPriorityThan(next priorityQueue.Interface) bool {
	return len(p) < len(next.(PathItem))
}

// When walking all directions, check if point is valid within our map
func validPoint(pad map[rune]image.Point, point image.Point) bool {
	for _, v := range pad {
		if v == point {
			return true
		}
	}
	return false
}

// Turn a path of points into a string
func pathToString(path PathItem) string {
	s := strings.Builder{}

	for i := 1; i < len(path); i++ {
		switch path[i-1].Sub(path[i]) {
		case image.Point{-1, 0}:
			s.WriteRune('>')
		case image.Point{0, -1}:
			s.WriteRune('v')
		case image.Point{1, 0}:
			s.WriteRune('<')
		case image.Point{0, 1}:
			s.WriteRune('^')
		}
	}

	s.WriteRune('A')
	return s.String()
}

// Turn map into a slice
func mapToSlice(m map[string]bool) []string {
	s := []string{}
	for k := range m {
		s = append(s, k)
	}
	return s
}

// Generate the best possible paths between two points
func bestPaths(pad map[rune]image.Point, start, end image.Point) []string {
	pathMap := map[string]bool{}
	best := math.MaxInt

	dirs := grid.Points(grid.DirectionsCardinal...)

	visited := map[image.Point]int{start: 1}

	// Priority queue is superflouous here, but was fun to use
	queue := priorityQueue.New()
	queue.Push(PathItem{start})

	for {
		item := queue.Pop()
		if item == nil {
			break
		}

		path := item.(PathItem)
		pathLen := len(path)
		current := path[len(path)-1]

		if pathLen > best {
			// No need to process the since they are worse
			return mapToSlice(pathMap)
		} else if current == end {
			// If we end up here we have a valid path, let's save the possible way to
			// get here
			best = pathLen
			pathMap[pathToString(path)] = true
		} else if visitScore, exists := visited[current]; !exists || pathLen <= visitScore {
			// We're only interested in paths that are equal or better than paths we've
			// seen before
			visited[current] = pathLen
			for _, dir := range dirs {
				new := current.Add(dir)

				// Double check it's actually a valid point (in how our pads look) before
				// proceeding
				if validPoint(pad, new) {
					queue.Push(append(PathItem{}, append(path, new)...))
				}
			}
		}

	}

	return mapToSlice(pathMap)
}

// Generate the possible paths between all the buttons in a keypad
func possiblePaths(pad map[rune]image.Point) map[RuneCombo][]string {
	paths := map[RuneCombo][]string{}

	for ak, a := range pad {
		for bk, b := range pad {
			paths[RuneCombo{ak, bk}] = bestPaths(pad, a, b)
		}
	}

	return paths
}

type DecodeQueue struct {
	length, code int
}

type CacheKey struct {
	codeHash uint64
	robots   int
}

// Generate a hash for the code string
func hashCode(code string) uint64 {
	hasher := fnv.New64a()
	hasher.Write([]byte(code))
	return hasher.Sum64()
}

// Recursive function to handle a variable number of robots that needs to be
// handled

func handleRobots(code string, robots int, padPath, arrowPaths map[RuneCombo][]string, cache map[CacheKey]int) int {
	// Using a hash for the cache key together with the level of robots
	hash := hashCode(code)
	cacheKey := CacheKey{hash, robots}
	if c, exists := cache[cacheKey]; exists {
		return c
	}

	bestPath := math.MaxInt
	code = "A" + code

	// We need to use a queue, since we might have several options getting from
	// one button to another
	queue := []DecodeQueue{{0, 1}}

	// Process queue BFS wise
	for len(queue) > 0 {
		decode := queue[0]
		queue = queue[1:]

		if decode.code == len(code) {
			// If we're at the end, check whether we have the best path so far
			if decode.length < bestPath {
				bestPath = decode.length
			}
		} else {
			// Process the step from one button to the next and record it recursivly
			// until we have the prober number of robots handled
			a, b := rune(code[decode.code-1]), rune(code[decode.code])
			paths := padPath[RuneCombo{a, b}]
			for _, path := range paths {
				length := len(path)
				if robots != 0 {
					length = handleRobots(path, robots-1, arrowPaths, arrowPaths, cache)
				}

				queue = append(queue, DecodeQueue{decode.length + length, decode.code + 1})
			}
		}
	}

	// Cache and return
	cache[cacheKey] = bestPath
	return bestPath
}

func run(part2 bool, input string) any {
	// Define the arrow pad and calculate the best paths between all the buttons
	arrowPad := map[rune]image.Point{
		'^': {1, 0}, 'A': {2, 0},
		'<': {0, 1}, 'v': {1, 1}, '>': {2, 1},
	}
	arrowPaths := possiblePaths(arrowPad)

	// Define the numeric pad and calculate the best paths between all the buttons
	numericPad := map[rune]image.Point{
		'7': {0, 0}, '8': {1, 0}, '9': {2, 0},
		'4': {0, 1}, '5': {1, 1}, '6': {2, 1},
		'1': {0, 2}, '2': {1, 2}, '3': {2, 2},
		'0': {1, 3}, 'A': {2, 3},
	}
	numericPaths := possiblePaths(numericPad)

	iterations := 2
	if part2 {
		iterations = 25
	}

	score := 0
	cache := map[CacheKey]int{}

	// Loop all the available codes
	for _, code := range strings.Split(input, "\n") {
		sequence := handleRobots(code, iterations, numericPaths, arrowPaths, cache)
		number, _ := strconv.Atoi(code[:len(code)-1])

		score += number * sequence
	}

	return score
}
