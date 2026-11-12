package l888

import (
	"slices"
	"testing"
)

func TestFairCandySwap(t *testing.T) {
	for _, tt := range []struct {
		name string
		a, b []int
	}{
		{"example", []int{1, 1}, []int{2, 2}}, {"different sizes", []int{1, 2}, []int{2, 3}},
		{"single", []int{2}, []int{1, 3}}, {"multiple answers", []int{1, 2, 5}, []int{2, 4}},
		{"equal totals", []int{3, 1}, []int{1, 3}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := fairCandySwap(slices.Clone(tt.a), slices.Clone(tt.b))
			if len(got) != 2 {
				t.Fatalf("invalid swap %v", got)
			}
			if !slices.Contains(tt.a, got[0]) || !slices.Contains(tt.b, got[1]) {
				t.Fatalf("swap not in inputs: %v", got)
			}
			a, b := 0, 0
			for _, v := range tt.a {
				a += v
			}
			for _, v := range tt.b {
				b += v
			}
			if a-got[0]+got[1] != b-got[1]+got[0] {
				t.Errorf("swap %v does not equalize totals", got)
			}
		})
	}
}
