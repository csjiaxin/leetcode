package l76

import "testing"

func TestMinWindow(t *testing.T) {
	for _, tt := range []struct{ name, s, target, want string }{
		{"example", "ADOBECODEBANC", "ABC", "BANC"}, {"single", "a", "a", "a"},
		{"insufficient copies", "a", "aa", ""}, {"duplicate target", "AAABBC", "AABC", "AABBC"},
		{"missing", "xyz", "a", ""}, {"shrink prefix", "bba", "ab", "ba"},
		{"case sensitive", "aA", "A", "A"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.s, tt.target); got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}
