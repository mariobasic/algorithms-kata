package _heap

import (
	"cmp"
	"errors"
)

type MinHeap[T cmp.Ordered] struct {
	len  int
	data []T
}

// Len because we are using slice and not array list we could just use this
func (h *MinHeap[T]) Len() int {
	return len(h.data)
}

func (h *MinHeap[T]) Insert(value T) {
	h.data = append(h.data, value)
	h.heapifyUp(h.len)
	h.len++
}

func (h *MinHeap[T]) Delete() (T, error) {
	if h.len == 0 {
		return *new(T), errors.New("heap is empty")
	}

	val := h.data[0] // head value
	h.len--
	// did we have just one node on heap?
	if h.len == 0 {
		h.data = make([]T, 0) // restart heap we are empty
		return val, nil
	}
	// swap head and tail
	h.data[0] = h.data[h.len]
	h.data = h.data[:h.len]
	// heapifyDown
	h.heapifyDown(0)
	return val, nil
}

// heapifyDown used when we delete an element
func (h *MinHeap[T]) heapifyDown(idx int) {
	lIdx := h.leftChild(idx)
	rIdx := h.rightChild(idx)
	// we are at the end of  the "tree"
	if idx >= h.len || lIdx >= h.len {
		return
	}
	lVal := h.data[lIdx]
	v := h.data[idx]

	if rIdx >= h.len {
		// we are at the end of the tree there is no right child
		// swap left child with parent if parent is larger
		if v > lVal {
			h.data[idx], h.data[lIdx] = lVal, v
			return
		}
	}
	rVal := h.data[rIdx]

	if lVal > rVal && v > rVal {
		//swap with rChild
		h.data[idx], h.data[rIdx] = rVal, v
		h.heapifyDown(rIdx)
	} else if rVal > lVal && v > lVal {
		//swap with lChild
		h.data[idx], h.data[lIdx] = lVal, v
		h.heapifyDown(lIdx)
	}
}

// heapifyUp used when we add element
func (h *MinHeap[T]) heapifyUp(idx int) {
	if idx == 0 {
		return
	}
	pIdx := h.parent(idx)
	pVal := h.data[pIdx]
	v := h.data[idx]

	if pVal > v {
		// swap nodes
		//h.data[idx], h.data[pIdx] = h.data[pIdx], h.data[idx]
		h.data[idx], h.data[pIdx] = pVal, v
		h.heapifyUp(pIdx)
	}

}

func (h *MinHeap[T]) parent(idx int) int {
	//return h.data[(idx-1)/2]
	return (idx - 1) / 2
}

func (h *MinHeap[T]) leftChild(idx int) int {
	return idx*2 + 1
}

func (h *MinHeap[T]) rightChild(idx int) int {
	return idx*2 + 2
}
