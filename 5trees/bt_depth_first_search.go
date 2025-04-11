package _trees

import "cmp"

type BinaryNode[T cmp.Ordered] struct {
	value       T
	left, right *BinaryNode[T]
}

func (t *BinaryNode[T]) walkPreOrder(yield func(T) bool) bool {
	if t == nil {
		return false
	}

	val := yield(t.value)
	lWalk := t.left.walkPreOrder(yield)
	rWalk := t.right.walkPreOrder(yield)

	return val && lWalk && rWalk
}

func (t *BinaryNode[T]) WalkPreOrder(yield func(T) bool) {
	t.walkPreOrder(yield)
}

func (t *BinaryNode[T]) walkPostOrder(yield func(T) bool) bool {
	if t == nil {
		return false
	}

	lWalk := t.left.walkPostOrder(yield)
	rWalk := t.right.walkPostOrder(yield)
	val := yield(t.value)

	return val && lWalk && rWalk
}

func (t *BinaryNode[T]) WalkPostOrder(yield func(T) bool) {
	t.walkPostOrder(yield)
}
func (t *BinaryNode[T]) walkInOrder(yield func(T) bool) bool {
	if t == nil {
		return false
	}

	lWalk := t.left.walkInOrder(yield)
	val := yield(t.value)
	rWalk := t.right.walkInOrder(yield)

	return val && lWalk && rWalk
}

func (t *BinaryNode[T]) WalkInOrder(yield func(T) bool) {
	t.walkInOrder(yield)
}
