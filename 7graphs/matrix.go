package _graphs

import "slices"

type WeightedAdjacencyMatrix = [][]int

func bfs(graph WeightedAdjacencyMatrix, source, needle int) []int {

	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	for i := 0; i < len(prev); i++ {
		prev[i] = -1
	}

	seen[source] = true
	q := []int{source}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == needle {
			break
		}
		adjs := graph[cur] // who is this node connected to
		for i := 0; i < len(adjs); i++ {
			if adjs[i] == 0 || seen[i] {
				continue
			}
			seen[i] = true
			prev[i] = cur
			q = append(q, i)

		}
	}
	// we didn't visit the needle node/vertex
	if prev[needle] == -1 {
		return nil
	}

	// build it backwards
	curr := needle
	var out []int
	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	out = append(out, source)
	slices.Reverse(out)
	return out
}
