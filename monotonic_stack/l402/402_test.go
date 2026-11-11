package l402

import (
	"strings"
	"testing"
)

func TestRemoveKdigits(t *testing.T) {
	for _, tt := range []struct {
		name, num string
		k         int
		want      string
	}{
		{"example one", "1432219", 3, "1219"}, {"leading zeros", "10200", 1, "200"},
		{"remove all", "10", 2, "0"}, {"increasing delete tail", "12345", 2, "123"},
		{"decreasing", "54321", 2, "321"}, {"equal digits", "1111", 2, "11"},
		{"all remaining zeros", "1000", 1, "0"}, {"single", "9", 1, "0"},
		{"multiple leading zeros", "100200", 1, "200"}, {"zero result", "10", 1, "0"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits(tt.num, tt.k); got != tt.want {
				t.Errorf("%s: got %q, want %q", t.Name(), got, tt.want)
			}
		})
	}
}
func TestRemoveKdigitsMaximum(t *testing.T) {
	num := strings.Repeat("9", 100000)
	want := strings.Repeat("9", 50000)
	if got := removeKdigits(num, 50000); got != want {
		t.Errorf("%s: result mismatch, length %d want %d", t.Name(), len(got), len(want))
	}
}
