package l1642

import "testing"

func TestFurthestBuilding(t *testing.T) {
	for _, tt := range []struct {
		name                  string
		heights               []int
		bricks, ladders, want int
	}{
		{"example one", []int{4, 2, 7, 6, 9, 14, 12}, 5, 1, 4}, {"example two", []int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2, 7}, {"no ladders", []int{14, 3, 19, 3}, 17, 0, 3}, {"single", []int{1}, 0, 0, 0}, {"descending", []int{5, 4, 3}, 0, 0, 2}, {"save ladder for largest climb", []int{1, 3, 10}, 2, 1, 2}, {"cannot climb", []int{1, 2}, 0, 0, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := furthestBuilding(tt.heights, tt.bricks, tt.ladders); got != tt.want {
				t.Errorf("%s: got %d want %d", t.Name(), got, tt.want)
			}
		})
	}
}
