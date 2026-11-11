package l567

import "testing"

func TestCheckInclusion(t *testing.T) {
	for _, tt := range []struct {
		name   string
		s1, s2 string
		want   bool
	}{
		{"example true", "ab", "eidbaooo", true},
		{"example false", "ab", "eidboaoo", false},
		{"pattern longer", "abcd", "abc", false},
		{"same length permutation", "abc", "cba", true},
		{"single present", "a", "ba", true},
		{"single absent", "a", "bbb", false},
		{"repeated letters", "aab", "eidbaaoo", true},
		{"wrong multiplicities", "aab", "abb", false},
		{"match at start", "ab", "abzzz", true},
		{"match at end", "ab", "zzba", true},
		{"window must shrink", "ab", "aaab", true},
		{"match after excess different letters", "adc", "dcda", true},
		{"missing and excess letters", "abc", "aaabbb", false},
		{"reset after unrelated letters", "ab", "aazba", true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.s1, tt.s2); got != tt.want {
				t.Errorf("checkInclusion(%q, %q) = %v, want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}
