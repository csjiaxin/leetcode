package l856

import "testing"

func TestScoreOfParentheses(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{"empty", "", 0},
		{"single pair", "()", 1},
		{"nested once", "(())", 2},
		{"two siblings", "()()", 2},
		{"nested with siblings", "(()(()))", 6},
		{"deep nesting", "((()))", 4},
		{"deeper nesting", "(((())))", 8},
		{"sibling groups", "(())()", 3},
		{"multiple inner siblings", "(()()())", 6},
		{"mixed", "()(())((()))", 7},
		{"complex", "(()(()()))", 10},
		{"many siblings", "()()()()()", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreOfParentheses(tt.in); got != tt.want {
				t.Errorf("scoreOfParentheses(%q) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
