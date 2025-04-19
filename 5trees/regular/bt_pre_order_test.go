package regular

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_preOrderSearch(t *testing.T) {
	// Create a sample tree:
	//          20
	//      /          \
	//     10            50
	//    /  \         /   \
	//   5   15      30      100
	//    \         /   \
	//     7       29     45
	tt := getTree()

	expected := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	got := preOrderSearch(tt)
	assert.Equal(t, expected, got)
}

func getTree() *BinaryNode[int] {
	tt := &BinaryNode[int]{
		value: 20,
		right: &BinaryNode[int]{
			value: 50,
			right: &BinaryNode[int]{
				value: 100,
				left:  nil,
				right: nil,
			},
			left: &BinaryNode[int]{
				value: 30,
				right: &BinaryNode[int]{
					value: 45,
					left:  nil,
					right: nil,
				},
				left: &BinaryNode[int]{
					value: 29,
					left:  nil,
					right: nil,
				},
			},
		},
		left: &BinaryNode[int]{
			value: 10,
			right: &BinaryNode[int]{
				value: 15,
				left:  nil,
				right: nil,
			},
			left: &BinaryNode[int]{
				value: 5,
				right: &BinaryNode[int]{
					value: 7,
					left:  nil,
					right: nil,
				},
				left: nil,
			},
		},
	}
	return tt
}
