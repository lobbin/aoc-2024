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

type Antinode struct {
	p image.Point
	t byte
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

	// Loop entire grip, looking for antennas
	antinodes := make(map[Antinode]struct{})
	g.Iter(func(p image.Point, b byte) bool {
		if b != '.' {
			// When found an antenna, loop grid again, trying to find matching ones
			g.Iter(func(p2 image.Point, b2 byte) bool {
				if p != p2 && b == b2 {
					// If found, calculate the diff and create antinode based on the diff
					diff := p.Sub(p2)
					antinode := p2

					// Since part2 resonances in each direction and we only a calculates
					// the antinodes going _from_ each antenna, we need to manually add
					// the antinode that will exist at each antenna
					if part2 {
						antinodes[Antinode{p, b2}] = struct{}{}
					}

					// If not out-of-bounds, we store the point. For part 2 the signal
					// will continue in that direction until we not longer within the
					// grid.
					for {
						antinode = antinode.Sub(diff)
						if _, ok := g.Get(antinode); ok {
							antinodes[Antinode{antinode, b2}] = struct{}{}
						} else {
							break
						}

						if !part2 {
							break
						}
					}
				}
				return true
			})
		}
		return true
	})

	// We're only interested in the unique points
	unique := make(map[image.Point]struct{})
	for antinode := range antinodes {
		unique[antinode.p] = struct{}{}
	}

	return len(unique)
}
