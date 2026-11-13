package l212

import "testing"

func TestFindWords(t *testing.T) {
	tests := []struct {
		name        string
		rows        []string
		words, want []string
	}{
		{"example", []string{"oaan", "etae", "ihkr", "iflv"}, []string{"oath", "pea", "eat", "rain"}, []string{"oath", "eat"}},
		{"cannot reuse a cell", []string{"ab", "cd"}, []string{"abcb"}, nil},
		{"single matching cell", []string{"a"}, []string{"a", "b", "aa"}, []string{"a"}},
		{"no match", []string{"ab", "cd"}, []string{"xyz", "zz"}, nil},
		{"complete word can have longer match", []string{"oath"}, []string{"oat", "oath", "oaths"}, []string{"oat", "oath"}},
		{"many paths output once", []string{"aa", "aa"}, []string{"a", "aa", "aaa", "aaaa", "aaaaa"}, []string{"a", "aa", "aaa", "aaaa"}},
		{"diagonals forbidden", []string{"ab", "cd"}, []string{"ad", "bc"}, nil},
		{"vertical and reversed paths", []string{"a", "b", "c"}, []string{"abc", "cba", "ab", "ac"}, []string{"abc", "cba", "ab"}},
		{"path state restored for other searches", []string{"ab", "cd"}, []string{"ab", "ac", "bd", "cd", "dcba"}, []string{"ab", "ac", "bd", "cd"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := make([][]byte, len(tt.rows))
			for i, row := range tt.rows {
				board[i] = []byte(row)
			}
			got := findWords(board, tt.words)
			want := make(map[string]bool, len(tt.want))
			for _, word := range tt.want {
				want[word] = true
			}
			seen := make(map[string]bool, len(got))
			for _, word := range got {
				if seen[word] {
					t.Errorf("duplicate result %q", word)
				}
				seen[word] = true
				if !want[word] {
					t.Errorf("unexpected result %q", word)
				}
			}
			for word := range want {
				if !seen[word] {
					t.Errorf("missing result %q; got %v", word, got)
				}
			}
		})
	}
}
