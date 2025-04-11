package _trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryNode_DFSonBST(t *testing.T) {
	// Create a sample tree:
	//          20
	//      /          \
	//     10            50
	//    /  \         /   \
	//   5   15      30      100
	//    \         /   \
	//     7       29     45
	tt := getTree()
	assert.True(t, tt.dfs(45))
	assert.True(t, tt.dfs(7))
	assert.False(t, tt.dfs(69))
}
