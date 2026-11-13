package l345

import "testing"

func TestReverseVowels(t *testing.T) {
	for _, tt := range []struct{ name, s, want string }{
		{"example", "hello", "holle"}, {"leetcode", "leetcode", "leotcede"},
		{"case", "aA", "Aa"}, {"none", "bcdf", "bcdf"}, {"single", "a", "a"}, {"all", "aeiou", "uoiea"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.s); got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}
