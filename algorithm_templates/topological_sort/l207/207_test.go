package l207

import "testing"

func TestCanFinish(t *testing.T) {
	for _, tt := range []struct {
		name          string
		n             int
		prerequisites [][]int
		want          bool
	}{
		{"edge", 2, [][]int{{1, 0}}, true}, {"cycle", 2, [][]int{{1, 0}, {0, 1}}, false},
		{"isolated", 4, nil, true}, {"diamond", 4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, true},
		{"cycle after root", 4, [][]int{{1, 0}, {2, 1}, {1, 2}, {3, 2}}, false},
		{"single", 1, nil, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.prerequisites))
			for i := range input {
				input[i] = append([]int(nil), tt.prerequisites[i]...)
			}
			if got := canFinish(tt.n, input); got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
