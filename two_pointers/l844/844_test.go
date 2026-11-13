package l844

import "testing"

func TestBackspaceCompare(t *testing.T) {
	for _, tt := range []struct {
		name, s, text string
		want          bool
	}{
		{"example", "ab#c", "ad#c", true}, {"both erased", "ab##", "c#d#", true},
		{"different", "a#c", "b", false}, {"leading deletes", "###a", "a", true},
		{"skip chain", "bxj##tw", "bxo#j##tw", true}, {"one erased", "a#", "a", false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.s, tt.text); got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
