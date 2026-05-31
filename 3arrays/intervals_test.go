package _arrays

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				intervals:   [][]int{{1, 3}, {4, 6}},
				newInterval: []int{2, 5},
			},
			want: [][]int{{1, 6}},
		},

		{
			name: "2",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {9, 10}},
				newInterval: []int{6, 7},
			},
			want: [][]int{{1, 2}, {3, 5}, {6, 7}, {9, 10}},
		},
		{
			name: "3",
			args: args{
				intervals:   [][]int{},
				newInterval: []int{1, 3},
			},
			want: [][]int{{1, 3}},
		},
		{
			name: "4",
			args: args{
				intervals:   [][]int{},
				newInterval: []int{1, 3},
			},
			want: [][]int{{1, 3}},
		},
		{
			name: "5",
			args: args{
				intervals:   [][]int{{1, 5}},
				newInterval: []int{2, 3},
			},
			want: [][]int{{1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{intervals: [][]int{{1, 3}, {1, 5}, {6, 7}}},
			want: [][]int{{1, 5}, {6, 7}},
		},
		{
			name: "2",
			args: args{intervals: [][]int{{1, 2}, {2, 3}}},
			want: [][]int{{1, 3}},
		},
		{
			name: "3",
			args: args{intervals: [][]int{{1, 4}, {2, 3}}},
			want: [][]int{{1, 4}},
		},
		{
			name: "4",
			args: args{intervals: [][]int{{1, 2}, {3, 4}}},
			want: [][]int{{1, 2}, {3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{intervals: [][]int{{1, 2}, {2, 4}, {1, 4}}},
			want: 1,
		},
		{
			name: "2",
			args: args{intervals: [][]int{{1, 2}, {2, 4}}},
			want: 0,
		},
		{
			name: "3",
			args: args{intervals: [][]int{{1, 4}, {2, 3}, {3, 5}}},
			want: 1,
		},
		{
			name: "4",
			args: args{intervals: [][]int{{1, 5}, {2, 3}, {3, 4}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eraseOverlapIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canAttendMeetings(t *testing.T) {
	tests := []struct {
		name      string
		intervals []Interval
		expected  bool
	}{
		// --- Given examples ---
		{
			name:      "given: overlap at start and middle",
			intervals: []Interval{{0, 30}, {5, 10}, {15, 20}},
			expected:  false,
		},
		{
			name:      "given: no overlap",
			intervals: []Interval{{5, 8}, {9, 15}},
			expected:  true,
		},

		// --- Touching endpoints (edge case) ---
		{
			name:      "touching endpoints are NOT a conflict",
			intervals: []Interval{{0, 8}, {8, 10}},
			expected:  true, // end == start is allowed
		},

		// --- Overlap variants ---
		{
			name:      "fully contained interval",
			intervals: []Interval{{1, 10}, {3, 5}},
			expected:  false,
		},
		{
			name:      "partial overlap",
			intervals: []Interval{{1, 5}, {4, 8}},
			expected:  false,
		},
		{
			name:      "exact duplicate intervals",
			intervals: []Interval{{1, 5}, {1, 5}},
			expected:  false,
		},

		// --- Order independence ---
		{
			name:      "unsorted input, no conflict",
			intervals: []Interval{{9, 15}, {5, 8}}, // reversed order
			expected:  true,
		},
		{
			name:      "unsorted input, with conflict",
			intervals: []Interval{{15, 20}, {0, 30}, {5, 10}}, // shuffled
			expected:  false,
		},

		// --- Boundary / trivial cases ---
		{
			name:      "single interval",
			intervals: []Interval{{1, 5}},
			expected:  true,
		},
		{
			name:      "empty intervals",
			intervals: []Interval{},
			expected:  true,
		},
		{
			name:      "multiple non-overlapping in order",
			intervals: []Interval{{1, 2}, {3, 4}, {5, 6}},
			expected:  true,
		},
		{
			name:      "multiple touching endpoints in sequence",
			intervals: []Interval{{0, 5}, {5, 10}, {10, 15}},
			expected:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canAttendMeetings(tt.intervals); got != tt.expected {
				t.Errorf("canAttendMeetings() = %v, want %v", got, tt.expected)
			}
		})
	}
}
