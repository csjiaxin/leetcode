package l567

import "testing"

func TestCheckInclusion(t *testing.T) {
	for _, tt := range []struct {
		name, s1, s2 string
		want         bool
	}{
		{"example true", "ab", "eidbaooo", true},
		{"example false", "ab", "eidboaoo", false},
		{"pattern longer", "abcd", "abc", false},
		{"same length", "abc", "cba", true},
		{"repeated letters", "aab", "eidbaaoo", true},
		{"wrong multiplicities", "aab", "abb", false},
		{"match at start", "ab", "abzzz", true},
		{"match at end", "ab", "zzba", true},
		{"excess duplicate", "ab", "aaab", true},
		{"single absent", "a", "bbb", false},
		{"reset after unrelated", "adc", "dcda", true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.s1, tt.s2); got != tt.want {
				t.Errorf("checkInclusion(%q, %q) = %v, want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}
