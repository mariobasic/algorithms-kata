package _trees

import (
	"fmt"
	"testing"
)

func TestTree_Inorder(t1 *testing.T) {
	// Create a sample tree:
	//
	//       10
	//      /  \
	//     20  40
	//    /  \
	//   30  39

	tt := &Tree[int]{
		10,
		&Tree[int]{
			20,
			&Tree[int]{30, nil, nil},
			&Tree[int]{39, nil, nil}},
		&Tree[int]{40, nil, nil},
	}

	for v := range tt.Walk {
		fmt.Println(v)
	}
}
