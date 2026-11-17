package l525

import "testing"

func TestFindMaxLength(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		want int
	}{
		{"example", []int{0, 1}, 2}, {"earliest prefix", []int{0, 1, 0, 1}, 4},
		{"odd length", []int{0, 1, 0}, 2}, {"all zeros", []int{0, 0, 0}, 0},
		{"single", []int{1}, 0}, {"interior", []int{0, 0, 1, 0, 0, 0, 1, 1}, 6},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLength(append([]int(nil), tt.nums...)); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
