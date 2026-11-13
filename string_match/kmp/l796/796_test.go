package l796

import "testing"

func TestRotateString(t *testing.T) {
	tests := []struct {
		name, s, goal string
		want          bool
	}{
		{"example", "abcde", "cdeab", true}, {"wrong order", "abcde", "abced", false},
		{"single", "a", "a", true}, {"unchanged", "abc", "abc", true},
		{"periodic", "abab", "baba", true}, {"unequal lengths", "ab", "abab", false},
		{"different character", "aaa", "aab", false}, {"last shift", "abc", "cab", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateString(tt.s, tt.goal); got != tt.want {
				t.Errorf("rotateString(%q, %q) = %v, want %v", tt.s, tt.goal, got, tt.want)
			}
		})
	}
}
