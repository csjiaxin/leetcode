package l1834

import (
	"slices"
	"testing"
)

func TestGetOrder(t *testing.T) {
	for _, tt := range []struct {
		name  string
		tasks [][]int
		want  []int
	}{
		{"example one", [][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}, []int{0, 2, 3, 1}}, {"same enqueue", [][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}, []int{4, 3, 2, 0, 1}}, {"idle jump", [][]int{{1, 1}, {10, 2}, {10, 1}}, []int{0, 2, 1}}, {"index tie", [][]int{{1, 3}, {1, 3}, {1, 3}}, []int{0, 1, 2}}, {"wide clock", [][]int{{1, 1000000000}, {1, 1000000000}, {1, 1000000000}, {1000000000, 1}}, []int{0, 3, 1, 2}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOrder(tt.tasks); !slices.Equal(got, tt.want) {
				t.Errorf("%s: got %v want %v", t.Name(), got, tt.want)
			}
		})
	}
}
