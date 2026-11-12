package l684

import (
	"slices"
	"testing"
)

func TestFindRedundantConnection(t *testing.T) {
	for _, tt := range []struct {
		name  string
		edges [][]int
		want  []int
	}{
		{"triangle", [][]int{{1, 2}, {1, 3}, {2, 3}}, []int{2, 3}},
		{"tail after cycle", [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}, []int{1, 4}},
		{"last closes cycle", [][]int{{1, 2}, {2, 3}, {1, 3}}, []int{1, 3}},
		{"square", [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}}, []int{1, 4}},
		{"components join", [][]int{{1, 2}, {3, 4}, {2, 3}, {1, 4}}, []int{1, 4}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.edges))
			for i := range input {
				input[i] = append([]int(nil), tt.edges[i]...)
			}
			if got := findRedundantConnection(input); !slices.Equal(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
