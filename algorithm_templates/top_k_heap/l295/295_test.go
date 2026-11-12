package l295

import (
	"math"
	"testing"
)

func TestMedianFinder(t *testing.T) {
	for _, tt := range []struct {
		name    string
		nums    []int
		medians []float64
	}{
		{"ascending", []int{1, 2, 3}, []float64{1, 1.5, 2}},
		{"negative", []int{-1, -2, -3}, []float64{-1, -1.5, -2}},
		{"mixed", []int{5, 15, 1, 3}, []float64{5, 10, 5, 4}},
		{"duplicates", []int{2, 2, 2}, []float64{2, 2, 2}},
		{"opposite extremes", []int{-100000, 100000, 0}, []float64{-100000, 0, 0}},
		{"single zero", []int{0}, []float64{0}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			stream := Constructor()
			for i, x := range tt.nums {
				stream.AddNum(x)
				got := stream.FindMedian()
				if math.IsNaN(got) || math.Abs(got-tt.medians[i]) > 1e-9 {
					t.Errorf("after adding %d got %v want %v", x, got, tt.medians[i])
				}
			}
		})
	}
}
