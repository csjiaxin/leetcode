package l733

import (
	"reflect"
	"testing"
)

func TestFloodFillImage(t *testing.T) {
	for _, tt := range []struct {
		name        string
		image       [][]int
		r, c, color int
		want        [][]int
	}{
		{"example", [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
		{"same color", [][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0, [][]int{{0, 0, 0}, {0, 0, 0}}},
		{"diagonal", [][]int{{1, 0}, {0, 1}}, 0, 0, 2, [][]int{{2, 0}, {0, 1}}},
		{"single", [][]int{{1}}, 0, 0, 3, [][]int{{3}}},
		{"row", [][]int{{1, 1, 0, 1}}, 0, 1, 2, [][]int{{2, 2, 0, 1}}},
		{"nil", nil, 0, 0, 1, nil},
		{"empty", [][]int{}, 0, 0, 1, [][]int{}},
		{"empty row", [][]int{{}}, 0, 0, 1, [][]int{{}}},
		{"invalid start", [][]int{{1}}, -1, 0, 2, [][]int{{1}}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var before [][]int
			if tt.image != nil {
				before = make([][]int, len(tt.image))
				for r := range tt.image {
					before[r] = append(make([]int, 0, len(tt.image[r])), tt.image[r]...)
				}
			}
			for i := 0; i < 2; i++ {
				got := floodFill(tt.image, tt.r, tt.c, tt.color)
				if !reflect.DeepEqual(tt.image, before) {
					t.Fatal("input mutated")
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("got %v want %v", got, tt.want)
				}
				if len(got) > 0 && len(got[0]) > 0 {
					old := tt.image[0][0]
					got[0][0] = 99
					if tt.image[0][0] != old {
						t.Fatal("output aliases input")
					}
				}
			}
		})
	}
}
