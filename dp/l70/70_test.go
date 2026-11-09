package l70

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "single stair", n: 1, want: 1},
		{name: "example one", n: 2, want: 2},
		{name: "example two", n: 3, want: 3},
		{name: "four stairs", n: 4, want: 5},
		{name: "five stairs", n: 5, want: 8},
		{name: "ten stairs", n: 10, want: 89},
		{name: "twenty stairs", n: 20, want: 10946},
		{name: "thirty stairs", n: 30, want: 1346269},
		{name: "forty stairs", n: 40, want: 165580141},
		{name: "maximum input", n: 45, want: 1836311903},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.n); got != tt.want {
				t.Errorf("climbStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
