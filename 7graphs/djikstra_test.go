package _graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pq_dijkstraList(t *testing.T) {
	//     >(1)<--->(4) ---->(5)
	//    /          |       /|
	// (0)     ------|------- |
	//    \   v      v        v
	//     >(2) --> (3) <----(6)
	tt := getGraph1()
	distancesToNodes := []int{0, 3, 1, 8, 4, 6, 7}
	assert.Equal(t, distancesToNodes, dijkstraPq(0, 6, tt))
}

func Test_dijkstraList(t *testing.T) {
	//     >(1)<--->(4) ---->(5)
	//    /          |       /|
	// (0)     ------|------- |
	//    \   v      v        v
	//     >(2) --> (3) <----(6)
	tt := getGraph1()
	expected := []int{0, 1, 4, 5, 6}
	assert.Equal(t, expected, dijkstraList(0, 6, tt))
	assert.Nil(t, dijkstraList(6, 0, tt))
}

func getGraph1() WeightedAdjacencyList {
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
