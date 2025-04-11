package _graphs

type GraphEdge struct {
	to, weight int
}

type WeightedAdjacencyList [][]GraphEdge

func walk(graph WeightedAdjacencyList, curr, needle int, seen []bool, path *[]int) bool {
	if seen[curr] {
		return false
	}

	// pre
	seen[curr] = true
	*path = append(*path, curr)
	if curr == needle {
		return true
	}
	// recurse
	list := graph[curr] // who is this node connected  to
	for i := 0; i < len(list); i++ {
		edge := list[i]
		if walk(graph, edge.to, needle, seen, path) {
			return true
		}
	}

	// post
	*path = (*path)[:len(*path)-1]
	return false
}

func Dfs(graph WeightedAdjacencyList, source, needle int) []int {
	seen := make([]bool, len(graph))
	var path []int
	walk(graph, source, needle, seen, &path)
	if len(path) == 0 {
		return nil
	}
	return path
}
