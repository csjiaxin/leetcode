package l1854

import "testing"

func TestMaximumPopulation(t *testing.T) {
	for _, tt := range []struct {
		name string
		logs [][]int
		want int
	}{
		{"disjoint tie", [][]int{{1993, 1999}, {2000, 2010}}, 1993},
		{"example", [][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}}, 1960},
		{"death excluded", [][]int{{1950, 1960}, {1960, 1970}}, 1950},
		{"single", [][]int{{2001, 2002}}, 2001},
		{"unsorted overlap", [][]int{{2000, 2010}, {1990, 2005}, {2003, 2004}}, 2003},
		{"earliest peak", [][]int{{1950, 1952}, {1951, 1953}, {1952, 1954}}, 1951},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.logs))
			for i := range input {
				input[i] = append([]int(nil), tt.logs[i]...)
			}
			if got := maximumPopulation(input); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
