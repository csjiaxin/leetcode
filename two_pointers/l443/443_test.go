package l443

import (
	"strings"
	"testing"
)

func TestCompress(t *testing.T) {
	for _, tt := range []struct{ name, input, want string }{
		{"example", "aabbccc", "a2b2c3"}, {"single", "a", "a"}, {"multi digit", "abbbbbbbbbbbb", "ab12"},
		{"no repeats", "abc", "abc"}, {"ten", "aaaaaaaaaa", "a10"},
		{"digit chars", "11122", "1322"},
		{"two equal", "aa", "a2"},
		{"separate groups", "aabbaa", "a2b2a2"},
		{"single after group", "aaab", "a3b"},
		{"digit after group", "aaa3", "a33"},
		{"three digit count", strings.Repeat("a", 123), "a123"},
		{"adjacent long groups", strings.Repeat("a", 12) + strings.Repeat("b", 10), "a12b10"},
		{"maximum length group", strings.Repeat("z", 2000), "z2000"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			chars := []byte(tt.input)
			k := compress(chars)
			if k < 0 || k > len(chars) {
				t.Fatalf("invalid length %d", k)
			}
			if got := string(chars[:k]); got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}
