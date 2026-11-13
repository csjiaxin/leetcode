package l459

import "testing"

func TestRepeatedSubstringPattern(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"abab", true}, {"aba", false}, {"abcabcabcabc", true}, {"a", false},
		{"aa", true}, {"abc", false}, {"ababa", false}, {"ababaa", false},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := repeatedSubstringPattern(tt.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
