package l15

import (
	"slices"
	"testing"
)

// Normalize order only: keep duplicates so duplicate answers fail the test.
func normalize(t *testing.T, values [][]int) [][]int {
	t.Helper()
	out := make([][]int, len(values))
	for i, v := range values {
		if len(v) != 3 {
			t.Fatalf("invalid triple: %v", v)
		}
		out[i] = slices.Clone(v)
		slices.Sort(out[i])
	}
	slices.SortFunc(out, func(a, b []int) int { return slices.Compare(a, b) })
	return out
}

func TestThreeSum(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		want [][]int
	}{
		{"example", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"no solution", []int{0, 1, 1}, nil},
		{"all zero", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{"duplicate pairs", []int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
		{"multiple", []int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
		{"all negative", []int{-5, -3, -1}, nil},
		{"all positive", []int{1, 2, 3}, nil},
		{"exactly three", []int{-1, 0, 1}, [][]int{{-1, 0, 1}}},
		{"repeated fixed value", []int{-1, -1, -1, 2, 2}, [][]int{{-1, -1, 2}}},
		{"both sides repeated", []int{-2, -2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
		{"value limits", []int{-100000, 0, 100000}, [][]int{{-100000, 0, 100000}}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := normalize(t, threeSum(slices.Clone(tt.nums)))
			want := normalize(t, tt.want)
			if !slices.EqualFunc(got, want, func(a, b []int) bool { return slices.Equal(a, b) }) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
