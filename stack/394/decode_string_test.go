package decodestring

import "testing"

func TestDecodeString(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"abc3[cd]xyz", "abccdcdcdxyz"},
		{"3[a3[b]c4[d]]", "abbbcddddabbbcddddabbbcdddd"},
		{"10[a]", "aaaaaaaaaa"},
		{"abc", "abc"},
		{"", ""},
		{"1[a]", "a"},
	}
	for _, tt := range tests {
		if got := decodeString(tt.in); got != tt.want {
			t.Errorf("decodeString(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
