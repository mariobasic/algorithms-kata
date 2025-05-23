package _graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWeightedAdjacencyMatrix_all(t *testing.T) {
	//     >(1)<--->(4) ---->(5)
	//    /          |       /|
	// (0)     ------|------- |
	//    \   v      v        v
	//     >(2) --> (3) <----(6)
	tt := getMatrix()
	expected := []int{0, 1, 4, 5, 6}
	got := bfs(tt, 0, 6)
	assert.Equal(t, expected, got)
	assert.Nil(t, bfs(tt, 6, 0))
}

func getMatrix() WeightedAdjacencyMatrix {
	return WeightedAdjacencyMatrix{

		{0, 3, 1, 0, 0, 0, 0}, // 0
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 5, 0, 2, 0},
		{0, 0, 18, 0, 0, 0, 1},
		{0, 0, 0, 1, 0, 0, 0},
	}
}
