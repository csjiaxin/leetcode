package l680

import "testing"

func TestValidPalindrome(t *testing.T) {
	for _, tt := range []struct {
		name, s string
		want    bool
	}{
		{"palindrome", "aba", true}, {"delete middle", "abca", true}, {"impossible", "abc", false},
		{"single", "a", true}, {"two", "ab", true}, {"delete left", "deeee", true},
		{"delete right", "eeeed", true}, {"two mismatches", "abcda", false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.s); got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
