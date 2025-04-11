package data_structure

type Queue[T any] struct {
	head *QNode[T]
	tail *QNode[T]
	len  int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(value T) {
	if q.tail == nil {
		q.head = &QNode[T]{value: value}
		q.tail = q.head
		q.len++
		return
	}
	q.len++
	q.tail.next = &QNode[T]{value: value}
	q.tail = q.tail.next
}

// Dequeue does this need to be a pointer if it's empty probably yes
func (q *Queue[T]) Dequeue() *T {
	if q.len == 0 || q.head == nil {
		return nil
	}

	current := q.head
	q.head = current.next
	q.len--

	if q.len == 0 {
		q.tail = nil
	}

	return &current.value
}

func (q *Queue[T]) Peek() *T {
	if q.len == 0 || q.head == nil {
		return nil
	}

	return &q.head.value
}

func (q *Queue[T]) Len() int {
	return q.len
}
