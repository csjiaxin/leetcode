package l1254

import (
	"reflect"
	"testing"
)

func TestClosedIsland(t *testing.T) {
	for _, tt := range []struct {
		grid [][]int
		want int
	}{
		{nil, 0}, {[][]int{{}}, 0}, {[][]int{{0}}, 0}, {[][]int{{1}}, 0},
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, 1},
		{[][]int{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}}, 1},
		{[][]int{{1, 0, 1}, {1, 0, 1}, {1, 1, 1}}, 0},
		{[][]int{{1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}}, 2},
		{[][]int{{0, 0}, {0, 0}}, 0}, {[][]int{{0, 1, 0}}, 0}, {[][]int{{0}, {1}, {0}}, 0},
		{[][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}}, 2},
	} {
		var before [][]int
		if tt.grid != nil {
			before = make([][]int, len(tt.grid))
			for r := range tt.grid {
				before[r] = append(make([]int, 0, len(tt.grid[r])), tt.grid[r]...)
			}
		}
		for i := 0; i < 2; i++ {
			if got := closedIsland(tt.grid); got != tt.want {
				t.Fatalf("%v: got %d want %d", tt.grid, got, tt.want)
			}
		}
		if !reflect.DeepEqual(before, tt.grid) {
			t.Fatal("input mutated")
		}
	}
}
