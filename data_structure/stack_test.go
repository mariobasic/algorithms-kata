package data_structure

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack[int]()
	require.NotNil(t, stack)
	require.Equal(t, 0, stack.len)
	require.Nil(t, stack.head)
}

func TestStack_Peek(t *testing.T) {
	type testCase[T any] struct {
		name string
		s    func(stack *Stack[T])
		want func(stack *Stack[T])
	}
	tests := []testCase[int]{
		{
			name: "first",
			s: func(stack *Stack[int]) {
				stack.Push(1)
				stack.Push(2)
				stack.Push(3)
			},
			want: func(stack *Stack[int]) {
				peek := stack.Peek()
				require.Equal(t, 3, *peek)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack[int]()
			tt.s(stack)
			tt.want(stack)
		})
	}
}

func TestStack_PushPop(t *testing.T) {
	type testCase[T any] struct {
		name string
		s    func(stack *Stack[T])
		want func(stack *Stack[T])
	}
	tests := []testCase[string]{
		{
			name: "first",
			s: func(stack *Stack[string]) {
				stack.Push("a")
				stack.Push("b")
				stack.Push("c")
			},
			want: func(stack *Stack[string]) {
				require.Equal(t, 3, stack.len)
				require.Equal(t, "c", *stack.Pop())
				require.Equal(t, 2, stack.len)
				require.Equal(t, "b", *stack.Pop())
				require.Equal(t, 1, stack.len)
				require.Equal(t, "a", *stack.Pop())
				require.Nil(t, stack.Pop())
				require.Equal(t, 0, stack.len)
				stack.Push("d")
				require.Equal(t, 1, stack.len)
				require.Equal(t, "d", *stack.Peek())
				require.Equal(t, "d", *stack.Pop())
				require.Nil(t, stack.Peek())
				require.Nil(t, stack.Pop())

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack[string]()
			tt.s(stack)
			tt.want(stack)
		})
	}
}
