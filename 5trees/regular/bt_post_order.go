package regular

func postOrderSearch(head *BinaryNode[int]) []int {
	return postOrderWalk(head, new(Path)).path
}

func postOrderWalk(curr *BinaryNode[int], path *Path) *Path {
	if curr == nil {
		return path
	}

	postOrderWalk(curr.left, path)
	postOrderWalk(curr.right, path)
	path.push(curr.value)
	return path
}
