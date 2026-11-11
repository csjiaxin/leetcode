package l438

import (
	"slices"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	for _, tt := range []struct {
		name string
		s, p string
		want []int
	}{
		{"example one", "cbaebabacd", "abc", []int{0, 6}},
		{"overlapping", "abab", "ab", []int{0, 1, 2}},
		{"pattern longer", "ab", "abc", nil},
		{"no match", "abcd", "ee", nil},
		{"same length", "bca", "abc", []int{0}},
		{"repeated letters", "aaaaa", "aa", []int{0, 1, 2, 3}},
		{"wrong multiplicity", "abb", "aab", nil},
		{"single character", "baa", "a", []int{1, 2}},
		{"match at end", "zzabc", "cba", []int{2}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := findAnagrams(tt.s, tt.p)
			// The problem allows indices in any order; nil and empty are equivalent.
			got = slices.Clone(got)
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("findAnagrams(%q, %q) = %v, want %v", tt.s, tt.p, got, tt.want)
			}
		})
	}
}
