package l214

import "testing"

func TestShortestPalindrome(t *testing.T) {
	tests := []struct{ name, s, want string }{
		{"example one", "aacecaaa", "aaacecaaa"}, {"example two", "abcd", "dcbabcd"},
		{"empty", "", ""}, {"single", "a", "a"}, {"already palindrome", "abba", "abba"},
		{"odd prefix", "abac", "cabac"}, {"repeated", "aaaa", "aaaa"}, {"two different", "ab", "bab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPalindrome(tt.s); got != tt.want {
				t.Errorf("shortestPalindrome(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
