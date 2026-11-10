package l1143

import (
	"strings"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name  string
		text1 string
		text2 string
		want  int
	}{
		{name: "example one", text1: "abcde", text2: "ace", want: 3},
		{name: "example two", text1: "abc", text2: "abc", want: 3},
		{name: "example three", text1: "abc", text2: "def", want: 0},
		{name: "single matching character", text1: "a", text2: "a", want: 1},
		{name: "single different characters", text1: "a", text2: "b", want: 0},
		{name: "match at beginning", text1: "abcd", text2: "a", want: 1},
		{name: "match at end", text1: "abcd", text2: "d", want: 1},
		{name: "second string longer", text1: "ace", text2: "abcde", want: 3},
		{name: "subsequence not substring", text1: "axbxcxd", text2: "abcd", want: 4},
		{name: "order matters", text1: "abc", text2: "cba", want: 1},
		{name: "repeated characters", text1: "aaaaa", text2: "aaa", want: 3},
		{name: "cannot reuse a character", text1: "a", text2: "aaaa", want: 1},
		{name: "skip characters in both strings", text1: "abcbdab", text2: "bdcaba", want: 4},
		{name: "multiple optimal subsequences", text1: "abab", text2: "baba", want: 3},
		{name: "shared suffix", text1: "xyzabc", text2: "defabc", want: 3},
		{name: "shared prefix", text1: "abcxyz", text2: "abcdef", want: 3},
		{name: "repeated character ordering", text1: "aabb", text2: "bbaa", want: 2},
		{name: "nonlocal matches", text1: "bsbininm", text2: "jmjkbkjkv", want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.text1, tt.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence(%q, %q) = %d, want %d", tt.text1, tt.text2, got, tt.want)
			}
		})
	}
}

func TestLongestCommonSubsequenceMaximumLength(t *testing.T) {
	tests := []struct {
		name  string
		text1 string
		text2 string
		want  int
	}{
		{name: "identical strings", text1: strings.Repeat("a", 1000), text2: strings.Repeat("a", 1000), want: 1000},
		{name: "disjoint strings", text1: strings.Repeat("a", 1000), text2: strings.Repeat("b", 1000), want: 0},
		{name: "reversed blocks", text1: strings.Repeat("a", 500) + strings.Repeat("b", 500), text2: strings.Repeat("b", 500) + strings.Repeat("a", 500), want: 500},
		{name: "very different lengths", text1: strings.Repeat("a", 1000), text2: "a", want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.text1, tt.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence(strings of lengths %d and %d) = %d, want %d", len(tt.text1), len(tt.text2), got, tt.want)
			}
		})
	}
}
