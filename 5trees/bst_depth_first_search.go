package _trees

func (t *BinaryNode[T]) dfs(needle T) bool {
	if t == nil {
		return false
	}

	if t.value == needle {
		return true
	}
	if t.value < needle {
		return t.right.dfs(needle)
	}

	return t.left.dfs(needle)
}
