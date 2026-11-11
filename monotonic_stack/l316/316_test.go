package l316

import (
	"strings"
	"testing"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	for _, tt := range []struct{ name, s, want string }{
		{"example one", "bcabc", "abc"}, {"example two", "cbacdcbc", "acdb"},
		{"single", "z", "z"}, {"identical", "aaaa", "a"}, {"already unique", "dcba", "dcba"},
		{"last occurrence must stay", "bbcaac", "bac"},
		{"repeated pattern", "abacb", "abc"},
		{"cannot sort arbitrarily", "ecbacba", "eacb"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateLetters(tt.s); got != tt.want {
				t.Errorf("%s: got %q, want %q", t.Name(), got, tt.want)
			}
		})
	}
}
func TestRemoveDuplicateLettersMaximum(t *testing.T) {
	if got := removeDuplicateLetters(strings.Repeat("a", 10000)); got != "a" {
		t.Errorf("%s: got %q, want a", t.Name(), got)
	}
}
