package l97

import "testing"

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		s3   string
		want bool
	}{
		{name: "example one", s1: "aabcc", s2: "dbbca", s3: "aadbbcbcac", want: true},
		{name: "example two", s1: "aabcc", s2: "dbbca", s3: "aadbbbaccc", want: false},
		{name: "all empty", want: true},
		{name: "only first string", s1: "abc", s3: "abc", want: true},
		{name: "only second string", s2: "abc", s3: "abc", want: true},
		{name: "first boundary mismatch", s1: "abc", s3: "acb", want: false},
		{name: "second boundary mismatch", s2: "abc", s3: "abd", want: false},
		{name: "target too short", s1: "a", s2: "b", s3: "a", want: false},
		{name: "target too long", s1: "a", s2: "b", s3: "abc", want: false},
		{name: "both sources empty but target nonempty", s3: "a", want: false},
		{name: "second source first", s1: "a", s2: "b", s3: "ba", want: true},
		{name: "alternating sources", s1: "ab", s2: "cd", s3: "acbd", want: true},
		{name: "must preserve source order", s1: "ab", s2: "cd", s3: "adbc", want: false},
		{name: "ambiguous repeated characters", s1: "aa", s2: "ab", s3: "aaba", want: true},
		{name: "repeated characters all matching", s1: "aa", s2: "aa", s3: "aaaa", want: true},
		{name: "matching counts but impossible order", s1: "ab", s2: "ab", s3: "abba", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave(tt.s1, tt.s2, tt.s3); got != tt.want {
				t.Errorf("isInterleave(%q, %q, %q) = %v, want %v", tt.s1, tt.s2, tt.s3, got, tt.want)
			}
		})
	}
}

func TestIsInterleaveExhaustiveSmallInputs(t *testing.T) {
	words := []string{"", "a", "b", "aa", "ab", "ba", "bb"}
	for _, s1 := range words {
		for _, s2 := range words {
			n := len(s1) + len(s2)
			for mask := 0; mask < 1<<n; mask++ {
				candidate := make([]byte, n)
				for i := range candidate {
					candidate[i] = 'a' + byte((mask>>i)&1)
				}
				s3 := string(candidate)
				want := interleaveBySearch(s1, s2, s3)
				if got := isInterleave(s1, s2, s3); got != want {
					t.Fatalf("isInterleave(%q, %q, %q) = %v, want %v", s1, s2, s3, got, want)
				}
			}
		}
	}
}

func interleaveBySearch(s1, s2, s3 string) bool {
	var visit func(i, j int) bool
	visit = func(i, j int) bool {
		if i+j == len(s3) {
			return true
		}
		if i < len(s1) && s1[i] == s3[i+j] && visit(i+1, j) {
			return true
		}
		return j < len(s2) && s2[j] == s3[i+j] && visit(i, j+1)
	}
	return visit(0, 0)
}
