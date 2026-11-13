package l1392

import "testing"

func TestLongestPrefix(t *testing.T) {
	tests := []struct{ s, want string }{
		{"level", "l"}, {"ababab", "abab"}, {"a", ""}, {"abcd", ""},
		{"aaaa", "aaa"}, {"abab", "ab"}, {"ababa", "aba"}, {"ababaa", "a"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := longestPrefix(tt.s); got != tt.want {
				t.Errorf("longestPrefix(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
