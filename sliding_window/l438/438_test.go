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
		{"overlap", "abab", "ab", []int{0, 1, 2}},
		{"pattern longer", "ab", "abc", nil},
		{"no match", "abcd", "ee", nil},
		{"same length", "bca", "abc", []int{0}},
		{"duplicates", "aaaaa", "aa", []int{0, 1, 2, 3}},
		{"wrong multiplicity", "abb", "aab", nil},
		{"single", "baa", "a", []int{1, 2}},
		{"last window", "zzabc", "cba", []int{2}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Clone(findAnagrams(tt.s, tt.p))
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("findAnagrams(%q, %q) = %v, want %v", tt.s, tt.p, got, tt.want)
			}
		})
	}
}
