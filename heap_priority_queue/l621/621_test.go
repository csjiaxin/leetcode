package l621

import (
	"math/rand"
	"testing"
)

func TestLeastInterval(t *testing.T) {
	for _, tt := range []struct {
		name, s string
		n, want int
	}{
		{"example", "AAABBB", 2, 8}, {"no cooldown", "AAABBB", 0, 6}, {"mixed", "ACABDB", 1, 6}, {"long cooldown", "AAABBB", 3, 10}, {"single kind", "AAAA", 2, 10}, {"all distinct", "ABC", 100, 3}, {"single", "A", 100, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval([]byte(tt.s), tt.n); got != tt.want {
				t.Errorf("%s: got %d want %d", t.Name(), got, tt.want)
			}
		})
	}
}
func TestLeastIntervalRandom(t *testing.T) {
	r := rand.New(rand.NewSource(621))
	for trial := 0; trial < 300; trial++ {
		tasks := make([]byte, r.Intn(100)+1)
		counts := [6]int{}
		for i := range tasks {
			x := r.Intn(6)
			counts[x]++
			tasks[i] = byte('A' + x)
		}
		highest, tied := 0, 0
		for _, c := range counts {
			if c > highest {
				highest = c
				tied = 1
			} else if c == highest {
				tied++
			}
		}
		n := r.Intn(8)
		want := max(len(tasks), (highest-1)*(n+1)+tied)
		if got := leastInterval(tasks, n); got != want {
			t.Fatalf("%s: trial %d got %d want %d", t.Name(), trial, got, want)
		}
	}
}
