package l3

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, tt := range []struct {
		name, s string
		want    int
	}{
		{"example one", "abcabcbb", 3},
		{"all equal", "bbbbb", 1},
		{"substring not subsequence", "pwwkew", 3},
		{"empty", "", 0},
		{"single", "a", 1},
		{"all distinct", "abcdef", 6},
		{"left cannot go backward", "abba", 2},
		{"duplicate outside window", "tmmzuxt", 5},
		{"space", " ", 1},
		{"symbols and case", "aA1! a", 5},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
