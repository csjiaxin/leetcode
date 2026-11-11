package l424

import (
	"strings"
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	for _, tt := range []struct {
		name, s string
		k, want int
	}{
		{"example one", "ABAB", 2, 4},
		{"example two", "AABABBA", 1, 4},
		{"single", "A", 0, 1},
		{"all equal no changes", "AAAA", 0, 4},
		{"no replacements", "ABBBCC", 0, 3},
		{"all distinct", "ABCDE", 1, 2},
		{"replace whole string", "ABCDE", 5, 5},
		{"best at start", "AAAABC", 1, 5},
		{"best at end", "ABCCCC", 1, 5},
		{"historical max frequency", "AAAAABCDE", 1, 6},
		{"alternating", "ABABABAB", 1, 3},
		{"alphabet endpoints", "AZZA", 1, 3},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.s, tt.k); got != tt.want {
				t.Errorf("characterReplacement(%q, %d) = %d, want %d", tt.s, tt.k, got, tt.want)
			}
		})
	}
}

func TestCharacterReplacementMaximum(t *testing.T) {
	const n = 100000
	if got := characterReplacement(strings.Repeat("Z", n), 0); got != n {
		t.Errorf("maximum equal input: got %d, want %d", got, n)
	}
}
