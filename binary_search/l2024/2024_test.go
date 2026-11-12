package l2024

import "testing"

func TestMaxConsecutiveAnswers(t *testing.T) {
	for _, tt := range []struct {
		name, key string
		k, want   int
	}{
		{"example", "TTFF", 2, 4}, {"example two", "TFFT", 1, 3},
		{"example three", "TTFTTFTT", 1, 5}, {"no changes", "TFTTTF", 0, 3},
		{"all equal", "FFFF", 0, 4}, {"single", "T", 0, 1}, {"empty", "", 0, 0},
		{"full budget", "TFTF", 4, 4},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxConsecutiveAnswers(tt.key, tt.k); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestExhaustiveKeys(t *testing.T) {
	for n := 1; n <= 8; n++ {
		for mask := 0; mask < 1<<n; mask++ {
			key := make([]byte, n)
			for i := range key {
				key[i] = 'F'
				if mask&(1<<i) != 0 {
					key[i] = 'T'
				}
			}
			for k := 0; k <= n; k++ {
				want := 0
				for l := 0; l < n; l++ {
					trues := 0
					for r := l; r < n; r++ {
						if key[r] == 'T' {
							trues++
						}
						length := r - l + 1
						if min(trues, length-trues) <= k {
							want = max(want, length)
						}
					}
				}
				if got := maxConsecutiveAnswers(string(key), k); got != want {
					t.Fatalf("key=%s k=%d got=%d want=%d", key, k, got, want)
				}
			}
		}
	}
}
