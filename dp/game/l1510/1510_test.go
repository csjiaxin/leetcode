package l1510

import "testing"

func TestWinnerSquareGame(t *testing.T) {
	for _, tt := range []struct {
		n    int
		want bool
	}{
		{0, false}, {1, true}, {2, false}, {3, true}, {4, true}, {7, false}, {17, false}, {100, true},
	} {
		for call := 0; call < 2; call++ {
			if got := winnerSquareGame(tt.n); got != tt.want {
				t.Fatalf("n=%d got %v want %v", tt.n, got, tt.want)
			}
		}
	}
	// 独立的自底向上胜负表，覆盖大量非平方状态。
	const limit = 1000
	win := make([]bool, limit+1)
	for n := 1; n <= limit; n++ {
		for root := 1; root*root <= n; root++ {
			if !win[n-root*root] {
				win[n] = true
				break
			}
		}
		if got := winnerSquareGame(n); got != win[n] {
			t.Fatalf("n=%d got %v want %v", n, got, win[n])
		}
	}
}

func TestWinnerSquareGameLarge(t *testing.T) {
	const n = 100000
	win := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		for x := 1; x*x <= i; x++ {
			if !win[i-x*x] {
				win[i] = true
				break
			}
		}
	}
	if got := winnerSquareGame(n); got != win[n] {
		t.Fatalf("got %v want %v", got, win[n])
	}
}
