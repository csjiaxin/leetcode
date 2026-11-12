package l2498

import "testing"

func TestMaxJump(t *testing.T) {
	for _, tt := range []struct {
		name   string
		stones []int
		want   int
	}{
		{"example", []int{0, 2, 5, 6, 7}, 5}, {"example two", []int{0, 3, 9}, 9},
		{"two", []int{0, 7}, 7}, {"uniform", []int{0, 1, 2, 3, 4}, 2},
		{"large gap", []int{0, 1, 2, 100}, 99}, {"large positions", []int{0, 1000000000}, 1000000000},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxJump(tt.stones); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestAgainstPathPartitions(t *testing.T) {
	// Exhaust all assignments of intermediate stones to outward/return paths.
	for n := 2; n <= 8; n++ {
		stones := make([]int, n)
		for i := 1; i < n; i++ {
			stones[i] = stones[i-1] + i%3 + 1
		}
		want := stones[n-1]
		for mask := 0; mask < 1<<(n-2); mask++ {
			last := [2]int{0, 0}
			cost := 0
			for i := 1; i < n-1; i++ {
				path := (mask >> (i - 1)) & 1
				cost = max(cost, stones[i]-stones[last[path]])
				last[path] = i
			}
			cost = max(cost, stones[n-1]-stones[last[0]], stones[n-1]-stones[last[1]])
			want = min(want, cost)
		}
		if got := maxJump(stones); got != want {
			t.Fatalf("stones=%v got=%d want=%d", stones, got, want)
		}
	}
}
