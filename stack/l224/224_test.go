package l224

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		// // LeetCode examples
		// {"example 1", "1 + 1", 2},
		// {"example 2", " 2-1 + 2 ", 3},
		// {"example 3", "(1+(4+5+2)-3)+(6+8)", 23},

		// // single numbers
		// {"zero", "0", 0},
		// {"single digit", "7", 7},
		// {"multi digit", "42", 42},
		// {"surrounded by spaces", "   7  ", 7},
		// {"max int32", "2147483647", 2147483647},
		// {"multi digit addition", "123+456", 579},

		// // subtraction / associativity
		// {"subtract to zero", "1-1", 0},
		// {"left associative", "10-2-3", 5},
		// {"negative result", "1-5", -4},

		// // unary minus
		// {"unary minus number", "-1", -1},
		// {"unary minus with space", "- 1 + 3", 2},
		// {"unary minus leading expr", "-2+ 1", -1},
		{"unary minus group", "-(2+3)", -5},
		{"unary minus inside parens", "1+(-2)", -1},
		{"minus negative group", "1-(-2)", 3},
		{"double unary minus", "-(-(1))", 1},
		{"unary minus nested", "- (3 + (4 + 5))", -12},

		// parentheses
		{"redundant parens", "(((5)))", 5},
		{"parens around each operand", "(1)-(2)", -1},
		{"minus distributes over group", "1-(2+3)", -4},
		{"minus flips inner minus", "1-(2-3)", 2},
		{"nested sign flips", "2-(5-(6-1))", 2},
		{"two groups", "((1+2)-(3-4))", 4},
		{"spaces everywhere", " ( 1 + ( 2 - 3 ) ) - ( 4 ) ", -4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.in); got != tt.want {
				t.Errorf("calculate(%q) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
