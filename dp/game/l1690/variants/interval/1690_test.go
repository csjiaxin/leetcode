package l1690

import "testing"

func TestStoneGameVII(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{"example one", []int{5, 3, 1, 4, 2}, 6},
		{"example two", []int{7, 90, 5, 1, 100, 10, 10, 2}, 122},
		{"two stones", []int{1, 9}, 9},
		{"large middle", []int{1, 1000, 1}, 1},
		{"large endpoint", []int{1000, 2, 1}, 2},
		{"equal stones", []int{5, 5, 5, 5, 5}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameVII(tt.stones); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}

func TestStoneGameVIIMaximum(t *testing.T) {
	stones := make([]int, 1000)
	for i := range stones {
		stones[i] = 1000
	}
	if got := stoneGameVII(stones); got != 500000 {
		t.Errorf("%s: got %d, want 500000", t.Name(), got)
	}
}
