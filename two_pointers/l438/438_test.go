package l438

import (
	"slices"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	for _, tt := range []struct {
		name, s, p string
		want       []int
	}{
		{"example", "cbaebabacd", "abc", []int{0, 6}},
		{"overlap", "abab", "ab", []int{0, 1, 2}}, {"too short", "a", "ab", nil},
		{"counts", "aaab", "aab", []int{1}}, {"no match", "abc", "d", nil},
		{"single", "a", "a", []int{0}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.s, tt.p); !slices.Equal(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
