package l2050

import "testing"

func TestMinimumTime(t *testing.T) {
	for _, tt := range []struct {
		name      string
		n         int
		relations [][]int
		times     []int
		want      int
	}{
		{"example", 3, [][]int{{1, 3}, {2, 3}}, []int{3, 2, 5}, 8},
		{"dag", 5, [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}}, []int{1, 2, 3, 4, 5}, 12},
		{"single", 1, nil, []int{7}, 7}, {"parallel", 3, nil, []int{2, 9, 5}, 9},
		{"chain", 3, [][]int{{1, 2}, {2, 3}}, []int{3, 2, 5}, 10},
		{"longest not sum", 4, [][]int{{1, 3}, {2, 3}, {3, 4}}, []int{10, 1, 2, 3}, 15},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.relations))
			for i := range input {
				input[i] = append([]int(nil), tt.relations[i]...)
			}
			if got := minimumTime(tt.n, input, append([]int(nil), tt.times...)); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
