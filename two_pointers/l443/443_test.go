package l443

import "testing"

func TestCompress(t *testing.T) {
	for _, tt := range []struct{ name, input, want string }{
		{"example", "aabbccc", "a2b2c3"}, {"single", "a", "a"}, {"multi digit", "abbbbbbbbbbbb", "ab12"},
		{"no repeats", "abc", "abc"}, {"ten", "aaaaaaaaaa", "a10"}, {"digit chars", "11122", "1322"},
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
