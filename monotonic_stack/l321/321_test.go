package l321

import (
	"slices"
	"testing"
)

func TestMaxNumber(t *testing.T) {
	for _, tt := range []struct {
		name string
		a, b []int
		k    int
		want []int
	}{
		{"example one", []int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5, []int{9, 8, 6, 5, 3}},
		{"example two", []int{6, 7}, []int{6, 0, 4}, 5, []int{6, 7, 6, 0, 4}},
		{"example three", []int{3, 9}, []int{8, 9}, 3, []int{9, 8, 9}},
		{"zero length", []int{1}, []int{2}, 0, []int{}},
		{"one digit", []int{1, 9}, []int{8}, 1, []int{9}},
		{"equal prefix compare suffix", []int{6, 0}, []int{6, 7}, 4, []int{6, 7, 6, 0}},
		{"all zeros", []int{0, 0}, []int{0}, 3, []int{0, 0, 0}},
		{"all from one input", []int{9, 8, 7}, []int{1, 2}, 3, []int{9, 8, 7}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumber(slices.Clone(tt.a), slices.Clone(tt.b), tt.k); !slices.Equal(got, tt.want) {
				t.Errorf("%s: got %v, want %v", t.Name(), got, tt.want)
			}
		})
	}
}
