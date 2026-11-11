package l767

import "testing"

func TestReorganizeString(t *testing.T) {
	for _, tt := range []struct {
		name, s  string
		possible bool
	}{
		{"example", "aab", true}, {"impossible", "aaab", false}, {"single", "a", true}, {"distinct", "abc", true}, {"balanced", "aaabbbccc", true}, {"limit", "aaaabbc", true}, {"over limit", "aaaaabc", false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := reorganizeString(tt.s)
			if !tt.possible {
				if got != "" {
					t.Errorf("%s: want empty got %q", t.Name(), got)
				}
				return
			}
			if len(got) != len(tt.s) {
				t.Fatalf("%s: length mismatch", t.Name())
			}
			counts := map[byte]int{}
			for i := range tt.s {
				counts[tt.s[i]]++
			}
			for i := range got {
				counts[got[i]]--
				if i > 0 && got[i] == got[i-1] {
					t.Fatalf("%s: adjacent duplicate in %q", t.Name(), got)
				}
			}
			for _, c := range counts {
				if c != 0 {
					t.Fatalf("%s: changed character counts", t.Name())
				}
			}
		})
	}
}
