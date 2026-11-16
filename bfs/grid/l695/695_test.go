package l695

import (
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	for _, tt := range []struct {
		grid [][]int
		want int
	}{
		{nil, 0}, {[][]int{{}}, 0}, {[][]int{{0}}, 0}, {[][]int{{1}}, 1},
		{[][]int{{1, 1, 0, 1}, {0, 1, 0, 1}}, 3},
		{[][]int{{1, 0}, {0, 1}}, 1}, {[][]int{{1, 1}, {1, 1}}, 4},
		{[][]int{{1, 1, 0, 1, 1, 1}}, 3}, {[][]int{{1}, {1}, {0}, {1}}, 2},
	} {
		before := clone(tt.grid)
		for i := 0; i < 2; i++ {
			if got := maxAreaOfIsland(tt.grid); got != tt.want {
				t.Fatalf("%v: got %d want %d", tt.grid, got, tt.want)
			}
		}
		if !reflect.DeepEqual(before, tt.grid) {
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

func TestFloodFillVisited(t *testing.T) {
	visited := [][]bool{make([]bool, 3), make([]bool, 3)}
	grid := [][]int{{1, 1, 0}, {1, 1, 1}}
	counts := map[[2]int]int{}
	canEnter := func(r, c int) bool { return grid[r][c] == 1 }
	visit := func(r, c int) {
		if !visited[r][c] {
			t.Fatal("visit before marking")
		}
		counts[[2]int{r, c}]++
	}
	FloodFill(2, 3, 0, 0, visited, canEnter, visit)
	if len(counts) != 5 {
		t.Fatalf("visited %v", counts)
	}
	for _, n := range counts {
		if n != 1 {
			t.Fatal("duplicate visit")
		}
	}
	FloodFill(2, 3, 1, 2, visited, canEnter, visit)
	if counts[[2]int{1, 2}] != 1 {
		t.Fatal("visited did not persist")
	}
	FloodFill(2, 3, 0, 2, visited, canEnter, func(int, int) { t.Fatal("visited blocked cell") })
	FloodFill(2, 3, -1, 0, visited, func(int, int) bool { t.Fatal("out of bounds callback"); return true }, visit)
	if visited[0][2] {
		t.Fatal("water marked")
	}
}
