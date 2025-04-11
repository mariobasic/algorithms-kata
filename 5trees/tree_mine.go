package _trees

type Tree[E comparable] struct {
	val         E
	left, right *Tree[E]
}

func (t *Tree[E]) walk(yield func(a E) bool) bool {
	if t == nil {
		return false
	}
	lInOrd := t.left.walk(yield)
	val := yield(t.val) // return value to range
	rInOrd := t.right.walk(yield)
	rslt := lInOrd && val && rInOrd

	return rslt
}

func (t *Tree[E]) Walk(yield func(a E) bool) {
	t.walk(yield)
}
