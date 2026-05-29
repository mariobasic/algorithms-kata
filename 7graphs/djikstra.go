package _graphs

import (
	"math"
	"slices"
)

func dijkstraList(
	source int,
	sink int,
	arr WeightedAdjacencyList,
) []int {

	prev := make([]int, len(arr))
	seen := make([]bool, len(arr))
	dists := make([]int, len(arr))
	for i := range dists {
		dists[i] = math.MaxInt
	}
	for i := range prev {
		prev[i] = -1
	}

	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true

		adjs := arr[curr]
		for i := range adjs {
			edge := adjs[i]
			if seen[edge.to] {
				continue
			}
			dist := dists[curr] + edge.weight
			if dist < dists[edge.to] {
				dists[edge.to] = dist
				prev[edge.to] = curr
			}
		}
	}

	var out []int
	curr := sink

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	if len(out) > 0 {
		out = append(out, source)
		slices.Reverse(out)
		return out
	}

	return nil
}

func hasUnvisited(seen []bool, dists []int) bool {

	for i := range seen {
		// I didn't visit it, but I do have weights for it?
		if !seen[i] && dists[i] < math.MaxInt {
			return true
		}
	}
	return false
}

func getLowestUnvisited(seen []bool, dists []int) int {
	idx := -1
	lowestDistance := math.MaxInt

	for i := range seen {
		if seen[i] {
			continue
		}

		if lowestDistance > dists[i] {
			lowestDistance = dists[i]
			idx = i
		}
	}

	return idx
}
