package regular

func inOrderSearch(head *BinaryNode[int]) []int {
	return inOrderWalk(head, new(Path)).path
}

func inOrderWalk(curr *BinaryNode[int], path *Path) *Path {
	if curr == nil {
		return path
	}

	inOrderWalk(curr.left, path)
	path.push(curr.value)
	inOrderWalk(curr.right, path)
	return path
}
