package l516

import (
	"strings"
	"testing"
)

func TestLongestPalindromeSubseq(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "example one", s: "bbbab", want: 4},
		{name: "example two", s: "cbbd", want: 2},
		{name: "single character", s: "a", want: 1},
		{name: "two matching characters", s: "aa", want: 2},
		{name: "two different characters", s: "ab", want: 1},
		{name: "odd length palindrome", s: "racecar", want: 7},
		{name: "even length palindrome", s: "abba", want: 4},
		{name: "all identical characters", s: "aaaaa", want: 5},
		{name: "all distinct characters", s: "abcdef", want: 1},
		{name: "subsequence not substring", s: "agbdba", want: 5},
		{name: "skip left endpoint", s: "xabba", want: 4},
		{name: "skip right endpoint", s: "abbax", want: 4},
		{name: "skip both endpoints", s: "xabbay", want: 4},
		{name: "matching endpoints with distinct interior", s: "abcda", want: 3},
		{name: "alternating characters", s: "ababab", want: 5},
		{name: "multiple optimal subsequences", s: "aebcbda", want: 5},
		{name: "separated repeated characters", s: "abcabc", want: 3},
		{name: "longer mixed example", s: "character", want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeSubseq(tt.s); got != tt.want {
				t.Errorf("longestPalindromeSubseq(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestLongestPalindromeSubseqMaximumLength(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "all identical", s: strings.Repeat("a", 1000), want: 1000},
		{name: "two distinct blocks", s: strings.Repeat("a", 500) + strings.Repeat("b", 500), want: 500},
		{name: "alternating characters", s: strings.Repeat("ab", 500), want: 999},
		{name: "matching outer blocks", s: strings.Repeat("a", 250) + strings.Repeat("b", 500) + strings.Repeat("a", 250), want: 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeSubseq(tt.s); got != tt.want {
				t.Errorf("longestPalindromeSubseq(string of length %d) = %d, want %d", len(tt.s), got, tt.want)
			}
		})
	}
}
