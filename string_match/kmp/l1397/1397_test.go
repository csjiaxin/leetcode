package l1397

import "testing"

func TestFindGoodStrings(t *testing.T) {
	tests := []struct {
		name         string
		n            int
		s1, s2, evil string
		want         int
	}{
		{"example one", 2, "aa", "da", "b", 51},
		{"example two", 8, "leetcode", "leetgoes", "leet", 0},
		{"example three", 2, "gx", "gz", "x", 2},
		{"single alphabet", 1, "a", "z", "b", 25},
		{"both bounds inclusive", 2, "aa", "ab", "aa", 1},
		{"one allowed string", 3, "abc", "abc", "bc", 0},
		{"one good string", 3, "abc", "abc", "aa", 1},
		{"evil longer than strings", 1, "a", "c", "aa", 3},
		{"overlapping forbidden", 3, "aaa", "aaz", "aa", 0},
		{"all length two except aa", 2, "aa", "zz", "aa", 675},
		{"modulo count", 7, "aaaaaaa", "zzzzzzz", "a", 103515583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findGoodStrings(tt.n, tt.s1, tt.s2, tt.evil); got != tt.want {
				t.Errorf("findGoodStrings(%d, %q, %q, %q) = %d, want %d", tt.n, tt.s1, tt.s2, tt.evil, got, tt.want)
			}
		})
	}
}
