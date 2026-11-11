package l295

import "testing"

func TestMedianFinder(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		want []float64
	}{
		{"example", []int{1, 2, 3}, []float64{1, 1.5, 2}}, {"negative", []int{-1, -2, -3, -4, -5}, []float64{-1, -1.5, -2, -2.5, -3}}, {"duplicates", []int{2, 2, 2, 2}, []float64{2, 2, 2, 2}}, {"mixed", []int{5, 1, 9, 3}, []float64{5, 3, 5, 4}}, {"large opposites", []int{100000, -100000}, []float64{100000, 0}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := Constructor()
			for i, x := range tt.nums {
				m.AddNum(x)
				if got := m.FindMedian(); got != tt.want[i] {
					t.Errorf("%s: after %d got %v want %v", t.Name(), i, got, tt.want[i])
				}
			}
		})
	}
}
