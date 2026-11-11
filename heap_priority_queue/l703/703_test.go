package l703

import "testing"

func TestKthLargest(t *testing.T) {
	for _, tt := range []struct {
		name                 string
		k                    int
		initial, added, want []int
	}{
		{"example", 3, []int{4, 5, 8, 2}, []int{3, 5, 10, 9, 4}, []int{4, 5, 5, 8, 8}},
		{"empty", 1, nil, []int{-3, -2, -4, 0, 4}, []int{-3, -2, -2, 0, 4}},
		{"duplicates", 2, []int{1}, []int{1, 1, 2, 3}, []int{1, 1, 1, 2}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			s := Constructor(tt.k, tt.initial)
			for i, x := range tt.added {
				if got := s.Add(x); got != tt.want[i] {
					t.Errorf("%s: Add(%d) got %d want %d", t.Name(), x, got, tt.want[i])
				}
			}
		})
	}
}
