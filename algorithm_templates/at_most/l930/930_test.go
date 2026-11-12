package l930

import "testing"

func TestNumSubarraysWithSum(t *testing.T) {
	for _, tt := range []struct {
		name       string
		nums       []int
		goal, want int
	}{
		{"example", []int{1, 0, 1, 0, 1}, 2, 4}, {"zeros", []int{0, 0, 0, 0, 0}, 0, 15},
		{"zero goal", []int{1, 0, 0, 1}, 0, 3}, {"single one", []int{1}, 1, 1},
		{"single zero", []int{0}, 0, 1}, {"missing", []int{0, 0}, 1, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarraysWithSum(append([]int(nil), tt.nums...), tt.goal); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
