package l76

import "testing"

func TestMinWindow(t *testing.T) {
	for _, tt := range []struct{ name, s, target, want string }{
		{"example", "ADOBECODEBANC", "ABC", "BANC"}, {"single", "a", "a", "a"},
		{"insufficient copies", "a", "aa", ""}, {"multiplicity", "AAABBC", "AABC", "AABBC"},
		{"case sensitive", "aA", "AA", ""}, {"whole", "ab", "ab", "ab"},
		{"missing", "abc", "d", ""},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.s, tt.target); got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}
