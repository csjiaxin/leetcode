package l72

import (
	"strings"
	"testing"
)

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		{name: "example one", word1: "horse", word2: "ros", want: 3},
		{name: "example two", word1: "intention", word2: "execution", want: 5},
		{name: "both empty", word1: "", word2: "", want: 0},
		{name: "insert into empty word", word1: "", word2: "abc", want: 3},
		{name: "delete entire word", word1: "abc", word2: "", want: 3},
		{name: "identical words", word1: "algorithm", word2: "algorithm", want: 0},
		{name: "single matching character", word1: "a", word2: "a", want: 0},
		{name: "single replacement", word1: "a", word2: "b", want: 1},
		{name: "insert at beginning", word1: "abc", word2: "xabc", want: 1},
		{name: "insert in middle", word1: "ac", word2: "abc", want: 1},
		{name: "insert at end", word1: "abc", word2: "abcd", want: 1},
		{name: "delete at beginning", word1: "xabc", word2: "abc", want: 1},
		{name: "delete in middle", word1: "abc", word2: "ac", want: 1},
		{name: "delete at end", word1: "abcd", word2: "abc", want: 1},
		{name: "replace in middle", word1: "abc", word2: "axc", want: 1},
		{name: "replace every character", word1: "abc", word2: "xyz", want: 3},
		{name: "adjacent swap is not one operation", word1: "ab", word2: "ba", want: 2},
		{name: "shift using delete and insert", word1: "abcde", word2: "bcdea", want: 2},
		{name: "repeated characters", word1: "aaaaa", word2: "aaa", want: 2},
		{name: "mixed operations", word1: "kitten", word2: "sitting", want: 3},
		{name: "shared suffix", word1: "walking", word2: "talking", want: 1},
		{name: "unequal lengths without shared characters", word1: "abc", word2: "xy", want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.word1, tt.word2); got != tt.want {
				t.Errorf("minDistance(%q, %q) = %d, want %d", tt.word1, tt.word2, got, tt.want)
			}
		})
	}
}

func TestMinDistanceMaximumLength(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		{name: "identical words", word1: strings.Repeat("a", 500), word2: strings.Repeat("a", 500), want: 0},
		{name: "replace all characters", word1: strings.Repeat("a", 500), word2: strings.Repeat("b", 500), want: 500},
		{name: "insert maximum length word", word1: "", word2: strings.Repeat("a", 500), want: 500},
		{name: "delete maximum length word", word1: strings.Repeat("a", 500), word2: "", want: 500},
		{name: "very different lengths", word1: strings.Repeat("a", 500), word2: "a", want: 499},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.word1, tt.word2); got != tt.want {
				t.Errorf("minDistance(words of lengths %d and %d) = %d, want %d", len(tt.word1), len(tt.word2), got, tt.want)
			}
		})
	}
}
