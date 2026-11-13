package l2851

import "testing"

func TestNumberOfWays(t *testing.T) {
	tests := []struct {
		name, s, target string
		k               int64
		want            int
	}{
		{"one rotation", "ab", "ba", 1, 1},
		{"proper suffix only", "ab", "ab", 1, 0},
		{"two operations", "ab", "ab", 2, 1},
		{"example", "abcd", "cdab", 2, 2},
		{"identical letters", "aaaa", "aaaa", 2, 9},
		{"multiple equal rotations", "ababab", "ababab", 1, 2},
		{"not a rotation", "abcd", "abdc", 3, 0},
		{"return to original", "abcd", "abcd", 2, 3},
		{"large odd parity", "ab", "ba", 1000000000000000, 0},
		{"large even parity", "ab", "ab", 1000000000000000, 1},
		{"modulo", "aaaa", "aaaa", 20, 486784380},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWays(tt.s, tt.target, tt.k); got != tt.want {
				t.Errorf("numberOfWays(%q, %q, %d) = %d, want %d", tt.s, tt.target, tt.k, got, tt.want)
			}
		})
	}
}
