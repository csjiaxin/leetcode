package l542

import (
	"reflect"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	for _, tt := range []struct {
		name      string
		mat, want [][]int
	}{
		{"example", [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}},
		{"single", [][]int{{0}}, [][]int{{0}}},
		{"row", [][]int{{0, 1, 1, 1}}, [][]int{{0, 1, 2, 3}}},
		{"multiple sources", [][]int{{0, 1, 1, 1, 0}}, [][]int{{0, 1, 2, 1, 0}}},
		{"corner source", [][]int{{0, 1}, {1, 1}}, [][]int{{0, 1}, {1, 2}}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.mat))
			for i := range input {
				input[i] = append([]int(nil), tt.mat[i]...)
			}
			if got := updateMatrix(input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
