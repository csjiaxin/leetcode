package l264

import (
	"fmt"
	"testing"
)

func TestNthUglyNumber(t *testing.T) {
	for _, tt := range []struct{ n, want int }{{1, 1}, {2, 2}, {7, 8}, {10, 12}, {15, 24}, {1690, 2123366400}} {
		t.Run(fmt.Sprint(tt.n), func(t *testing.T) {
			if got := nthUglyNumber(tt.n); got != tt.want {
				t.Errorf("%s: got %d want %d", t.Name(), got, tt.want)
			}
		})
	}
}
