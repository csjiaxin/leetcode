package l60

import "testing"

func TestGetPermutation(t *testing.T) {
	for _, tt := range []struct {
		name string
		n, k int
		want string
	}{
		{"example one", 3, 3, "213"},
		{"example two", 4, 9, "2314"},
		{"first", 3, 1, "123"},
		{"last", 3, 6, "321"},
		{"single", 1, 1, "1"},
		{"factorial boundary", 4, 6, "1432"},
		{"next factorial block", 4, 7, "2134"},
		{"maximum n first", 9, 1, "123456789"},
		{"maximum n last", 9, 362880, "987654321"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPermutation(tt.n, tt.k); got != tt.want {
				t.Errorf("getPermutation(%d, %d) = %q, want %q", tt.n, tt.k, got, tt.want)
			}
		})
	}
}
