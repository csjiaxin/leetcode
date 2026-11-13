package l344

import "testing"

func TestReverseString(t *testing.T) {
	for _, tt := range []struct{ name, s, want string }{
		{"odd", "hello", "olleh"}, {"even", "abcd", "dcba"}, {"single", "a", "a"},
		{"equal", "aaaa", "aaaa"}, {"symbols", "a !", "! a"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := []byte(tt.s)
			reverseString(got)
			if string(got) != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}
