package l567

import "testing"

func TestCheckInclusion(t *testing.T) {
	for _, tt := range []struct {
		name, s1, s2 string
		want         bool
	}{
		{"example", "ab", "eidbaooo", true}, {"missing", "ab", "eidboaoo", false},
		{"too short", "abc", "ab", false}, {"counts", "aab", "caba", true},
		{"insufficient repeats", "aab", "abbb", false}, {"single", "a", "a", true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.s1, tt.s2); got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
