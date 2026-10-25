package main

import "testing"

func TestIsBipartite(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  bool
	}{
		{name: "empty graph", graph: nil, want: true},
		{name: "isolated vertices", graph: [][]int{{}, {}, {}}, want: true},
		{name: "single edge", graph: [][]int{{1}, {0}}, want: true},
		{name: "even cycle", graph: [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}, want: true},
		{name: "odd cycle", graph: [][]int{{1, 2}, {0, 2}, {0, 1}}, want: false},
		{name: "example with diagonal", graph: [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}, want: false},
		{name: "disconnected edges", graph: [][]int{{1}, {0}, {3}, {2}, {}}, want: true},
		{name: "disconnected odd cycle", graph: [][]int{{1}, {0}, {3, 4}, {2, 4}, {2, 3}}, want: false},
		{name: "self loop", graph: [][]int{{0}}, want: false},
		{name: "branching tree", graph: [][]int{{1, 2}, {0, 3}, {0, 4}, {1}, {2}}, want: true},
		{name: "queue traversal regression", graph: [][]int{{1}, {0, 2}, {1, 3}, {2}}, want: true},
		{name: "reordered path", graph: [][]int{{3}, {2, 4}, {1}, {0, 4}, {1, 3}}, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBipartite(tt.graph); got != tt.want {
				t.Errorf("isBipartite(%v) = %v, want %v", tt.graph, got, tt.want)
			}
		})
	}
}
