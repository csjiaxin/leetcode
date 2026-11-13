package l392

import "testing"

func TestIsSubsequence(t *testing.T) {
	for _, tt := range []struct {
		name, s, text string
		want          bool
	}{
		{"example", "abc", "ahbgdc", true}, {"missing", "axc", "ahbgdc", false},
		{"empty source", "", "abc", true}, {"empty target", "a", "", false},
		{"both empty", "", "", true}, {"order matters", "ba", "ab", false}, {"repeats", "aaa", "aa", false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.s, tt.text); got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
