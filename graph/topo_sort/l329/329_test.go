package l329

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestLongestIncreasingPath(t *testing.T) {
	for _, tt := range []struct {
		matrix [][]int
		want   int
	}{
		{nil, 0}, {[][]int{{}}, 0}, {[][]int{{1}}, 1},
		{[][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}, 4},
		{[][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}, 4},
		{[][]int{{7, 7}, {7, 7}}, 1},
		{[][]int{{1, 2, 3, 4}}, 4}, {[][]int{{4}, {3}, {2}, {1}}, 4},
		{[][]int{{1, 2, 3}, {6, 5, 4}}, 6},
		{[][]int{{-3, -2}, {0, -1}}, 4},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 5},
	} {
		var before [][]int
		if tt.matrix != nil {
			before = make([][]int, len(tt.matrix))
			for r := range before {
				before[r] = append(make([]int, 0, len(tt.matrix[r])), tt.matrix[r]...)
			}
		}
		for call := 0; call < 2; call++ {
			if got := longestIncreasingPath(tt.matrix); got != tt.want {
				t.Fatalf("%v got %d want %d", tt.matrix, got, tt.want)
			}
		}
		if !reflect.DeepEqual(before, tt.matrix) {
			t.Fatal("input mutated")
		}
	}
}

// 独立的记忆化 DFS 对照，不使用入度或队列。
func memoLongest(g [][]int) int {
	rows, cols := len(g), len(g[0])
	memo := make([][]int, rows)
	for r := range memo {
		memo[r] = make([]int, cols)
	}
	var dfs func(int, int) int
	dfs = func(r, c int) int {
		if memo[r][c] != 0 {
			return memo[r][c]
		}
		best := 1
		for _, d := range [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && g[nr][nc] > g[r][c] {
				candidate := 1 + dfs(nr, nc)
				if candidate > best {
					best = candidate
				}
			}
		}
		memo[r][c] = best
		return best
	}
	answer := 0
	for r := range g {
		for c := range g[r] {
			if value := dfs(r, c); value > answer {
				answer = value
			}
		}
	}
	return answer
}
func TestMatrixAgainstMemoDFS(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	for trial := 0; trial < 300; trial++ {
		rows, cols := 1+rng.Intn(5), 1+rng.Intn(5)
		g := make([][]int, rows)
		for r := range g {
			g[r] = make([]int, cols)
			for c := range g[r] {
				g[r][c] = rng.Intn(11) - 5
			}
		}
		want := memoLongest(g)
		if got := longestIncreasingPath(g); got != want {
			t.Fatalf("%v got %d want %d", g, got, want)
		}
	}
}
