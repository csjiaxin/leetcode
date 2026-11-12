package l1170

import (
	"slices"
	"testing"
)

func TestNumSmallerByFrequency(t *testing.T) {
	for _, tt := range []struct {
		name           string
		queries, words []string
		want           []int
	}{
		{"example", []string{"cbd"}, []string{"zaaaz"}, []int{1}},
		{"multiple", []string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}, []int{1, 2}},
		{"equal", []string{"ab", "zz"}, []string{"a", "ba", "cc"}, []int{1, 0}},
		{"smallest not most frequent", []string{"abbbb"}, []string{"zz", "aab"}, []int{2}},
		{"no words", []string{"a"}, nil, []int{0}}, {"no queries", nil, []string{"a"}, []int{}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSmallerByFrequency(tt.queries, tt.words); !slices.Equal(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
