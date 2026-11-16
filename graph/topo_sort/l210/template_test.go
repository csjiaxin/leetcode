package l210

import (
	"reflect"
	"testing"
)

func TestOrderBoundaries(t *testing.T) {
	for _, tt := range []struct {
		n        int
		edges    [][]int
		possible bool
	}{
		{0, nil, true}, {1, [][]int{{0, 0}}, false}, {4, [][]int{{1, 0}, {0, 1}}, false}, {3, [][]int{{1, 0}, {1, 0}, {2, 1}}, true},
	} {
		before := make([][]int, len(tt.edges))
		for i := range before {
			before[i] = append([]int(nil), tt.edges[i]...)
		}
		for call := 0; call < 2; call++ {
			got := findOrder(tt.n, tt.edges)
			if !tt.possible {
				if len(got) != 0 {
					t.Fatalf("cycle: %v", got)
				}
				continue
			}
			if len(got) != tt.n {
				t.Fatalf("got %v", got)
			}
			pos := make([]int, tt.n)
			seen := make([]bool, tt.n)
			for i, u := range got {
				if u < 0 || u >= tt.n || seen[u] {
					t.Fatalf("invalid %v", got)
				}
				seen[u] = true
				pos[u] = i
			}
			for _, e := range tt.edges {
				if pos[e[1]] >= pos[e[0]] {
					t.Fatalf("order violates %v", e)
				}
			}
		}
		for i := range before {
			if !reflect.DeepEqual(tt.edges[i], before[i]) {
				t.Fatal("input mutated")
			}
		}
	}
}

func TestKahnLifecycle(t *testing.T) {
	graph := [][]int{{2}, {2}, {3}, nil}
	degree := []int{0, 0, 2, 1}
	// 额外容量用于检测 append 是否污染调用方底层数组。
	backing := []int{0, 1, 99, 99, 99}
	seeds := backing[:2]
	for call := 0; call < 2; call++ {
		contributions := make([]int, 4)
		popped := make([]bool, 4)
		got := Kahn(degree, seeds, func(u int, yield func(int)) {
			for _, v := range graph[u] {
				yield(v)
			}
		},
			func(u int) {
				if contributions[u] != degree[u] {
					t.Fatalf("node %d popped too early", u)
				}
				popped[u] = true
			},
			func(u, v int) {
				if !popped[u] || popped[v] {
					t.Fatal("invalid relax timing")
				}
				contributions[v]++
			},
		)
		if got != 4 {
			t.Fatalf("processed %d", got)
		}
	}
	if !reflect.DeepEqual(degree, []int{0, 0, 2, 1}) || !reflect.DeepEqual(backing, []int{0, 1, 99, 99, 99}) {
		t.Fatal("template mutated inputs")
	}
	if got := Kahn(degree, seeds, func(u int, y func(int)) {
		for _, v := range graph[u] {
			y(v)
		}
	}, nil, nil); got != 4 {
		t.Fatal(got)
	}
}
