package l51

import (
	"reflect"
	"strings"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	counts := func(boards [][]string) map[string]int {
		result := make(map[string]int)
		for _, board := range boards {
			result[strings.Join(board, "\n")]++
		}
		return result
	}
	for _, tt := range []struct {
		name string
		n    int
		want [][]string
	}{
		{"single", 1, [][]string{{"Q"}}},
		{"two impossible", 2, nil},
		{"three impossible", 3, nil},
		{"example", 4, [][]string{
			{".Q..", "...Q", "Q...", "..Q."},
			{"..Q.", "Q...", "...Q", ".Q.."},
		}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := solveNQueens(tt.n)
			if !reflect.DeepEqual(counts(got), counts(tt.want)) {
				t.Errorf("solveNQueens(%d) = %v, want %v (in any order)", tt.n, got, tt.want)
			}
		})
	}
}
