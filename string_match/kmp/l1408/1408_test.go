package l1408

import (
	"slices"
	"testing"
)

func TestStringMatching(t *testing.T) {
	tests := []struct {
		name        string
		words, want []string
	}{
		{"example", []string{"mass", "as", "hero", "superhero"}, []string{"as", "hero"}},
		{"second example", []string{"leetcode", "et", "code"}, []string{"et", "code"}},
		{"absent", []string{"blue", "green", "bu"}, []string{}},
		{"nested", []string{"a", "aa", "aaa"}, []string{"a", "aa"}},
		{"multiple witnesses", []string{"ab", "abc", "zab"}, []string{"ab"}},
		{"suffix", []string{"xyz", "yz"}, []string{"yz"}},
		{"equal lengths", []string{"ab", "ba"}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Clone(stringMatching(slices.Clone(tt.words)))
			want := slices.Clone(tt.want)
			slices.Sort(got)
			slices.Sort(want)
			if !slices.Equal(got, want) {
				t.Errorf("stringMatching(%v) = %v, want %v (any order)", tt.words, got, want)
			}
		})
	}
}
