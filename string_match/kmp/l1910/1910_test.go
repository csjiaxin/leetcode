package l1910

import "testing"

func TestRemoveOccurrences(t *testing.T) {
	tests := []struct{ name, s, part, want string }{
		{"example", "daabcbaabcbc", "abc", "dab"},
		{"cascade", "axxxxyyyyb", "xy", "ab"},
		{"all removed", "aaaa", "aa", ""},
		{"leftmost overlap", "ababa", "aba", "ba"},
		{"absent", "abcd", "ef", "abcd"},
		{"whole string", "abc", "abc", ""},
		{"restore previous state", "aabcbc", "abc", ""},
		{"single character", "banana", "a", "bnn"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOccurrences(tt.s, tt.part); got != tt.want {
				t.Errorf("removeOccurrences(%q, %q) = %q, want %q", tt.s, tt.part, got, tt.want)
			}
		})
	}
}
