package l1092

import "testing"

func TestShortestCommonSupersequence(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want int
	}{
		{name: "official overlapping", a: "abac", b: "cab", want: 5},
		{name: "official identical", a: "aaaaaaaa", b: "aaaaaaaa", want: 8},
		{name: "both empty", a: "", b: "", want: 0},
		{name: "first empty", a: "", b: "abc", want: 3},
		{name: "second empty", a: "abc", b: "", want: 3},
		{name: "disjoint", a: "abc", b: "def", want: 6},
		{name: "first contained", a: "ac", b: "abc", want: 3},
		{name: "second contained", a: "abc", b: "ac", want: 3},
		{name: "multiple optimal answers", a: "ab", b: "ba", want: 3},
		{name: "repeated letters", a: "aab", b: "aba", want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortestCommonSupersequence(tt.a, tt.b)
			if !isSubsequence(tt.a, got) || !isSubsequence(tt.b, got) {
				t.Errorf("result %q does not contain both %q and %q as subsequences", got, tt.a, tt.b)
			}
			if len(got) != tt.want {
				t.Errorf("result %q has length %d; want optimal length %d", got, len(got), tt.want)
			}
		})
	}
}

// Exhaust all pairs of binary strings of length at most four. Compute the
// independent optimal-length oracle by brute-force LCS subset enumeration,
// rather than repeating the solution's dynamic-programming recurrence.
func TestShortestCommonSupersequenceExhaustive(t *testing.T) {
	words := []string{""}
	for length := 1; length <= 4; length++ {
		for bits := 0; bits < 1<<length; bits++ {
			word := make([]byte, length)
			for i := range word {
				word[i] = 'a' + byte(bits>>i&1)
			}
			words = append(words, string(word))
		}
	}

	for _, a := range words {
		for _, b := range words {
			got := shortestCommonSupersequence(a, b)
			want := len(a) + len(b) - bruteLCSLength(a, b)
			if !isSubsequence(a, got) || !isSubsequence(b, got) || len(got) != want {
				t.Fatalf("inputs (%q, %q): result %q; want both subsequences and length %d", a, b, got, want)
			}
		}
	}
}

func isSubsequence(needle string, haystack string) bool {
	i := 0
	for j := 0; j < len(haystack) && i < len(needle); j++ {
		if needle[i] == haystack[j] {
			i++
		}
	}
	return i == len(needle)
}

func bruteLCSLength(a string, b string) int {
	best := 0
	for mask := 0; mask < 1<<len(a); mask++ {
		candidate := make([]byte, 0, len(a))
		for i := 0; i < len(a); i++ {
			if mask>>i&1 != 0 {
				candidate = append(candidate, a[i])
			}
		}
		if len(candidate) > best && isSubsequence(string(candidate), b) {
			best = len(candidate)
		}
	}
	return best
}
