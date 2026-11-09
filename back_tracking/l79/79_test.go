package l79

import "testing"

func TestExist(t *testing.T) {
	tests := []struct {
		name string
		rows []string
		word string
		want bool
	}{
		{
			name: "example one",
			rows: []string{"ABCE", "SFCS", "ADEE"},
			word: "ABCCED",
			want: true,
		},
		{
			name: "example two",
			rows: []string{"ABCE", "SFCS", "ADEE"},
			word: "SEE",
			want: true,
		},
		{
			name: "example three cannot reuse a cell",
			rows: []string{"ABCE", "SFCS", "ADEE"},
			word: "ABCB",
			want: false,
		},
		{
			name: "single matching cell",
			rows: []string{"a"},
			word: "a",
			want: true,
		},
		{
			name: "single nonmatching cell",
			rows: []string{"a"},
			word: "b",
			want: false,
		},
		{
			name: "single row reverse direction",
			rows: []string{"ABCD"},
			word: "DCBA",
			want: true,
		},
		{
			name: "single column upward direction",
			rows: []string{"A", "B", "C"},
			word: "CBA",
			want: true,
		},
		{
			name: "diagonal moves are forbidden",
			rows: []string{"AX", "XB"},
			word: "AB",
			want: false,
		},
		{
			name: "word longer than board",
			rows: []string{"AA", "AA"},
			word: "AAAAA",
			want: false,
		},
		{
			name: "repeated letters in distinct cells",
			rows: []string{"AA", "AA"},
			word: "AAAA",
			want: true,
		},
		{
			name: "try later starting cell",
			rows: []string{"AXAB"},
			word: "AB",
			want: true,
		},
		{
			name: "backtrack after a dead end",
			rows: []string{"ABCE", "SFES", "ADEE"},
			word: "ABCESEEEFS",
			want: true,
		},
		{
			name: "case sensitive",
			rows: []string{"aB"},
			word: "AB",
			want: false,
		},
		{
			name: "horizontal edges do not wrap",
			rows: []string{"XA", "BX"},
			word: "AB",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := make([][]byte, len(tt.rows))
			for i, row := range tt.rows {
				board[i] = []byte(row)
			}
			if got := exist(board, tt.word); got != tt.want {
				t.Errorf("exist(%q, %q) = %v, want %v", tt.rows, tt.word, got, tt.want)
			}
		})
	}
}
