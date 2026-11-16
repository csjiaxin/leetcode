package l827

import (
	"reflect"
	"testing"
)

func TestLargestIsland(t *testing.T) {
	for _, tt := range []struct {
		grid [][]int
		want int
	}{
		{nil, 0}, {[][]int{{}}, 0}, {[][]int{{0}}, 1}, {[][]int{{1}}, 1},
		{[][]int{{1, 0}, {0, 1}}, 3}, {[][]int{{1, 1}, {1, 0}}, 4}, {[][]int{{1, 1}, {1, 1}}, 4},
		{[][]int{{0, 0}, {0, 0}}, 1}, {[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, 9},
		{[][]int{{1, 0, 1}}, 3}, {[][]int{{1}, {0}, {1}}, 3},
		{[][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}, 5},
	} {
		before := clone(tt.grid)
		for i := 0; i < 2; i++ {
			if got := largestIsland(tt.grid); got != tt.want {
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

// 独立 BFS：枚举翻转点，每次重新计算最大连通块。
func bruteLargest(g [][]int) int {
	rows, cols := len(g), len(g[0])
	best := 0
	for flip := -1; flip < rows*cols; flip++ {
		if flip >= 0 && g[flip/cols][flip%cols] == 1 {
			continue
		}
		seen := make([]bool, rows*cols)
		land := func(p int) bool { return p == flip || g[p/cols][p%cols] == 1 }
		for start := 0; start < rows*cols; start++ {
			if seen[start] || !land(start) {
				continue
			}
			q := []int{start}
			seen[start] = true
			for head := 0; head < len(q); head++ {
				p := q[head]
				r, c := p/cols, p%cols
				for _, d := range [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
					nr, nc := r+d[0], c+d[1]
					if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
						continue
					}
					next := nr*cols + nc
					if !seen[next] && land(next) {
						seen[next] = true
						q = append(q, next)
					}
				}
			}
			if len(q) > best {
				best = len(q)
			}
		}
	}
	return best
}
func TestLargestIslandExhaustive(t *testing.T) {
	for rows := 1; rows <= 3; rows++ {
		for cols := 1; cols <= 3; cols++ {
			for mask := 0; mask < 1<<(rows*cols); mask++ {
				g := make([][]int, rows)
				for r := range g {
					g[r] = make([]int, cols)
					for c := range g[r] {
						g[r][c] = (mask >> (r*cols + c)) & 1
					}
				}
				want := bruteLargest(g)
				if got := largestIsland(g); got != want {
					t.Fatalf("grid %v: got %d want %d", g, got, want)
				}
			}
		}
	}
}
