package l1850

import "testing"

func TestGetMinSwaps(t *testing.T) {
	for _, tt := range []struct {
		name    string
		num     string
		k, want int
	}{
		{"example one", "5489355142", 4, 2},
		{"example two", "11112", 4, 4},
		{"leading zeros", "00123", 1, 1},
		{"two digits", "12", 1, 1},
		{"duplicate digits", "115", 1, 1},
		{"move across duplicate digits", "112", 2, 2},
		{"reverse suffix", "132", 1, 2},
		{"multiple next permutations", "123", 3, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinSwaps(tt.num, tt.k); got != tt.want {
				t.Errorf("getMinSwaps(%q, %d) = %d, want %d", tt.num, tt.k, got, tt.want)
			}
		})
	}
}
