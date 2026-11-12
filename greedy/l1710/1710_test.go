package l1710

import "testing"

func TestMaximumUnits(t *testing.T) {
	for _, tt := range []struct {
		name       string
		boxes      [][]int
		size, want int
	}{
		{"example one", [][]int{{1, 3}, {2, 2}, {3, 1}}, 4, 8},
		{"example two", [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10, 91},
		{"partial type", [][]int{{5, 4}}, 2, 8},
		{"more capacity than boxes", [][]int{{1, 3}, {2, 2}}, 10, 7},
		{"unit value not box count", [][]int{{10, 1}, {1, 100}}, 1, 100},
		{"equal unit values", [][]int{{2, 3}, {4, 3}}, 3, 9},
		{"single box", [][]int{{1, 1}}, 1, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			boxes := make([][]int, len(tt.boxes))
			for i, b := range tt.boxes {
				boxes[i] = append([]int(nil), b...)
			}
			if got := maximumUnits(boxes, tt.size); got != tt.want {
				t.Errorf("maximumUnits(%v, %d) = %d, want %d", tt.boxes, tt.size, got, tt.want)
			}
		})
	}
}
