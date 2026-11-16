package l1462

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestCheckIfPrerequisite(t *testing.T) {
	for _, tt := range []struct {
		n              int
		edges, queries [][]int
		want           []bool
	}{
		{2, [][]int{{1, 0}}, [][]int{{0, 1}, {1, 0}}, []bool{false, true}},
		{2, nil, [][]int{{0, 1}, {1, 0}}, []bool{false, false}},
		{4, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}, [][]int{{0, 3}, {1, 2}, {2, 1}, {3, 0}, {1, 3}}, []bool{true, false, false, false, true}},
		{5, [][]int{{0, 1}, {1, 2}, {3, 4}}, [][]int{{0, 2}, {0, 4}, {3, 4}, {4, 3}}, []bool{true, false, true, false}},
		{0, nil, nil, []bool{}},
	} {
		edges := clone(tt.edges)
		queries := clone(tt.queries)
		for call := 0; call < 2; call++ {
			if got := checkIfPrerequisite(tt.n, tt.edges, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v want %v", got, tt.want)
			}
		}
		if !reflect.DeepEqual(edges, tt.edges) || !reflect.DeepEqual(queries, tt.queries) {
			t.Fatal("input mutated")
		}
	}
}
func clone(g [][]int) [][]int {
	if g == nil {
		return nil
	}
	out := make([][]int, len(g))
	for i := range g {
		out[i] = append([]int(nil), g[i]...)
	}
	return out
}

func TestPrerequisitesAgainstFloyd(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	for trial := 0; trial < 200; trial++ {
		n := 1 + rng.Intn(9)
		perm := rng.Perm(n)
		edges := [][]int{}
		reach := make([][]bool, n)
		for i := range reach {
			reach[i] = make([]bool, n)
		}
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if rng.Intn(3) == 0 {
					u, v := perm[i], perm[j]
					edges = append(edges, []int{u, v})
					reach[u][v] = true
				}
			}
		}
		for k := 0; k < n; k++ {
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					reach[i][j] = reach[i][j] || (reach[i][k] && reach[k][j])
				}
			}
		}
		queries := [][]int{}
		want := []bool{}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				queries = append(queries, []int{i, j})
				want = append(want, reach[i][j])
			}
		}
		if got := checkIfPrerequisite(n, edges, queries); !reflect.DeepEqual(got, want) {
			t.Fatalf("trial %d edges %v got %v want %v", trial, edges, got, want)
		}
	}
}
