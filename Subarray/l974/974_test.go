package l974

import "testing"

func TestSubarraysDivByK(t *testing.T) {
	for _, tt := range []struct {
		name    string
		nums    []int
		k, want int
	}{
		{"example", []int{4, 5, 0, -2, -3, 1}, 5, 7}, {"missing", []int{5}, 9, 0},
		{"negative remainder", []int{-1, 2, 9}, 2, 2}, {"negative multiple", []int{-6}, 3, 1},
		{"zeros", []int{0, 0, 0}, 2, 6}, {"same remainder", []int{-1, 1}, 2, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysDivByK(append([]int(nil), tt.nums...), tt.k); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
