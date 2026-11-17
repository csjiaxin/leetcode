package l583

import "testing"

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name         string
		word1, word2 string
		want         int
	}{
		{"official_sea_eat", "sea", "eat", 2},
		{"official_subsequence", "leetcode", "etco", 4},
		{"empty_both", "", "", 0},
		{"empty_first", "", "abc", 3},
		{"empty_second", "abc", "", 3},
		{"equal", "abc", "abc", 0},
		{"replacement_requires_two_deletions", "a", "b", 2},
		{"disjoint", "ab", "cd", 4},
		{"repeated", "aaaa", "aa", 2},
		{"order_matters", "ab", "ba", 2},
		{"repeated_different_counts", "aaab", "abbb", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.word1, tt.word2); got != tt.want {
				t.Fatalf("minDistance(%q, %q) = %d, want %d", tt.word1, tt.word2, got, tt.want)
			}
			if got := minDistance(tt.word2, tt.word1); got != tt.want {
				t.Fatalf("swapped minDistance(%q, %q) = %d, want %d", tt.word2, tt.word1, got, tt.want)
			}
		})
	}
}
