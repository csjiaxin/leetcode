package l1371

import "testing"

func TestFindTheLongestSubstring(t *testing.T) {
	for _, tt := range []struct {
		name, s string
		want    int
	}{
		{"example", "eleetminicoworoep", 13}, {"example two", "leetcodeisgreat", 5},
		{"no vowels", "bcbcbc", 6}, {"single vowel", "a", 0},
		{"even vowel", "aa", 2}, {"different vowels", "aeiou", 0},
		{"all even", "aeiouaeiou", 10},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheLongestSubstring(tt.s); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
