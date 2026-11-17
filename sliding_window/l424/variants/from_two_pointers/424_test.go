package l424

import "testing"

func TestCharacterReplacement(t *testing.T) {
	for _, tt := range []struct {
		name, s string
		k, want int
	}{
		{"example", "ABAB", 2, 4}, {"shrink", "AABABBA", 1, 4},
		{"no replacement", "ABBB", 0, 3}, {"single", "A", 0, 1},
		{"replace all", "ABCDE", 5, 5}, {"equal", "AAAA", 0, 4},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.s, tt.k); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
