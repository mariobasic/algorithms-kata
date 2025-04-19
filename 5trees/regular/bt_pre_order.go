package regular

type BinaryNode[T comparable] struct {
	value       T
	left, right *BinaryNode[T]
}
type Path struct {
	path []int
}

func (p *Path) push(val int) {
	p.path = append(p.path, val)
}

func preOrderSearch(head *BinaryNode[int]) []int {
	return walk(head, new(Path)).path
}

func walk(curr *BinaryNode[int], path *Path) *Path {
	if curr == nil {
		return path
	}

	path.push(curr.value)
	walk(curr.left, path)
	walk(curr.right, path)
	return path
}
