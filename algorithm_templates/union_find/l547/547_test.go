package l547

import "testing"

func TestFindCircleNum(t *testing.T) {
	for _, tt := range []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{"example", [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, 2},
		{"isolated", [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, 3},
		{"single", [][]int{{1}}, 1}, {"connected", [][]int{{1, 1}, {1, 1}}, 1},
		{"transitive", [][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}}, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.matrix))
			for i := range input {
				input[i] = append([]int(nil), tt.matrix[i]...)
			}
			if got := findCircleNum(input); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
