package l375

import (
	"fmt"
	"testing"
)

func TestGetMoneyAmount(t *testing.T) {
	for _, tt := range []struct{ n, want int }{
		{1, 0}, {2, 1}, {3, 2}, {4, 4}, {5, 6}, {6, 8}, {7, 10}, {10, 16}, {20, 49}, {100, 400}, {200, 952},
	} {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			if got := getMoneyAmount(tt.n); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
