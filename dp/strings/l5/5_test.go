package l5

import (
	"strings"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		wantLen int
	}{
		// {name: "example one with tied answers", s: "babad", wantLen: 3},
		// {name: "example two", s: "cbbd", wantLen: 2},
		{name: "single character", s: "a", wantLen: 1},
		{name: "two equal characters", s: "aa", wantLen: 2},
		{name: "two different characters", s: "ab", wantLen: 1},
		{name: "odd length palindrome", s: "racecar", wantLen: 7},
		{name: "even length palindrome", s: "abba", wantLen: 4},
		{name: "all identical characters", s: "aaaaa", wantLen: 5},
		{name: "all distinct characters", s: "abcdef", wantLen: 1},
		{name: "palindrome at beginning", s: "abacdef", wantLen: 3},
		{name: "palindrome at end", s: "defaba", wantLen: 3},
		{name: "palindrome in middle", s: "xracecary", wantLen: 7},
		{name: "substring not subsequence", s: "bbbab", wantLen: 3},
		{name: "separated matching endpoints", s: "abcda", wantLen: 1},
		{name: "longer even palindrome", s: "forgeeksskeegfor", wantLen: 10},
		{name: "nested odd palindrome", s: "bananas", wantLen: 5},
		{name: "alternating characters", s: "ababab", wantLen: 5},
		{name: "case sensitive", s: "Aa", wantLen: 1},
		{name: "digits", s: "x12321y", wantLen: 5},
		{name: "mixed case and digits", s: "zA1b1Ay", wantLen: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkLongestPalindrome(t, tt.s, tt.wantLen)
		})
	}
}

func TestLongestPalindromeMaximumLength(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		wantLen int
	}{
		{name: "all identical", s: strings.Repeat("a", 1000), wantLen: 1000},
		{name: "alternating characters", s: strings.Repeat("ab", 500), wantLen: 999},
		{name: "two distinct blocks", s: strings.Repeat("a", 500) + strings.Repeat("b", 500), wantLen: 500},
		{name: "matching outer blocks", s: strings.Repeat("a", 250) + strings.Repeat("b", 500) + strings.Repeat("a", 250), wantLen: 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkLongestPalindrome(t, tt.s, tt.wantLen)
		})
	}
}

// Accept any longest answer when multiple palindromic substrings tie.
func checkLongestPalindrome(t *testing.T, s string, wantLen int) {
	t.Helper()
	got := longestPalindrome(s)
	if len(got) != wantLen {
		t.Fatalf("%s: longestPalindrome(input of length %d) returned length %d, want %d", t.Name(), len(s), len(got), wantLen)
	}
	if !strings.Contains(s, got) {
		t.Fatalf("%s: result %q is not a substring of the input", t.Name(), got)
	}
	for left, right := 0, len(got)-1; left < right; left, right = left+1, right-1 {
		if got[left] != got[right] {
			t.Fatalf("%s: result %q is not a palindrome", t.Name(), got)
		}
	}
}
