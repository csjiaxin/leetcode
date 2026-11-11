package l2402

import "testing"

func TestMostBooked(t *testing.T) {
	for _, tt := range []struct {
		name     string
		n        int
		meetings [][]int
		want     int
	}{
		{"example one", 2, [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}, 0}, {"example two", 3, [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}, 1}, {"simultaneous release and tie", 2, [][]int{{0, 5}, {1, 5}, {5, 6}, {6, 7}}, 0}, {"unsorted", 2, [][]int{{2, 3}, {0, 10}, {1, 2}}, 1}, {"single room", 1, [][]int{{0, 10}, {1, 2}, {3, 5}}, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostBooked(tt.n, tt.meetings); got != tt.want {
				t.Errorf("%s: got %d want %d", t.Name(), got, tt.want)
			}
		})
	}
}
func TestMostBookedLargeDelayedTime(t *testing.T) {
	// With two rooms and equal long durations, assignments alternate. The last
	// meeting completes well beyond int32; 10001 meetings give room zero one more.
	meetings := make([][]int, 10001)
	for i := range meetings {
		meetings[i] = []int{i, i + 480000}
	}
	if got := mostBooked(2, meetings); got != 0 {
		t.Errorf("%s: got %d want 0", t.Name(), got)
	}
}
