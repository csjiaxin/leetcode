package l1035

import "testing"

func TestMaxUncrossedLines(t *testing.T) {
	tests := []struct {
		name         string
		nums1, nums2 []int
		want         int
	}{
		{"official_crossing", []int{1, 4, 2}, []int{1, 2, 4}, 2},
		{"official_repeated", []int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}, 3},
		{"official_multiple_choices", []int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}, 2},
		{"empty_both", nil, nil, 0},
		{"empty_first", []int{}, []int{1, 2}, 0},
		{"empty_second", []int{1, 2}, []int{}, 0},
		{"equal", []int{1, 2, 3}, []int{1, 2, 3}, 3},
		{"disjoint", []int{1, 2}, []int{3, 4}, 0},
		{"repeated_values", []int{2, 2, 2}, []int{2, 2}, 2},
		{"order_matters", []int{1, 2, 3}, []int{3, 2, 1}, 1},
		{"alternating", []int{1, 2, 1, 2}, []int{2, 1, 2, 1}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxUncrossedLines(tt.nums1, tt.nums2); got != tt.want {
				t.Fatalf("maxUncrossedLines(%v, %v) = %d, want %d", tt.nums1, tt.nums2, got, tt.want)
			}
			if got := maxUncrossedLines(tt.nums2, tt.nums1); got != tt.want {
				t.Fatalf("swapped maxUncrossedLines(%v, %v) = %d, want %d", tt.nums2, tt.nums1, got, tt.want)
			}
		})
	}
}
