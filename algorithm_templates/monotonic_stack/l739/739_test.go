package l739

import (
	"slices"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	for _, tt := range []struct {
		name         string
		values, want []int
	}{
		{"example", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"ascending", []int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{"descending", []int{90, 60, 30}, []int{0, 0, 0}},
		{"equal", []int{50, 50, 50}, []int{0, 0, 0}},
		{"equal before warmer", []int{30, 30, 40}, []int{2, 1, 0}},
		{"single", []int{100}, []int{0}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := dailyTemperatures(slices.Clone(tt.values)); !slices.Equal(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
