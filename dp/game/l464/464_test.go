package l464

import "testing"

func TestCanIWin(t *testing.T) {
	for _, tt := range []struct {
		n, target int
		want      bool
	}{
		{4, 6, true},
		{10, 11, false}, {10, 0, true}, {10, 1, true}, {10, 10, true},
		{5, 50, false}, {1, 1, true}, {1, 2, false}, {3, 6, true}, {4, 10, false},
		{20, 210, false}, {20, 211, false},
	} {
		for call := 0; call < 2; call++ {
			if got := canIWin(tt.n, tt.target); got != tt.want {
				t.Fatalf("n=%d target=%d got %v want %v", tt.n, tt.target, got, tt.want)
			}
		}
	}
}

// 小规模直接枚举，不使用 mask 或缓存。
func bruteCanWin(n, target int) bool {
	if target <= 0 {
		return true
	}
	if n*(n+1)/2 < target {
		return false
	}
	used := make([]bool, n+1)
	var play func(int) bool
	play = func(remaining int) bool {
		for x := 1; x <= n; x++ {
			if !used[x] {
				if x >= remaining {
					return true
				}
				used[x] = true
				opponent := play(remaining - x)
				used[x] = false
				if !opponent {
					return true
				}
			}
		}
		return false
	}
	return play(target)
}
func TestCanWinAgainstBrute(t *testing.T) {
	for n := 1; n <= 7; n++ {
		for target := 0; target <= n*(n+1)/2+1; target++ {
			want := bruteCanWin(n, target)
			if got := canIWin(n, target); got != want {
				t.Fatalf("n=%d target=%d got %v want %v", n, target, got, want)
			}
		}
	}
}
