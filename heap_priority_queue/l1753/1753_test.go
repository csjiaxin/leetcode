package l1753

import "testing"

func TestMaximumScore(t *testing.T) {
	for _, tt := range []struct {
		name          string
		a, b, c, want int
	}{
		{"example one", 2, 4, 6, 6}, {"example two", 4, 4, 6, 7}, {"example three", 1, 8, 8, 8}, {"dominant pile", 1, 1, 100, 2}, {"equal odd total", 3, 3, 3, 4}, {"minimum", 1, 1, 1, 1}, {"maximum", 100000, 100000, 100000, 150000},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.a, tt.b, tt.c); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
