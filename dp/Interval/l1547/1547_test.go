package l1547

import "testing"

func TestMinCost(t *testing.T) {
	tests := []struct {
		name string
		n    int
		cuts []int
		want int
	}{
		{"example one", 7, []int{1, 3, 4, 5}, 16},
		{"example two unsorted", 9, []int{5, 6, 1, 4, 2}, 22},
		{"single cut", 5, []int{2}, 5},
		{"two cuts", 10, []int{3, 7}, 17},
		{"minimum length", 2, []int{1}, 2},
		{"maximum length", 1000000, []int{500000}, 1000000},
		{"balanced cuts", 8, []int{6, 2, 4}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := append([]int(nil), tt.cuts...)
			if got := minCost(tt.n, tt.cuts); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
			for i := range original {
				if original[i] != tt.cuts[i] {
					t.Fatalf("%s: cuts mutated", t.Name())
				}
			}
		})
	}
}

func TestMinCostMaximumCuts(t *testing.T) {
	// 101 unit segments: optimal leaf depths are six or seven.
	cuts := make([]int, 100)
	for i := range cuts {
		cuts[i] = i + 1
	}
	if got := minCost(101, cuts); got != 680 {
		t.Errorf("%s: got %d, want 680", t.Name(), got)
	}
}
