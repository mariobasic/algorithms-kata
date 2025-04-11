package data_structure

type QNode[T any] struct {
	value T
	next  *QNode[T]
}

type SNode[T any] struct {
	value T
	prev  *SNode[T]
}
