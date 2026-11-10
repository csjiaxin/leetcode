package l664

import (
	"strings"
	"testing"
)

func TestStrangePrinter(t *testing.T) {
	tests := []struct {
		name, s string
		want    int
	}{
		{"example one", "aaabbb", 2},
		{"example two", "aba", 2},
		{"single", "a", 1},
		{"identical", "aaaa", 1},
		{"distinct", "abc", 3},
		{"nested", "abcba", 3},
		{"alternating", "ababab", 4},
		{"nonadjacent duplicates", "ababc", 4},
		{"compress adjacent only", "aaabbbaa", 2},
		{"maximum identical", strings.Repeat("a", 100), 1},
		{"maximum alternating", strings.Repeat("ab", 50), 51},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strangePrinter(tt.s); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
