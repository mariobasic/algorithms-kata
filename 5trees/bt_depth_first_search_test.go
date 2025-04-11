package _trees

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryNode_Walk(t *testing.T) {
	// Create a sample tree:
	//
	//       10
	//      /  \
	//     20  40
	//    /  \
	//   30  39

	tt := &BinaryNode[int]{
		value: 10,
		left: &BinaryNode[int]{
			value: 20,
			left:  &BinaryNode[int]{value: 30, left: nil, right: nil},
			right: &BinaryNode[int]{value: 39, left: nil, right: nil},
		},
		right: &BinaryNode[int]{value: 40, left: nil, right: nil},
	}

	for v := range tt.WalkPreOrder {
		fmt.Println(v)
	}
}

func TestBinaryNode_WalkPreOrder(t *testing.T) {
	// Create a sample tree:
	//
	//          20
	//      /          \
	//     10            50
	//    /  \         /   \
	//   5   15      30      100
	//    \         /   \
	//     7       29     45
	tt := getTree()

	expected := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	var got []int
	for v := range tt.WalkPreOrder {
		got = append(got, v)
	}

	assert.Equal(t, expected, got)
}

func TestBinaryNode_WalkPostOrder(t *testing.T) {
	// Create a sample tree:
	//
	//          20
	//      /          \
	//     10            50
	//    /  \         /   \
	//   5   15      30      100
	//    \         /   \
	//     7       29     45
	tt := getTree()

	expected := []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}

	var got []int
	for v := range tt.WalkPostOrder {
		got = append(got, v)
	}

	assert.Equal(t, expected, got)

}

func TestBinaryNode_WalkInOrder(t *testing.T) {
	// Create a sample tree:
	//          20
	//      /          \
	//     10            50
	//    /  \         /   \
	//   5   15      30      100
	//    \         /   \
	//     7       29     45
	tt := getTree()

	expected := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}

	var got []int
	for v := range tt.WalkInOrder {
		got = append(got, v)
	}

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

func getTree2() *BinaryNode[int] {
	return &BinaryNode[int]{
		value: 20,
		right: &BinaryNode[int]{
			value: 50,
			right: nil,
			left: &BinaryNode[int]{
				value: 30,
				right: &BinaryNode[int]{
					value: 45,
					right: &BinaryNode[int]{
						value: 49,
						left:  nil,
						right: nil,
					},
					left: nil,
				},
				left: &BinaryNode[int]{
					value: 29,
					right: nil,
					left: &BinaryNode[int]{
						value: 21,
						left:  nil,
						right: nil,
					},
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

}
