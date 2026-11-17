package l1025

import "testing"

func TestDivisorGame(t *testing.T) {
	for n := 1; n <= 1000; n++ {
		want := n%2 == 0
		for call := 0; call < 2; call++ {
			if got := divisorGame(n); got != want {
				t.Fatalf("n=%d got %v want %v", n, got, want)
			}
		}
	}
}

func TestMemoGameCachesZeroValues(t *testing.T) {
	calls := map[int]int{}
	solve := MemoGame(func(n int, recur func(int) bool) bool {
		calls[n]++
		if n == 0 {
			return false
		}
		left, right := recur(n-1), recur(n-1)
		return !left && !right
	})
	for i := 0; i < 2; i++ {
		if solve(10) {
			t.Fatal("expected false")
		}
	}
	for n := 0; n <= 10; n++ {
		if calls[n] != 1 {
			t.Fatalf("state %d evaluated %d times", n, calls[n])
		}
	}
	other := MemoGame(func(int, func(int) bool) bool { return true })
	if !other(10) || solve(10) {
		t.Fatal("caches not independent")
	}
	zeroCalls := 0
	zero := MemoGame(func(int, func(int) int) int {
		zeroCalls++
		return 0
	})
	if zero(1) != 0 || zero(1) != 0 || zeroCalls != 1 {
		t.Fatal("zero value not cached")
	}
}
