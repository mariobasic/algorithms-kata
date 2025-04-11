package _recursion

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		maze  []string
		wall  rune
		start Point
		end   Point
	}
	tests := []struct {
		name string
		args args
		want []Point
	}{
		{
			name: "first",
			args: args{
				maze: []string{
					"xxxxxxxxxx x",
					"x        x x",
					"x        x x",
					"x xxxxxxxx x",
					"x          x",
					"x xxxxxxxxxx",
				},
				wall:  'x',
				start: Point{x: 10, y: 0},
				end:   Point{x: 1, y: 5},
			},
			want: getResultSlice(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.maze, tt.args.wall, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getResultSlice() []Point {
	return []Point{
		{x: 10, y: 0},
		{x: 10, y: 1},
		{x: 10, y: 2},
		{x: 10, y: 3},
		{x: 10, y: 4},
		{x: 9, y: 4},
		{x: 8, y: 4},
		{x: 7, y: 4},
		{x: 6, y: 4},
		{x: 5, y: 4},
		{x: 4, y: 4},
		{x: 3, y: 4},
		{x: 2, y: 4},
		{x: 1, y: 4},
		{x: 1, y: 5},
	}
}
