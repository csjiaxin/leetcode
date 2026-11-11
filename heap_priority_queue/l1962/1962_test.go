package l1962

import "testing"

func TestMinStoneSum(t *testing.T) {
	for _, tt := range []struct {
		name    string
		piles   []int
		k, want int
	}{
		{"example one", []int{5, 4, 9}, 2, 12}, {"example two", []int{4, 3, 6, 7}, 3, 12}, {"ones", []int{1, 1}, 10, 2}, {"odd rounding", []int{5}, 1, 3}, {"even rounding", []int{6}, 1, 3}, {"repeat", []int{10000}, 14, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := minStoneSum(tt.piles, tt.k); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
