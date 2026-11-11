package l556

import "testing"

func TestNextGreaterElement(t *testing.T) {
	for _, tt := range []struct {
		name    string
		n, want int
	}{
		{"example one", 12, 21},
		{"example two", 21, -1},
		{"single", 7, -1},
		{"all equal", 111, -1},
		{"duplicates", 115, 151},
		{"zero digits", 1020, 1200},
		{"reverse suffix", 12443322, 13222344},
		{"largest valid result", 2147483476, 2147483647},
		{"result overflows", 1999999999, -1},
		{"max int32 input", 2147483647, -1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.n); got != tt.want {
				t.Errorf("nextGreaterElement(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
