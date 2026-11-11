package l2375

import "testing"

func TestSmallestNumber(t *testing.T) {
	for _, tt := range []struct {
		pattern, want string
	}{
		{"IIIDIDDD", "123549876"},
		{"DDD", "4321"},
		{"I", "12"},
		{"D", "21"},
		{"III", "1234"},
		{"IDID", "13254"},
		{"IIDDD", "126543"},
		{"DII", "2134"},
		{"IIIIIIII", "123456789"},
		{"DDDDDDDD", "987654321"},
	} {
		t.Run(tt.pattern, func(t *testing.T) {
			if got := smallestNumber(tt.pattern); got != tt.want {
				t.Errorf("smallestNumber(%q) = %q, want %q", tt.pattern, got, tt.want)
			}
		})
	}
}
