package l200

import (
	"reflect"
	"testing"
)

func TestNumIslands(t *testing.T) {
	for _, tt := range []struct {
		rows []string
		want int
	}{
		{nil, 0}, {[]string{""}, 0}, {[]string{"0"}, 0}, {[]string{"1"}, 1},
		{[]string{"11110", "11010", "11000", "00000"}, 1},
		{[]string{"11000", "11000", "00100", "00011"}, 3},
		{[]string{"10", "01"}, 2}, {[]string{"11011"}, 2},
		{[]string{"1", "0", "1"}, 2}, {[]string{"111", "101", "111"}, 1},
	} {
		grid := make([][]byte, len(tt.rows))
		for r, s := range tt.rows {
			grid[r] = []byte(s)
		}
		for i := 0; i < 2; i++ {
			if got := numIslands(grid); got != tt.want {
				t.Fatalf("%v: got %d want %d", tt.rows, got, tt.want)
			}
			for r, s := range tt.rows {
				if !reflect.DeepEqual(grid[r], []byte(s)) {
					t.Fatal("input mutated")
				}
			}
		}
	}
}
