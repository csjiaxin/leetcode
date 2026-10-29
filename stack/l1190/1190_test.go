package l1190

import "testing"

func TestReverseParentheses(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		// LeetCode examples
		{"example 1", "(abcd)", "dcba"},
		{"example 2", "(u(love)i)", "iloveu"},
		{"example 3", "(ed(et(oc))el)", "leetcode"},
		{"example 4", "a(bcdefghijkl(mno)p)q", "apmnolkjihgfedcbq"},

		// edge cases
		{"empty", "", ""},
		{"no parentheses", "abc", "abc"},
		{"single char", "a", "a"},
		{"single char in parens", "(a)", "a"},
		{"empty parens", "()", ""},
		{"empty parens between letters", "a()b", "ab"},
		{"nested empty parens", "(())", ""},

		// nesting
		{"double nesting cancels", "((ab))", "ab"},
		{"triple nesting reverses", "(((abc)))", "cba"},
		{"inner then outer", "((a)b)", "ba"},
		{"outer with inner on right", "(a(bc))", "bca"},

		// siblings
		{"adjacent groups", "(ab)(cd)", "badc"},
		{"groups with surrounding text", "x(ab)y(cd)z", "xbaydcz"},
		{"siblings inside group", "((ab)(cd))", "cdab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseParentheses(tt.in); got != tt.want {
				t.Errorf("reverseParentheses(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
