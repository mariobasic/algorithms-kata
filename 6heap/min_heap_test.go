package _heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinHeap_All(t *testing.T) {
	heap := &MinHeap[int]{}

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(69)
	heap.Insert(420)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(7)

	assert.Equal(t, 8, heap.len)
	assert.Equal(t, 1, assertDelete(t, heap))
	assert.Equal(t, 3, assertDelete(t, heap))
	assert.Equal(t, 4, assertDelete(t, heap))
	assert.Equal(t, 5, assertDelete(t, heap))

	assert.Equal(t, 4, heap.len)

	assert.Equal(t, 7, assertDelete(t, heap))
	assert.Equal(t, 8, assertDelete(t, heap))
	assert.Equal(t, 69, assertDelete(t, heap))
	assert.Equal(t, 420, assertDelete(t, heap))

	assert.Equal(t, 0, heap.len)

}

func assertDelete(t *testing.T, heap *MinHeap[int]) int {
	val, err := heap.Delete()
	assert.NoError(t, err)
	return val
}
