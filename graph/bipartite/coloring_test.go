package main

import "testing"

func TestBipartiteColoring(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  bool
	}{
		{"empty", nil, true},
		{"isolated", [][]int{{}, {}, {}}, true},
		{"disconnected", [][]int{{1}, {0}, {3}, {2}, {}}, true},
		{"parallel edges", [][]int{{1, 1}, {0, 0}}, true},
		{"self loop", [][]int{{0}}, false},
		{"odd cycle", [][]int{{1, 2}, {0, 2}, {0, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkColoring(t, tt.graph, tt.want)
		})
	}
}

func checkColoring(t *testing.T, graph [][]int, want bool) {
	t.Helper()
	color, ok := bipartiteColoring(graph)
	if ok != want {
		t.Fatalf("bipartiteColoring(%v) = (%v, %v), want success %v", graph, color, ok, want)
	}
	if !ok {
		if color != nil {
			t.Fatalf("failure returned partial coloring: %v", color)
		}
		return
	}
	if len(color) != len(graph) {
		t.Fatalf("got %d colors for %d vertices", len(color), len(graph))
	}
	for u, neighbors := range graph {
		if color[u] != 1 && color[u] != -1 {
			t.Fatalf("vertex %d has invalid color %d", u, color[u])
		}
		for _, v := range neighbors {
			if color[u] == color[v] {
				t.Fatalf("edge %d--%d has equal colors in %v", u, v, color)
			}
		}
	}
}

// Enumerate every simple undirected graph with at most five vertices and
// compare BFS with independently enumerated assignments to the two groups.
func TestBipartiteColoringExhaustive(t *testing.T) {
	for n := 0; n <= 5; n++ {
		edgeCount := n * (n - 1) / 2
		for mask := 0; mask < 1<<edgeCount; mask++ {
			graph := make([][]int, n)
			bit := 0
			for u := 0; u < n; u++ {
				for v := u + 1; v < n; v++ {
					if mask&(1<<bit) != 0 {
						graph[u] = append(graph[u], v)
						graph[v] = append(graph[v], u)
					}
					bit++
				}
			}
			want := false
			for groups := 0; groups < 1<<n; groups++ {
				valid := true
				for u, neighbors := range graph {
					for _, v := range neighbors {
						if (groups>>u)&1 == (groups>>v)&1 {
							valid = false
						}
					}
				}
				if valid {
					want = true
					break
				}
			}
			checkColoring(t, graph, want)
		}
	}
}
