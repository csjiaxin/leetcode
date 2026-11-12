package l210

import "testing"

func TestFindOrder(t *testing.T) {
	for _, tt := range []struct {
		name          string
		n             int
		prerequisites [][]int
		possible      bool
	}{
		{"edge", 2, [][]int{{1, 0}}, true}, {"cycle", 2, [][]int{{1, 0}, {0, 1}}, false},
		{"isolated", 4, nil, true}, {"diamond", 4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, true},
		{"cycle after root", 4, [][]int{{1, 0}, {2, 1}, {1, 2}, {3, 2}}, false},
		{"single", 1, nil, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.prerequisites))
			for i := range input {
				input[i] = append([]int(nil), tt.prerequisites[i]...)
			}
			got := findOrder(tt.n, input)
			if !tt.possible {
				if len(got) != 0 {
					t.Errorf("expected empty, got %v", got)
				}
				return
			}
			if len(got) != tt.n {
				t.Fatalf("got %v want %d courses", got, tt.n)
			}
			pos := make([]int, tt.n)
			seen := make([]bool, tt.n)
			for i, v := range got {
				if v < 0 || v >= tt.n || seen[v] {
					t.Fatalf("invalid permutation %v", got)
				}
				seen[v] = true
				pos[v] = i
			}
			for _, e := range tt.prerequisites {
				if pos[e[1]] >= pos[e[0]] {
					t.Errorf("order %v violates %v", got, e)
				}
			}
		})
	}
}
