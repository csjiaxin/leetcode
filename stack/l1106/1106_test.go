package l1106

import "testing"

func TestParseBoolExpr(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		// LeetCode examples
		{"example 1", "&(|(f))", false},
		{"example 2", "|(f,f,f,t)", true},
		{"example 3", "!(&(f,t))", true},

		// literals
		{"literal true", "t", true},
		{"literal false", "f", false},

		// NOT
		{"not true", "!(t)", false},
		{"not false", "!(f)", true},
		{"double not", "!(!(t))", true},
		{"triple not", "!(!(!(t)))", false},

		// AND
		{"and single true", "&(t)", true},
		{"and single false", "&(f)", false},
		{"and all true", "&(t,t,t)", true},
		{"and one false", "&(t,f,t)", false},
		{"and last false", "&(t,t,f)", false},

		// OR
		{"or single true", "|(t)", true},
		{"or single false", "|(f)", false},
		{"or all false", "|(f,f,f)", false},
		{"or first true", "|(t,f,f)", true},

		// nested / mixed
		{"or of and and not", "|(&(t,f,t),!(t))", false},
		{"and of or not literal", "&(|(f,t),!(f),t)", true},
		{"or of two ands", "|(&(t,f),&(t,t))", true},
		{"and of two ors", "&(|(f,f),|(t))", false},
		{"deeply nested", "!(&(!(&(f)),&(t),|(f,f,t)))", false},
		{"deep or chain", "|(f,|(f,|(f,|(t))))", true},
		{"deep and chain", "&(t,&(t,&(t,&(f))))", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBoolExpr(tt.in); got != tt.want {
				t.Errorf("parseBoolExpr(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
