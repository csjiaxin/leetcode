package l202

import "testing"

func TestIsHappy(t *testing.T) {
	for _, tt := range []struct {
		name string
		n    int
		want bool
	}{
		{"example", 19, true}, {"unhappy", 2, false}, {"one", 1, true},
		{"ten", 10, true}, {"cycle", 4, false}, {"seven", 7, true}, {"max int32", 2147483647, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.n); got != tt.want {
				t.Errorf("isHappy(%d)=%v want %v", tt.n, got, tt.want)
			}
		})
	}
}
