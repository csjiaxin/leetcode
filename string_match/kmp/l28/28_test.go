package l28

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		name, text, pattern string
		want                int
	}{
		{"first occurrence", "sadbutsad", "sad", 0},
		{"absent", "leetcode", "leeto", -1},
		{"single", "a", "a", 0},
		{"at end", "hello", "lo", 3},
		{"overlap", "aaaaa", "aaa", 0},
		{"fallback", "abababc", "ababc", 2},
		{"repeated fallback", "mississippi", "issip", 4},
		{"longer pattern", "ab", "abc", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.text, tt.pattern); got != tt.want {
				t.Errorf("strStr(%q, %q) = %d, want %d", tt.text, tt.pattern, got, tt.want)
			}
		})
	}
}
