package l2558

import "testing"

func TestPickGifts(t *testing.T) {
	for _, tt := range []struct {
		name  string
		gifts []int
		k     int
		want  int64
	}{
		{"example", []int{25, 64, 9, 4, 100}, 4, 29}, {"ones", []int{1, 1, 1, 1}, 4, 4}, {"floor sqrt", []int{15}, 1, 3}, {"perfect square", []int{1000000000}, 1, 31622}, {"large total", []int{1000000000, 1000000000, 1000000000, 1000000000}, 1, 3000031622}, {"repeated", []int{81}, 3, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := pickGifts(tt.gifts, tt.k); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
