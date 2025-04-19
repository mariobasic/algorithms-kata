package sort

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want []int
	}{
		{
			name: "first",
			arg:  []int{1, 3, 2, 4, 6, 5, 7, 9, 8, 10},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "second",
			arg:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//got := bubbleSort(tt.arg)
			//require.Equal(t, tt.want, got)

			got := bubbleSort(tt.arg)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_quickSort(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want []int
	}{
		{
			name: "first",
			arg:  []int{1, 3, 2, 4, 6, 5, 7, 9, 8, 10},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "second",
			arg:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "third",
			arg:  []int{9, 3, 7, 4, 69, 420, 42},
			want: []int{3, 4, 7, 9, 42, 69, 420},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := quickSort(tt.arg)
			require.Equal(t, tt.want, got)
		})
	}
}
