package l630

import (
	"slices"
	"testing"
)

func TestScheduleCourse(t *testing.T) {
	for _, tt := range []struct {
		name    string
		courses [][]int
		want    int
	}{
		{"example", [][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}, 3}, {"single", [][]int{{1, 2}}, 1}, {"impossible", [][]int{{3, 2}, {4, 3}}, 0}, {"replace longest", [][]int{{5, 5}, {4, 6}, {2, 6}}, 2}, {"equal deadlines", [][]int{{3, 6}, {2, 6}, {1, 6}}, 3}, {"exact deadline", [][]int{{2, 2}}, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			before := make([][]int, len(tt.courses))
			for i := range before {
				before[i] = slices.Clone(tt.courses[i])
			}
			if got := scheduleCourse(tt.courses); got != tt.want {
				t.Errorf("%s: got %d want %d", t.Name(), got, tt.want)
			}
			for i := range before {
				if !slices.Equal(before[i], tt.courses[i]) {
					t.Fatalf("%s: input modified", t.Name())
				}
			}
		})
	}
}
