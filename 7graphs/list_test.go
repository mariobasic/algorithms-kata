package _graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWeightedAdjacencyList_All(t *testing.T) {
	//     >(1)<--->(4) ---->(5)
	//    /          |       /|
	// (0)     ------|------- |
	//    \   v      v        v
	//     >(2) --> (3) <----(6)
	tt := getGraph()
	expected := []int{0, 1, 4, 5, 6}
	assert.Equal(t, expected, Dfs(tt, 0, 6))
	assert.Nil(t, Dfs(tt, 6, 0))
}

func getGraph() WeightedAdjacencyList {
	// Create the adjacency list with 7 nodes
	list := make(WeightedAdjacencyList, 7)

	// Populate the graph edges
	list[0] = []GraphEdge{
		{to: 1, weight: 3},
		{to: 2, weight: 1},
	}
	list[1] = []GraphEdge{
		{to: 4, weight: 1},
	}
	list[2] = []GraphEdge{
		{to: 3, weight: 7},
	}
	list[3] = []GraphEdge{}
	list[4] = []GraphEdge{
		{to: 1, weight: 1},
		{to: 3, weight: 5},
		{to: 5, weight: 2},
	}
	list[5] = []GraphEdge{
		{to: 2, weight: 18},
		{to: 6, weight: 1},
	}
	list[6] = []GraphEdge{
		{to: 3, weight: 1},
	}

	return list
}
