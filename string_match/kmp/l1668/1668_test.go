package l1668

import "testing"

func TestMaxRepeating(t *testing.T) {
	tests := []struct {
		name, sequence, word string
		want                 int
	}{
		{"example", "ababc", "ab", 2}, {"shifted", "ababc", "ba", 1},
		{"absent", "ababc", "ac", 0}, {"overlap not repetition", "aaaaa", "aa", 2},
		{"gaps break run", "abxxab", "ab", 1}, {"longer word", "a", "aa", 0},
		{"whole sequence", "abc", "abc", 1}, {"offset run", "xababab", "ab", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRepeating(tt.sequence, tt.word); got != tt.want {
				t.Errorf("maxRepeating(%q, %q) = %d, want %d", tt.sequence, tt.word, got, tt.want)
			}
		})
	}
}
