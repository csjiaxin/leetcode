package l3008

import (
	"slices"
	"strings"
	"testing"
)

func TestBeautifulIndices(t *testing.T) {
	tests := []struct {
		name, s, a, b string
		k             int
		want          []int
	}{
		{"example", "isawsquirrelnearmysquirrelhouseohmy", "my", "squirrel", 15, []int{16, 33}},
		{"exact distance", "abcdef", "a", "c", 2, []int{0}},
		{"outside distance", "abcdef", "a", "c", 1, []int{}},
		{"overlaps", "aaaa", "aa", "aa", 0, []int{0, 1, 2}},
		{"b before a", "abcxxabc", "abc", "x", 2, []int{5}},
		{"missing a", "abc", "z", "a", 1, []int{}},
		{"missing b", "abab", "ab", "z", 3, []int{}},
		{"multiple witnesses no duplicate", "ababa", "a", "b", 1, []int{0, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := beautifulIndices(tt.s, tt.a, tt.b, tt.k); !slices.Equal(got, tt.want) {
				t.Errorf("beautifulIndices(%q, %q, %q, %d) = %v, want %v", tt.s, tt.a, tt.b, tt.k, got, tt.want)
			}
		})
	}
}

func TestBeautifulIndicesLarge(t *testing.T) {
	s := strings.Repeat("a", 100000)
	got := beautifulIndices(s, "aa", "a", 0)
	if len(got) != len(s)-1 {
		t.Fatalf("got %d indices, want %d", len(got), len(s)-1)
	}
	for i, v := range got {
		if v != i {
			t.Fatalf("index %d = %d, want %d", i, v, i)
		}
	}
}
