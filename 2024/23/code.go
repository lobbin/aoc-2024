package main

import (
	"slices"
	"strings"

	"github.com/hmdsefi/gograph"
	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func findLargestSet(g gograph.Graph[string]) []string {
	vertices := g.GetAllVertices()
	n := len(vertices)
	largestSet := []string{}

	// Using Perplexity suggested recursive backtrack method
	var backtrack func(start int, currentSet []string)
	backtrack = func(start int, currentSet []string) {
		// If this is better than the current set
		if len(currentSet) > len(largestSet) {
			largestSet = append([]string{}, currentSet...)
		}

		// Loop all nodes to see if we have it connected to this set
		for i := start; i < n; i++ {
			v := vertices[i]
			if isSet(g, currentSet, v) {
				// See if we can connect more to this set
				backtrack(i+1, append(currentSet, v.Label()))
			}
		}
	}

	// Start empty-handed
	backtrack(0, []string{})
	return largestSet
}

// Check if vertex can be connected to current set
func isSet(g gograph.Graph[string], set []string, v *gograph.Vertex[string]) bool {
	for _, u := range set {
		if !g.ContainsEdge(g.GetVertexByID(u), v) {
			return false
		}
	}
	return true
}

func run(part2 bool, input string) any {
	graph := gograph.New[string]()

	// Add all nodes
	for _, line := range strings.Split(input, "\n") {
		nodes := strings.Split(line, "-")
		graph.AddEdge(gograph.NewVertex(nodes[0]), gograph.NewVertex(nodes[1]))
	}

	if part2 {
		// Find the largest set
		largestSet := findLargestSet(graph)
		slices.Sort(largestSet)

		return strings.Join(largestSet, ",")
	} else {
		// Find nodes where the chief historian might be
		vertices := graph.GetAllVertices()
		n := len(vertices)

		possibleGroups := 0
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if graph.ContainsEdge(vertices[i], vertices[j]) {
					for k := j + 1; k < n; k++ {
						if graph.ContainsEdge(vertices[i], vertices[k]) && graph.ContainsEdge(vertices[j], vertices[k]) {
							if vertices[i].Label()[0] == 't' || vertices[j].Label()[0] == 't' || vertices[k].Label()[0] == 't' {
								possibleGroups++
							}
						}
					}
				}
			}
		}

		return possibleGroups
	}
}
