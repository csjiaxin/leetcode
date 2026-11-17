package template

import (
	"reflect"
	"testing"
)

func TestKahnReadinessAndInputOwnership(t *testing.T) {
	graph := [][]int{{2, 2}, {2}, {3}, nil}
	indegree := []int{0, 0, 3, 1}
	backing := []int{0, 1, 99, 99, 99}
	for call := 0; call < 2; call++ {
		contributions := make([]int, 4)
		popped := make([]bool, 4)
		processed := Kahn(indegree, backing[:2], func(u int, yield func(int)) {
			for _, v := range graph[u] {
				yield(v)
			}
		}, func(u int) {
			if popped[u] || contributions[u] != indegree[u] {
				t.Fatalf("node %d popped before all dependencies or more than once", u)
			}
			popped[u] = true
		}, func(u, v int) {
			if !popped[u] || popped[v] {
				t.Fatal("invalid edge callback timing")
			}
			contributions[v]++
		})
		if processed != 4 {
			t.Fatalf("processed %d nodes", processed)
		}
	}
	if !reflect.DeepEqual(indegree, []int{0, 0, 3, 1}) || !reflect.DeepEqual(backing, []int{0, 1, 99, 99, 99}) {
		t.Fatal("caller input or backing array was modified")
	}
}

func TestKahnCycleAndExplicitSeeds(t *testing.T) {
	graph := [][]int{nil, {2}, {1}, nil}
	degree := []int{0, 1, 1, 0}
	next := func(u int, yield func(int)) {
		for _, v := range graph[u] {
			yield(v)
		}
	}
	if got := Kahn(degree, ZeroIndegree(degree), next, nil, nil); got != 2 {
		t.Fatalf("cycle should leave two blocked nodes, processed %d", got)
	}
	if got := Kahn(degree, []int{3}, next, nil, nil); got != 1 {
		t.Fatalf("unseeded zero-indegree node should stay unprocessed, got %d", got)
	}
	if got := Kahn(nil, nil, next, nil, nil); got != 0 {
		t.Fatal("empty graph is not empty")
	}
}
