package l2402

import (
	"math/rand"
	"sort"
	"testing"
)

func TestMostBookedRandom(t *testing.T) {
	r := rand.New(rand.NewSource(2402))
	for trial := 0; trial < 500; trial++ {
		n, m := r.Intn(5)+1, r.Intn(40)+1
		meetings := make([][]int, m)
		for i := range meetings {
			meetings[i] = []int{2 * i, 2*i + r.Intn(30) + 1}
		}
		end := make([]int64, n)
		counts := make([]int, n)
		for _, meeting := range meetings {
			chosen := -1
			start := int64(meeting[0])
			for id := 0; id < n; id++ {
				if end[id] <= start {
					chosen = id
					break
				}
			}
			if chosen < 0 {
				chosen = 0
				for id := 1; id < n; id++ {
					if end[id] < end[chosen] {
						chosen = id
					}
				}
			}
			end[chosen] = max(start, end[chosen]) + int64(meeting[1]-meeting[0])
			counts[chosen]++
		}
		want := 0
		for id := 1; id < n; id++ {
			if counts[id] > counts[want] {
				want = id
			}
		}
		r.Shuffle(m, func(i, j int) { meetings[i], meetings[j] = meetings[j], meetings[i] })
		if got := mostBooked(n, meetings); got != want {
			t.Fatalf("%s: trial %d got %d want %d", t.Name(), trial, got, want)
		}
		// Verify the implementation did not reorder the caller's meetings.
		before := append([][]int(nil), meetings...)
		mostBooked(n, meetings)
		for i := range before {
			if before[i][0] != meetings[i][0] {
				t.Fatalf("%s: input reordered", t.Name())
			}
		}
	}
}

func TestMostBookedOverflowOracle(t *testing.T) {
	// Independently scan two room clocks rather than relying on an assignment pattern.
	meetings := make([][]int, 12000)
	for i := range meetings {
		meetings[i] = []int{i, 500000}
	}
	end := [2]int64{}
	counts := [2]int{}
	for _, m := range meetings {
		id := 0
		if end[1] < end[0] {
			id = 1
		}
		end[id] = max(end[id], int64(m[0])) + int64(m[1]-m[0])
		counts[id]++
	}
	want := 0
	if counts[1] > counts[0] {
		want = 1
	}
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][0] > meetings[j][0] })
	if got := mostBooked(2, meetings); got != want {
		t.Fatalf("%s: got %d want %d", t.Name(), got, want)
	}
}
