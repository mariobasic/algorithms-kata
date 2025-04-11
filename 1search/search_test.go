package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_linearSearch(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5}

	found := linearSearch(haystack, 4)
	assert.True(t, found)
}

func Test_binarySearchList(t *testing.T) {
	tests := []struct {
		name     string
		want     bool
		haystack []int
		needle   int
	}{
		{
			name:     "first",
			want:     true,
			haystack: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
			needle:   7,
		},
		{
			name:     "second",
			want:     true,
			haystack: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
			needle:   9,
		},
		{
			name:     "third",
			want:     true,
			haystack: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
			needle:   10,
		},
		{
			name:     "forth",
			want:     false,
			haystack: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			needle:   59,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := binarySearchListRecursion(tt.haystack, tt.needle)
			assert.Equal(t, tt.want, got)

			got = binarySearchList(tt.haystack, tt.needle)
			assert.Equal(t, tt.want, got)
		})
	}

}

func Test_twoCristalBalls(t *testing.T) {
	tests := []struct {
		name      string
		args      []bool
		wantIndex int
	}{
		{
			name:      "first",
			args:      []bool{false, false, false, false, false, false, false, false, false, true, true},
			wantIndex: 9,
		},
		{
			name:      "second",
			args:      []bool{false, false, false, false, false, false, false, false, false, false, true},
			wantIndex: 10,
		},
		{
			name:      "out",
			args:      []bool{false, false, false, false, false, false, false, false, false, false, false},
			wantIndex: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			index := twoCristalBalls(tt.args)
			require.Equal(t, tt.wantIndex, index)
		})
	}
}
