package l1905

import (
	"reflect"
	"testing"
)

func TestCountSubIslands(t *testing.T) {
	for _, tt := range []struct {
		a, b [][]int
		want int
	}{
		{nil, nil, 0}, {[][]int{{}}, [][]int{{}}, 0},
		{[][]int{{1}}, [][]int{{1}}, 1}, {[][]int{{0}}, [][]int{{1}}, 0},
		{[][]int{{1}}, [][]int{{0}}, 0},
		{[][]int{{1, 0, 1}}, [][]int{{1, 1, 1}}, 0},
		{[][]int{{1, 1, 1}}, [][]int{{1, 0, 1}}, 2},
		{[][]int{{1, 0}, {0, 1}}, [][]int{{1, 0}, {0, 1}}, 2},
		{[][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}, [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}, 3},
	} {
		a, b := clone(tt.a), clone(tt.b)
		for i := 0; i < 2; i++ {
			if got := countSubIslands(tt.a, tt.b); got != tt.want {
				t.Fatalf("a=%v b=%v: got %d want %d", tt.a, tt.b, got, tt.want)
			}
		}
		if !reflect.DeepEqual(a, tt.a) || !reflect.DeepEqual(b, tt.b) {
			t.Fatal("input mutated")
		}
	}
}
func clone(g [][]int) [][]int {
	if g == nil {
		return nil
	}
	out := make([][]int, len(g))
	for r := range g {
		out[r] = append(make([]int, 0, len(g[r])), g[r]...)
	}
	return out
}
