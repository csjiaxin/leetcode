package l125

import "testing"

func TestIsPalindrome(t *testing.T) {
	for _, tt := range []struct {
		name, s string
		want    bool
	}{
		{"example", "A man, a plan, a canal: Panama", true}, {"false", "race a car", false},
		{"space", " ", true}, {"digit letter", "0P", false}, {"single", "a", true},
		{"digits", "12,21", true}, {"punctuation", ".,!", true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
