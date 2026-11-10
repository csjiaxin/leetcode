package l1690

import "testing"

func TestStoneGameVII(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{name: "example one", stones: []int{5, 3, 1, 4, 2}, want: 6},
		{name: "example two", stones: []int{7, 90, 5, 1, 100, 10, 10, 2}, want: 122},
		{name: "two equal stones", stones: []int{4, 4}, want: 4},
		{name: "two stones remove left", stones: []int{1, 9}, want: 9},
		{name: "two stones remove right", stones: []int{9, 1}, want: 9},
		{name: "minimum weights", stones: []int{1, 1}, want: 1},
		{name: "maximum weights two stones", stones: []int{1000, 1000}, want: 1000},
		{name: "three equal stones", stones: []int{3, 3, 3}, want: 3},
		{name: "large middle stone", stones: []int{1, 1000, 1}, want: 1},
		{name: "large left endpoint", stones: []int{1000, 2, 1}, want: 2},
		{name: "large right endpoint", stones: []int{1, 2, 1000}, want: 2},
		{name: "small middle stone", stones: []int{9, 1, 8}, want: 1},
		{name: "increasing three stones", stones: []int{1, 2, 3}, want: 2},
		{name: "decreasing three stones", stones: []int{3, 2, 1}, want: 2},
		{name: "four equal stones", stones: []int{5, 5, 5, 5}, want: 10},
		{name: "five equal stones", stones: []int{5, 5, 5, 5, 5}, want: 10},
		{name: "reversed first example", stones: []int{2, 4, 1, 3, 5}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stones := append([]int(nil), tt.stones...)
			if got := stoneGameVII(stones); got != tt.want {
				t.Errorf("%s: stoneGameVII(%v) = %d, want %d", t.Name(), tt.stones, got, tt.want)
			}
		})
	}
}

func TestStoneGameVIILargeInput(t *testing.T) {
	tests := []struct {
		name  string
		count int
		value int
		want  int
	}{
		{name: "maximum length minimum weights", count: 1000, value: 1, want: 500},
		{name: "maximum length maximum weights", count: 1000, value: 1000, want: 500000},
		{name: "large odd length", count: 999, value: 1000, want: 499000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stones := make([]int, tt.count)
			for i := range stones {
				stones[i] = tt.value
			}
			if got := stoneGameVII(stones); got != tt.want {
				t.Errorf("%s: stoneGameVII(%d stones of weight %d) = %d, want %d", t.Name(), tt.count, tt.value, got, tt.want)
			}
		})
	}
}
