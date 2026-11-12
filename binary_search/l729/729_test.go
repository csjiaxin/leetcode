package l729

import "testing"

func TestBook(t *testing.T) {
	for _, tt := range []struct {
		name     string
		bookings [][2]int
		want     []bool
	}{
		{"example", [][2]int{{10, 20}, {15, 25}, {20, 30}}, []bool{true, false, true}},
		{"out of order", [][2]int{{30, 40}, {10, 20}, {20, 30}, {0, 10}}, []bool{true, true, true, true}},
		{"overlap neighbors", [][2]int{{10, 20}, {30, 40}, {19, 31}, {20, 31}, {19, 30}, {20, 30}}, []bool{true, true, false, false, false, true}},
		{"same start", [][2]int{{10, 20}, {10, 15}, {10, 20}, {5, 25}, {12, 18}}, []bool{true, false, false, false, false}},
		{"bounds", [][2]int{{0, 1}, {1, 1000000000}, {0, 1000000000}}, []bool{true, true, false}},
		{"invalid", [][2]int{{2, 2}, {3, 1}, {1, 3}}, []bool{false, false, true}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			calendar := Constructor()
			for i, b := range tt.bookings {
				if got := calendar.Book(b[0], b[1]); got != tt.want[i] {
					t.Errorf("booking %d %v got %v want %v", i, b, got, tt.want[i])
				}
			}
		})
	}
}
