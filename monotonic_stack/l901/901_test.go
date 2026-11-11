package l901

import "testing"

func TestStockSpanner(t *testing.T) {
	for _, tt := range []struct {
		name         string
		prices, want []int
	}{
		{"example", []int{100, 80, 60, 70, 60, 75, 85}, []int{1, 1, 1, 2, 1, 4, 6}},
		{"increasing", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"decreasing", []int{4, 3, 2, 1}, []int{1, 1, 1, 1}},
		{"equal prices", []int{5, 5, 5}, []int{1, 2, 3}},
		{"merge spans", []int{10, 5, 5, 8, 10}, []int{1, 1, 2, 3, 5}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			s := Constructor()
			for i, p := range tt.prices {
				if got := s.Next(p); got != tt.want[i] {
					t.Errorf("%s: day %d Next(%d) = %d, want %d", t.Name(), i, p, got, tt.want[i])
				}
			}
		})
	}
}

func TestStockSpannerIndependentInstances(t *testing.T) {
	a, b := Constructor(), Constructor()
	for _, tt := range []struct {
		s           *StockSpanner
		price, want int
	}{{&a, 100, 1}, {&b, 1, 1}, {&a, 50, 1}, {&b, 2, 2}} {
		if got := tt.s.Next(tt.price); got != tt.want {
			t.Errorf("%s: Next(%d) = %d, want %d", t.Name(), tt.price, got, tt.want)
		}
	}
}
