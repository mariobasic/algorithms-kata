package _trees

import (
	ds "algorithms/data_structure"
	"cmp"
)

func (t *BinaryNode[T]) breadthFirstSearch(needle T) bool {
	q := &ds.Queue[BinaryNode[T]]{}
	q.Enqueue(*t)

	for q.Len() > 0 {
		curr := q.Dequeue()
		if curr.value == needle {
			return true
		}
		if curr.left != nil {
			q.Enqueue(*curr.left)
		}
		if curr.right != nil {
			q.Enqueue(*curr.right)

		}
	}
	return false
}

func bfs[T cmp.Ordered](head *BinaryNode[T], needle T) bool {
	q := []*BinaryNode[T]{head}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.value == needle {
			return true
		}
		if curr.left != nil {
			q = append(q, curr.left)
		}
		if curr.right != nil {
			q = append(q, curr.right)
		}
	}
	return false
}
