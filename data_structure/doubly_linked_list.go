package data_structure

import (
	"errors"
)

type LinkedList[T comparable] interface {
	length() int
	insertAt(item T, index int) error
	remove(item T)
	removeAt(index int) T
	append(item T)
	prepend(item T)
	get(index int) T
	getAt(index int) *Node[T]
}

type Node[T comparable] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type DoublyLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

func NewDoublyLinkedList[T comparable]() *Node[T] {
	return &Node[T]{}
}

func (d *DoublyLinkedList[T]) length() int {
	//TODO implement me
	panic("implement me")
}

func (d *DoublyLinkedList[T]) insertAt(item T, index int) error {
	if index < 0 || index > d.len {
		return errors.New("index out of range")
	}
	if index == d.len {
		d.append(item)
		return nil
	} else if index == 0 {
		d.prepend(item)
		return nil
	}
	d.len++
	cur := d.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	node := Node[T]{value: item, next: cur, prev: cur.prev}
	cur.prev.next = &node
	cur.prev = &node

	return nil
}

func (d *DoublyLinkedList[T]) remove(item T) {
	curr := d.head

	for i := 0; i < d.len; i++ {
		if curr.value == item {
			break
		}
		curr = curr.next
	}

	if curr == nil {
		return
	}

	if d.len == 0 {
		return
	}

	if curr.prev != nil {
		curr.prev.next = curr.next
	}
	if curr.next != nil {
		curr.next.prev = curr.prev
	}

	if curr == d.head {
		d.head = curr.next
	}
	if curr == d.tail {
		d.tail = curr.prev
	}
	curr.next = nil
	curr.prev = nil

	d.len--
}

func (d *DoublyLinkedList[T]) removeAt(index int) T {
	node := d.getAt(index)
	if node == nil {
		return *new(T)
	}
	d.remove(node.value)
	return node.value // this is bad
}

func (d *DoublyLinkedList[T]) append(item T) {
	d.len++
	node := &Node[T]{value: item}
	if d.len == 0 {
		d.head = node
		d.tail = node
		return
	}
	node.prev = d.tail
	d.tail.next = node
	d.tail = node
}

func (d *DoublyLinkedList[T]) prepend(item T) {
	node := Node[T]{value: item}
	d.len++
	if d.head == nil {
		d.head = &node
		d.tail = &node
		return
	}

	node.next = d.head
	d.head.prev = &node
	d.head = &node
}

func (d *DoublyLinkedList[T]) get(index int) T {
	return d.getAt(index).value
}

func (d *DoublyLinkedList[T]) getAt(index int) *Node[T] {
	curr := d.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	return curr
}
