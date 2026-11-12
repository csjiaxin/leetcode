package l1358

import (
	"strings"
	"testing"
)

func TestNumberOfSubstrings(t *testing.T) {
	for _, tt := range []struct {
		s    string
		want int
	}{
		// {"abcabc", 10},
		// {"aaacb", 3},
		{"abc", 1},
		{"aaa", 0},
		{"abab", 0},
		{"aabbcc", 4},
		{"abcc", 2},
		{"cba", 1},
	} {
		t.Run(tt.s, func(t *testing.T) {
			if got := numberOfSubstrings(tt.s); got != tt.want {
				t.Errorf("numberOfSubstrings(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestNumberOfSubstringsMaximum(t *testing.T) {
	const n = 50000
	s := strings.Repeat("abc", n/3) + "abc"[:n%3]
	// In this periodic string every substring of length at least 3 qualifies.
	want := int(int64(n-1) * int64(n-2) / 2)
	if got := numberOfSubstrings(s); got != want {
		t.Errorf("maximum periodic input: got %d, want %d", got, want)
	}
}
