package l1049

import "testing"

func TestLastStoneWeightII(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{name: "example one", stones: []int{2, 7, 4, 1, 8, 1}, want: 1},
		{name: "example two", stones: []int{31, 26, 33, 21, 40}, want: 5},
		{name: "single stone", stones: []int{7}, want: 7},
		{name: "equal stones", stones: []int{4, 4}, want: 0},
		{name: "unequal stones", stones: []int{3, 8}, want: 5},
		{name: "exact partition", stones: []int{1, 2, 3, 6}, want: 0},
		{name: "even total without exact partition", stones: []int{2, 2, 2}, want: 2},
		{name: "cannot reuse a stone", stones: []int{1, 4}, want: 3},
		{name: "dominant stone", stones: []int{1, 2, 10}, want: 7},
		{name: "odd total", stones: []int{1, 1, 1}, want: 1},
		{name: "greedy smashing is not optimal", stones: []int{2, 3, 4, 5, 6}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stones := append([]int(nil), tt.stones...)
			if got := lastStoneWeightII(stones); got != tt.want {
				t.Errorf("lastStoneWeightII(%v) = %d, want %d", tt.stones, got, tt.want)
			}
			for i, stone := range stones {
				if stone != tt.stones[i] {
					t.Fatalf("input modified: got %v, want %v", stones, tt.stones)
				}
			}
		})
	}
}

func TestLastStoneWeightIIMaximumInput(t *testing.T) {
	for _, count := range []int{29, 30} {
		stones := make([]int, count)
		for i := range stones {
			stones[i] = 100
		}
		want := (count % 2) * 100
		if got := lastStoneWeightII(stones); got != want {
			t.Errorf("lastStoneWeightII(%d stones of weight 100) = %d, want %d", count, got, want)
		}
	}
}
