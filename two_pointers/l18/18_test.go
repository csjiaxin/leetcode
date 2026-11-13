package l18

import (
	"slices"
	"testing"
)

func normalize(t *testing.T, values [][]int) [][]int {
	t.Helper()
	out := make([][]int, len(values))
	for i, v := range values {
		if len(v) != 4 {
			t.Fatalf("invalid quadruplet: %v", v)
		}
		out[i] = slices.Clone(v)
		slices.Sort(out[i])
	}
	slices.SortFunc(out, func(a, b []int) int { return slices.Compare(a, b) })
	return out
}

func TestFourSum(t *testing.T) {
	for _, tt := range []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{"example", []int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{"duplicates", []int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
		{"too short", []int{1, 2, 3}, 6, nil},
		{"no solution", []int{1, 2, 3, 4}, 0, nil},
		{"zeros", []int{0, 0, 0, 0, 0}, 0, [][]int{{0, 0, 0, 0}}},
		{"overflow trap", []int{1000000000, 1000000000, 1000000000, 1000000000}, -294967296, nil},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := normalize(t, fourSum(slices.Clone(tt.nums), tt.target))
			want := normalize(t, tt.want)
			if !slices.EqualFunc(got, want, func(a, b []int) bool { return slices.Equal(a, b) }) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
