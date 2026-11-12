package l496

import (
	"slices"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	for _, tt := range []struct {
		name       string
		a, b, want []int
	}{
		{"example", []int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{"example two", []int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
		{"single", []int{1}, []int{1}, []int{-1}},
		{"descending", []int{3, 2, 1}, []int{3, 2, 1}, []int{-1, -1, -1}},
		{"nearest not maximum", []int{1}, []int{1, 3, 2, 4}, []int{3}},
		{"query order", []int{2, 1}, []int{1, 2, 3}, []int{3, 2}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(slices.Clone(tt.a), slices.Clone(tt.b)); !slices.Equal(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
