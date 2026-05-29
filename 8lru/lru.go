package _lru

import "cmp"

type Node[T any] struct {
	value      T
	prev, next *Node[T]
}

func createNode[V any](value V) *Node[V] {
	return &Node[V]{value: value}
}

type LRU[K cmp.Ordered, V cmp.Ordered] struct {
	len           int
	cap           int
	head          *Node[V] // list of values
	tail          *Node[V]
	lookup        map[K]*Node[V]
	reverseLookup map[*Node[V]]K
}

func newLRU[K, V cmp.Ordered](cap int) LRU[K, V] {
	return LRU[K, V]{
		lookup:        make(map[K]*Node[V], cap),
		reverseLookup: make(map[*Node[V]]K, cap),
		cap:           cap,
	}
}

func (l *LRU[K, V]) update(key K, value V) {
	// does it exist?
	node, ok := l.lookup[key]

	// if yes -> update -> move to the front
	if ok {
		l.detach(node)
		l.prepend(node)
		node.value = value
		return
	}

	// if not -> insert
	// 	- capacity greater than needed -> evict
	node = createNode(value)
	l.len++
	l.prepend(node)
	l.trimCache()
	l.lookup[key] = node
	l.reverseLookup[node] = key
}

func (l *LRU[K, V]) get(key K) (V, bool) {
	// does it exist?
	node, ok := l.lookup[key]
	if !ok {
		var zero V
		return zero, false
	}

	// update the value we found and move it to the front
	l.detach(node)
	l.prepend(node)

	// return out the value found or nil if it doesn't exist
	return node.value, true
}

func (l *LRU[K, V]) detach(node *Node[V]) {
	// point node's previous next to nodes next
	if node.prev != nil {
		node.prev.next = node.next
	}
	// point node's next previous to nodes previous
	if node.next != nil {
		node.next.prev = node.prev
	}
	//are we the head node?
	if l.head == node {
		l.head = l.head.next
	}
	// are we the tail?
	if l.tail == node {
		l.tail = l.tail.prev
	}

	node.next = nil
	node.prev = nil
	node = nil
}

func (l *LRU[K, V]) prepend(node *Node[V]) {
	if l.head == nil {
		l.head, l.tail = node, node
	}
	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *LRU[K, V]) trimCache() {
	if l.len <= l.cap {
		return
	}
	tail := l.tail
	l.detach(l.tail)

	// remove from lookups
	key := l.reverseLookup[tail]
	delete(l.lookup, key)
	delete(l.reverseLookup, tail)
	l.len--
}
