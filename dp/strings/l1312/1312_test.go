package l1312

import (
	"strings"
	"testing"
)

func TestMinInsertions(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "example one", s: "zzazz", want: 0},
		{name: "example two", s: "mbadm", want: 2},
		{name: "example three", s: "leetcode", want: 5},
		{name: "single character", s: "a", want: 0},
		{name: "two matching characters", s: "aa", want: 0},
		{name: "two different characters", s: "ab", want: 1},
		{name: "odd length palindrome", s: "racecar", want: 0},
		{name: "even length palindrome", s: "abba", want: 0},
		{name: "all identical characters", s: "aaaaa", want: 0},
		{name: "all distinct characters", s: "abcdef", want: 5},
		{name: "one insertion at beginning", s: "abb", want: 1},
		{name: "one insertion at end", s: "aab", want: 1},
		{name: "one insertion in middle", s: "abca", want: 1},
		{name: "matching endpoints with distinct interior", s: "abcda", want: 2},
		{name: "use subsequence not substring", s: "bbbab", want: 1},
		{name: "skip interior characters", s: "agbdba", want: 1},
		{name: "alternating characters", s: "ababab", want: 1},
		{name: "separated repeated characters", s: "abcabc", want: 3},
		{name: "multiple optimal insertion choices", s: "aebcbda", want: 2},
		{name: "mixed characters", s: "character", want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInsertions(tt.s); got != tt.want {
				t.Errorf("%s: minInsertions(%q) = %d, want %d", t.Name(), tt.s, got, tt.want)
			}
		})
	}
}

func TestMinInsertionsMaximumLength(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "all identical", s: strings.Repeat("a", 500), want: 0},
		{name: "two distinct blocks", s: strings.Repeat("a", 250) + strings.Repeat("b", 250), want: 250},
		{name: "alternating characters", s: strings.Repeat("ab", 250), want: 1},
		{name: "matching outer blocks", s: strings.Repeat("a", 125) + strings.Repeat("b", 250) + strings.Repeat("a", 125), want: 0},
		{name: "single different endpoint", s: strings.Repeat("a", 499) + "b", want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInsertions(tt.s); got != tt.want {
				t.Errorf("%s: minInsertions(string of length %d) = %d, want %d", t.Name(), len(tt.s), got, tt.want)
			}
		})
	}
}
