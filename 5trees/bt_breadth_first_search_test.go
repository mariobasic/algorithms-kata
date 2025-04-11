package _trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryNode_breadthFirstSearch(t1 *testing.T) {
	// Create a sample tree:
	//
	//          20
	//      /          \
	//     10            50
	//    /  \         /    \
	//   5   15       30     100
	//    \         /   \
	//     7       29     45
	tt := getTree()
	assert.True(t1, tt.breadthFirstSearch(45))
	assert.True(t1, tt.breadthFirstSearch(7))
	assert.False(t1, tt.breadthFirstSearch(69))

	assert.True(t1, bfs(tt, 45))
	assert.True(t1, bfs(tt, 7))
	assert.False(t1, bfs(tt, 69))
}
