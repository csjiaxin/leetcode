package l1000

import "testing"

func TestMergeStones(t *testing.T) {
	tests := []struct {
		name    string
		stones  []int
		k, want int
	}{
		{"example one", []int{3, 2, 4, 1}, 2, 20},
		{"example two impossible", []int{3, 2, 4, 1}, 3, -1},
		{"example three", []int{3, 5, 1, 2, 6}, 3, 25},
		{"single pile", []int{5}, 2, 0},
		{"k greater than pile count", []int{1, 2}, 3, -1},
		{"one merge", []int{1, 2, 3}, 3, 6},
		{"binary merges", []int{1, 2, 3}, 2, 9},
		{"two ternary merges", []int{1, 2, 3, 4, 5}, 3, 21},
		{"seven equal piles", []int{1, 1, 1, 1, 1, 1, 1}, 3, 13},
		{"four way merging", []int{1, 1, 1, 1, 1, 1, 1}, 4, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeStones(tt.stones, tt.k); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}

func TestMergeStonesMaximumInput(t *testing.T) {
	stones := make([]int, 30)
	for i := range stones {
		stones[i] = 100
	}
	for _, tt := range []struct {
		name    string
		k, want int
	}{
		{"binary", 2, 14800}, {"all at once", 30, 3000}, {"impossible", 3, -1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeStones(stones, tt.k); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
