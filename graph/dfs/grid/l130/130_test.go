package l130

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	for _, tt := range []struct{ rows, want []string }{
		{nil, nil}, {[]string{}, []string{}}, {[]string{""}, []string{""}},
		{[]string{"XXXX", "XOOX", "XXOX", "XOXX"}, []string{"XXXX", "XXXX", "XXXX", "XOXX"}},
		{[]string{"O"}, []string{"O"}}, {[]string{"X"}, []string{"X"}},
		{[]string{"OOOO", "OOOO"}, []string{"OOOO", "OOOO"}},
		{[]string{"OXO"}, []string{"OXO"}}, {[]string{"O", "X", "O"}, []string{"O", "X", "O"}},
		{[]string{"OXX", "XOX", "XXO"}, []string{"OXX", "XXX", "XXO"}},
		{[]string{"XOXX", "XOOX", "XXOX", "XXXX"}, []string{"XOXX", "XOOX", "XXOX", "XXXX"}},
	} {
		var board, want [][]byte
		if tt.rows != nil {
			board = make([][]byte, len(tt.rows))
			want = make([][]byte, len(tt.want))
			for r := range tt.rows {
				board[r] = []byte(tt.rows[r])
				want[r] = []byte(tt.want[r])
			}
		}
		for i := 0; i < 2; i++ {
			got := solve(board)
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("%v: got %q want %q", tt.rows, got, want)
			}
			for r, s := range tt.rows {
				if string(board[r]) != s {
					t.Fatal("input mutated")
				}
			}
			if len(got) > 0 && len(got[0]) > 0 {
				old := board[0][0]
				got[0][0] = '!'
				if board[0][0] != old {
					t.Fatal("output aliases input")
				}
			}
		}
	}
}
