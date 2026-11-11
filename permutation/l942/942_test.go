package l942

import "testing"

func TestDiStringMatch(t *testing.T) {
	for _, s := range []string{"IDID", "III", "DDI", "I", "D", "DDDD", "DIDI", "IIDDD"} {
		t.Run(s, func(t *testing.T) {
			got := diStringMatch(s)
			if len(got) != len(s)+1 {
				t.Fatalf("diStringMatch(%q) length = %d, want %d", s, len(got), len(s)+1)
			}
			seen := make([]bool, len(got))
			for _, v := range got {
				if v < 0 || v >= len(got) {
					t.Fatalf("result %v contains out-of-range value %d", got, v)
				}
				if seen[v] {
					t.Fatalf("result %v contains duplicate value %d", got, v)
				}
				seen[v] = true
			}
			for i := range s {
				if (s[i] == 'I' && got[i] >= got[i+1]) || (s[i] == 'D' && got[i] <= got[i+1]) {
					t.Errorf("result %v violates %c at index %d", got, s[i], i)
				}
			}
		})
	}
}
