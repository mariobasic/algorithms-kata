package data_structure

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	queue := NewQueue[int]()
	require.NotNil(t, queue)
	require.Equal(t, 0, queue.len)
	require.Nil(t, queue.head)
	require.Nil(t, queue.tail)
}

func TestQueue_PeekEnqueueDequeue(t *testing.T) {
	type testCase[T any] struct {
		name     string
		populate func(queue *Queue[T])
		check    func(queue *Queue[T])
	}
	tests := []testCase[string]{
		{
			name: "first",
			populate: func(queue *Queue[string]) {
				queue.Enqueue("a")
				queue.Enqueue("b")
				queue.Enqueue("c")
			},
			check: func(queue *Queue[string]) {
				require.Equal(t, "a", queue.Peek())
				require.Equal(t, queue.head, queue.Peek())
				require.Equal(t, 3, queue.len)
				require.Equal(t, "a", *queue.Dequeue())
				require.Equal(t, 2, queue.len)

				require.Equal(t, "b", *queue.Dequeue())
				require.Equal(t, 1, queue.len)
				require.Equal(t, "c", *queue.Dequeue())
				require.Equal(t, 0, queue.len)
				require.Nil(t, queue.head)
				require.Nil(t, queue.tail)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewQueue[string]()
			tt.populate(queue)
		})
	}
}
