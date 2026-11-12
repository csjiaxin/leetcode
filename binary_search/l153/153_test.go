package l153

import "testing"

func TestFindMin(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		want int
	}{
		{"example", []int{3, 4, 5, 1, 2}, 1}, {"sorted", []int{1, 2, 3}, 1},
		{"single", []int{7}, 7}, {"two", []int{2, 1}, 1},
		{"last", []int{2, 3, 4, 1}, 1}, {"negative", []int{0, 2, -4, -1}, -4},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.nums); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
