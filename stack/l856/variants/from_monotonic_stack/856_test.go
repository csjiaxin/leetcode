package l856

import (
	"strings"
	"testing"
)

func TestScoreOfParentheses(t *testing.T) {
	for _, tt := range []struct {
		name, s string
		want    int
	}{
		{"pair", "()", 1}, {"nested", "(())", 2}, {"concatenated", "()()", 2},
		{"mixed", "(()(()))", 6}, {"three nested", "((()))", 4},
		{"sum then double", "(()())", 4}, {"multiple groups", "()(())", 3},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreOfParentheses(tt.s); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
func TestScoreOfParenthesesMaximum(t *testing.T) {
	for _, tt := range []struct {
		name, s string
		want    int
	}{
		{"maximum nesting", strings.Repeat("(", 25) + strings.Repeat(")", 25), 1 << 24},
		{"maximum flat", strings.Repeat("()", 25), 25},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreOfParentheses(tt.s); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
