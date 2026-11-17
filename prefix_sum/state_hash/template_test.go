package prefixstate

import "testing"

func TestCountPrefixStateAdapters(t *testing.T) {
	add := func(state, value int) int { return state + value }
	identity := func(state int) int { return state }
	for _, tc := range []struct {
		values []int
		k      int
		want   int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, -1, 1}, 1, 3},
		{[]int{0, 0, 0}, 0, 6},
		{nil, 0, 0},
	} {
		got := Count(tc.values, 0, add, func(state int) int { return state - tc.k })
		if got != tc.want {
			t.Errorf("sum %v target %d: got %d, want %d", tc.values, tc.k, got, tc.want)
		}
	}
	got := Count([]int{4, 5, 0, -2, -3, 1}, 0,
		func(state, value int) int { return ((state+value)%5 + 5) % 5 }, identity)
	if got != 7 {
		t.Errorf("remainder adapter: got %d, want 7", got)
	}
	got = Count([]int{1, 1, 2, 1, 1}, 0,
		func(state, value int) int { return state + value%2 },
		func(state int) int { return state - 3 })
	if got != 2 {
		t.Errorf("odd-count adapter: got %d, want 2", got)
	}
}

func TestLongestPrefixStateAdapters(t *testing.T) {
	identity := func(state int) int { return state }
	got := Longest([]int{0, 1, 0}, 0,
		func(state, value int) int { return state + 2*value - 1 }, identity)
	if got != 2 {
		t.Errorf("balanced binary adapter: got %d, want 2", got)
	}
	got = Longest([]byte("eleetminicoworoep"), 0,
		func(state int, value byte) int {
			for i, vowel := range []byte("aeiou") {
				if value == vowel {
					return state ^ (1 << i)
				}
			}
			return state
		}, identity)
	if got != 13 {
		t.Errorf("parity-mask adapter: got %d, want 13", got)
	}
	got = Longest([]int{1, -1, 1, -1}, 0,
		func(state, value int) int { return state + value }, identity)
	if got != 4 {
		t.Errorf("earliest prefix must survive repeated states: got %d, want 4", got)
	}
	if got := Longest([]int(nil), 0, func(state, value int) int { return state + value }, identity); got != 0 {
		t.Errorf("empty input: got %d, want 0", got)
	}
}
