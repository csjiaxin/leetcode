package l2050

import (
	"reflect"
	"testing"
)

func TestMinimumTimeBoundaries(t *testing.T) {
	for _, tt := range []struct {
		n     int
		edges [][]int
		times []int
		want  int
	}{
		{0, nil, nil, 0}, {1, [][]int{{1, 1}}, []int{1}, -1},
		{3, [][]int{{1, 2}, {2, 1}}, []int{1, 1, 9}, -1},
		{3, [][]int{{1, 2}, {1, 2}, {2, 3}}, []int{2, 3, 4}, 9},
		{4, [][]int{{1, 3}, {2, 3}, {3, 4}}, []int{7, 10, 2, 3}, 15},
	} {
		before := make([][]int, len(tt.edges))
		for i := range before {
			before[i] = append([]int(nil), tt.edges[i]...)
		}
		times := append([]int(nil), tt.times...)
		for call := 0; call < 2; call++ {
			if got := minimumTime(tt.n, tt.edges, tt.times); got != tt.want {
				t.Fatalf("got %d want %d", got, tt.want)
			}
		}
		if !reflect.DeepEqual(tt.times, times) {
			t.Fatal("times mutated")
		}
		for i := range before {
			if !reflect.DeepEqual(tt.edges[i], before[i]) {
				t.Fatal("edges mutated")
			}
		}
	}
	// 最大合法耗时与长链，确保累计而不是计算拓扑层数。
	const n = 50000
	times := make([]int, n)
	edges := make([][]int, 0, n-1)
	for i := range times {
		times[i] = 100000
		if i > 0 {
			edges = append(edges, []int{i, i + 1})
		}
	}
	if got := minimumTime(n, edges, times); int64(got) != 5000000000 {
		t.Fatalf("large chain got %d", got)
	}
}
