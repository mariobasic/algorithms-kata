package _trees

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_compare(t *testing.T) {
	// Create a sample tree:
	//
	//          20
	//      /          \
	//     10           50
	//    /  \         /   \
	//   5   15      30     100
	//    \         /   \
	//     7       29     45
	tt1 := getTree()

	// Create a sample tree:
	//
	//          20
	//      /          \
	//     10            50
	//    /  \         /
	//   5   15      30
	//    \         /   \
	//     7       29     45
	//            /        \
	//          21         49
	tt2 := getTree2()

	require.True(t, compare(tt1, tt1))
	require.False(t, compare(tt1, tt2))

}
