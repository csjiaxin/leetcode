package l526

import (
	"strconv"
	"testing"
)

func TestCountArrangement(t *testing.T) {
	for _, tt := range []struct {
		n, want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 8},
		{5, 10},
		{6, 36},
		{7, 41},
		{8, 132},
		{10, 700},
		{15, 24679},
	} {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			if got := countArrangement(tt.n); got != tt.want {
				t.Errorf("countArrangement(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
