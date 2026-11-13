package l1268

import "testing"

func TestSuggestedProducts(t *testing.T) {
	tests := []struct {
		name     string
		products []string
		search   string
		want     [][]string
	}{
		{"example", []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse", [][]string{
			{"mobile", "moneypot", "monitor"}, {"mobile", "moneypot", "monitor"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}, {"mouse", "mousepad"},
		}},
		{"single product", []string{"havana"}, "havana", [][]string{{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}}},
		{"sort and limit to three", []string{"az", "ab", "a", "aa", "ac"}, "a", [][]string{{"a", "aa", "ab"}}},
		{"missing prefix stays missing", []string{"a", "ab"}, "azb", [][]string{{"a", "ab"}, {}, {}}},
		{"no first character", []string{"apple", "banana"}, "zoo", [][]string{{}, {}, {}}},
		{"exact word and longer words", []string{"appetite", "apple", "app"}, "apple", [][]string{
			{"app", "appetite", "apple"}, {"app", "appetite", "apple"}, {"app", "appetite", "apple"}, {"apple"}, {"apple"},
		}},
		{"two candidates", []string{"zebra", "zen"}, "ze", [][]string{{"zebra", "zen"}, {"zebra", "zen"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Sorting products in place is permitted by the original problem.
			products := append([]string(nil), tt.products...)
			got := suggestedProducts(products, tt.search)
			if len(got) != len(tt.want) {
				t.Fatalf("got %d prefixes, want %d: %v", len(got), len(tt.want), got)
			}
			for i, want := range tt.want {
				if len(got[i]) != len(want) {
					t.Errorf("prefix %q = %v, want %v", tt.search[:i+1], got[i], want)
					continue
				}
				for j := range want {
					if got[i][j] != want[j] {
						t.Errorf("prefix %q = %v, want %v", tt.search[:i+1], got[i], want)
						break
					}
				}
			}
		})
	}
}
