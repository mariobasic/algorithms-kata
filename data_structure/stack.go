package data_structure

type Stack[T any] struct {
	head *SNode[T]
	len  int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	s.head = &SNode[T]{value: value, prev: s.head}
	s.len++

	// alternative general idea for some other language maybe

	//if s.head == nil {
	//	s.head = &SNode[T]{value: value}
	//	s.len++
	//	return
	//}
	//
	//node := &SNode[T]{value: value, prev: s.head}
	//s.head = node
	//s.len++
}

func (s *Stack[T]) Pop() *T {
	if s.len == 0 {
		return nil
	}
	val := s.head.value
	s.head = s.head.prev
	s.len--

	return &val
}

func (s *Stack[T]) Peek() *T {
	if s.len == 0 {
		return nil
	}
	return &s.head.value
}
