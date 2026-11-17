package l115

import (
	"strings"
	"testing"
)

func TestNumDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want int
	}{
		{name: "official rabbits", s: "rabbbit", t: "rabbit", want: 3},
		{name: "official bags", s: "babgbag", t: "bag", want: 5},
		{name: "both empty", s: "", t: "", want: 1},
		{name: "empty target", s: "abc", t: "", want: 1},
		{name: "empty source", s: "", t: "abc", want: 0},
		{name: "identical", s: "abc", t: "abc", want: 1},
		{name: "target longer", s: "abc", t: "abcd", want: 0},
		{name: "no match", s: "abc", t: "d", want: 0},
		{name: "order matters", s: "abc", t: "ba", want: 0},
		{name: "repeated letters", s: "aaaaa", t: "aaa", want: 10},
		{
			name: "large exact bounded result",
			s:    strings.Repeat("a", 33),
			t:    strings.Repeat("a", 16),
			want: 1166803110, // C(33, 16), below MaxInt32.
		},
		{
			name: "huge intermediate counts with unmatched suffix",
			s:    strings.Repeat("a", 100) + "b",
			t:    strings.Repeat("a", 50) + "c",
			want: 0,
		},
		{
			name: "huge later prefix counts leave completed target unchanged",
			s:    strings.Repeat("a", 50) + "b" + strings.Repeat("a", 100),
			t:    strings.Repeat("a", 50) + "b",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDistinct(tt.s, tt.t); got != tt.want {
				t.Errorf("numDistinct(%q, %q) = %d; want %d", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
