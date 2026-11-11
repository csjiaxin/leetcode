package l1046

import "testing"

func TestLastStoneWeight(t *testing.T) {
	for _, tt := range []struct {
		name   string
		stones []int
		want   int
	}{
		{"example", []int{2, 7, 4, 1, 8, 1}, 1}, {"single", []int{1}, 1}, {"equal", []int{4, 4}, 0}, {"odd count", []int{2, 2, 2}, 2}, {"all smashed", []int{1, 1, 2}, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeight(tt.stones); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
