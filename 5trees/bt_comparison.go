package _trees

import "cmp"

func compare[T cmp.Ordered](a, b *BinaryNode[T]) bool {
	// structural check
	// are we structurally the same?
	if a == nil && b == nil {
		return true
	}
	// structural check
	// are we structurally not the same?
	if a == nil || b == nil {
		return false
	}

	// value check
	if a.value != b.value {
		return false
	}

	return compare(a.left, b.left) && compare(a.right, b.right)
}
