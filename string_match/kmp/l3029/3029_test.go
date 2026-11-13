package l3029

import "testing"

func TestMinimumTimeToInitialState(t *testing.T) {
	tests := []struct {
		name, word string
		k, want    int
	}{
		{"example one", "abacaba", 3, 2}, {"example two", "abacaba", 4, 1},
		{"single", "a", 1, 1}, {"no border", "abcdef", 2, 3},
		{"round up", "abcde", 2, 3}, {"repeated", "aaaa", 2, 1},
		{"partial period", "ababa", 2, 1}, {"shorter valid border", "ababab", 3, 2},
		{"replace whole word", "abcde", 5, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTimeToInitialState(tt.word, tt.k); got != tt.want {
				t.Errorf("minimumTimeToInitialState(%q, %d) = %d, want %d", tt.word, tt.k, got, tt.want)
			}
		})
	}
}
