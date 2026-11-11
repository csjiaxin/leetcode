package l692

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	for _, tt := range []struct {
		name  string
		words []string
		k     int
		want  []string
	}{
		{"example one", []string{"i", "love", "leetcode", "i", "love", "coding"}, 2, []string{"i", "love"}},
		{"example two", []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4, []string{"the", "is", "sunny", "day"}},
		{"lexical tie", []string{"b", "a", "c"}, 2, []string{"a", "b"}},
		{"single", []string{"z", "z"}, 1, []string{"z"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.words, tt.k); !slices.Equal(got, tt.want) {
				t.Errorf("%s: got %v want %v", t.Name(), got, tt.want)
			}
		})
	}
}
