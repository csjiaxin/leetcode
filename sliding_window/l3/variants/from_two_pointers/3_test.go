package l3

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, tt := range []struct {
		name, s string
		want    int
	}{
		{"example", "abcabcbb", 3}, {"equal", "bbbbb", 1}, {"middle", "pwwkew", 3},
		{"empty", "", 0}, {"left cannot retreat", "abba", 2}, {"space", " ", 1}, {"symbols", "a!b a", 4},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
