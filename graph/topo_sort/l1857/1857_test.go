package l1857

import (
	"reflect"
	"testing"
)

func TestLargestPathValue(t *testing.T) {
	for _, tt := range []struct {
		name, colors string
		edges        [][]int
		want         int
	}{
		{"example", "abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}, 3},
		{"self cycle", "a", [][]int{{0, 0}}, -1},
		{"disconnected cycle", "aaaa", [][]int{{0, 1}, {2, 3}, {3, 2}}, -1},
		{"diamond not branch sum", "aaaa", [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}, 3},
		{"fork", "aaa", [][]int{{0, 1}, {0, 2}}, 2},
		{"distinct colors", "abc", [][]int{{0, 1}, {1, 2}}, 1},
		{"disconnected", "aab", [][]int{{0, 1}}, 2},
		{"isolated", "abc", nil, 1}, {"single", "z", nil, 1}, {"empty", "", nil, 0},
		{"duplicate edges", "aa", [][]int{{0, 1}, {0, 1}}, 2},
		{"multiple color states", "abbba", [][]int{{0, 1}, {1, 2}, {2, 4}, {0, 3}, {3, 4}}, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var before [][]int
			if tt.edges != nil {
				before = make([][]int, len(tt.edges))
				for i := range before {
					before[i] = append([]int(nil), tt.edges[i]...)
				}
			}
			for call := 0; call < 2; call++ {
				if got := largestPathValue(tt.colors, tt.edges); got != tt.want {
					t.Fatalf("got %d want %d", got, tt.want)
				}
			}
			if !reflect.DeepEqual(before, tt.edges) {
				t.Fatal("input mutated")
			}
		})
	}
}
