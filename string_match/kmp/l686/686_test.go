package l686

import "testing"

func TestRepeatedStringMatch(t *testing.T) {
	tests := []struct {
		name, a, b string
		want       int
	}{
		{"example", "abcd", "cdabcdab", 3}, {"single repeated", "a", "aa", 2},
		{"absent character", "a", "b", -1}, {"already present", "abcd", "bc", 1},
		{"cross boundary", "abc", "cab", 2}, {"needs extra copy", "abc", "cabcab", 3},
		{"wrong order", "ab", "aab", -1}, {"identical", "abc", "abc", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedStringMatch(tt.a, tt.b); got != tt.want {
				t.Errorf("repeatedStringMatch(%q, %q) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
