package l76

import "testing"

func TestMinWindow(t *testing.T) {
	for _, tt := range []struct {
		name, s, target, want string
	}{
		{"example", "ADOBECODEBANC", "ABC", "BANC"},
		{"single", "a", "a", "a"},
		{"empty target", "abc", "", ""},
		{"empty source", "", "a", ""},
		{"both empty", "", "", ""},
		{"excess copies remain satisfied", "AAAB", "AAB", "AAB"},
		{"irrelevant characters inside", "AxB", "AB", "AxB"},
		{"not enough copies", "a", "aa", ""},
		{"no match", "abc", "z", ""},
		{"case sensitive", "aA", "A", "A"},
		{"duplicate demand", "AAABBC", "AABC", "AABBC"},
		{"whole string", "abc", "cba", "abc"},
		{"shrink across excess", "bbaa", "aba", "baa"},
		{"best at end", "abdbca", "abc", "bca"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.s, tt.target); got != tt.want {
				t.Errorf("minWindow(%q, %q) = %q, want %q", tt.s, tt.target, got, tt.want)
			}
		})
	}
}
