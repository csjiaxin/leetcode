package l3

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, tt := range []struct {
		name, s string
		want    int
	}{
		{"example", "abcabcbb", 3}, {"same", "bbbbb", 1}, {"subsequence not substring", "pwwkew", 3},
		{"empty", "", 0}, {"shrink repeatedly", "abba", 2}, {"space", " ", 1},
		{"no backward left", "tmmzuxt", 5},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
