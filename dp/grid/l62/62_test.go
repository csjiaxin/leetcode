package l62

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m, n int
		want int
	}{
		{name: "example one", m: 3, n: 7, want: 28},
		{name: "example two", m: 3, n: 2, want: 3},
		{name: "single cell", m: 1, n: 1, want: 1},
		{name: "single row", m: 1, n: 100, want: 1},
		{name: "single column", m: 100, n: 1, want: 1},
		{name: "two by two", m: 2, n: 2, want: 2},
		{name: "three by three", m: 3, n: 3, want: 6},
		{name: "wide grid", m: 2, n: 100, want: 100},
		{name: "tall grid", m: 100, n: 2, want: 100},
		{name: "transposed example", m: 7, n: 3, want: 28},
		{name: "medium square", m: 10, n: 10, want: 48620},
		{name: "large valid result", m: 17, n: 17, want: 601080390},
		{name: "near result limit", m: 17, n: 18, want: 1166803110},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.m, tt.n); got != tt.want {
				t.Errorf("uniquePaths(%d, %d) = %d, want %d", tt.m, tt.n, got, tt.want)
			}
		})
	}
}
