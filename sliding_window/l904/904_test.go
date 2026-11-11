package l904

import (
	"slices"
	"testing"
)

func TestTotalFruit(t *testing.T) {
	for _, tt := range []struct {
		name   string
		fruits []int
		want   int
	}{
		{"example one", []int{1, 2, 1}, 3},
		{"example two", []int{0, 1, 2, 2}, 3},
		{"example three", []int{1, 2, 3, 2, 2}, 4},
		{"multiple shrinking steps", []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}, 5},
		{"single", []int{0}, 1},
		{"one type", []int{2, 2, 2}, 3},
		{"two alternating types", []int{0, 1, 0, 1, 0, 1}, 6},
		{"all distinct", []int{0, 1, 2, 3}, 2},
		{"best at start", []int{1, 1, 2, 2, 3, 4}, 4},
		{"best at end", []int{0, 1, 2, 2, 1, 1}, 5},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalFruit(slices.Clone(tt.fruits)); got != tt.want {
				t.Errorf("totalFruit(%v) = %d, want %d", tt.fruits, got, tt.want)
			}
		})
	}
}

func TestTotalFruitMaximum(t *testing.T) {
	const n = 100000
	fruits := make([]int, n)
	for i := range fruits {
		fruits[i] = n - 1
	}
	if got := totalFruit(fruits); got != n {
		t.Errorf("maximum equal input: got %d, want %d", got, n)
	}
}
