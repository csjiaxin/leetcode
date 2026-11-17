package l721

import (
	"reflect"
	"slices"
	"sort"
	"testing"
)

func cloneAccounts(a [][]string) [][]string {
	b := make([][]string, len(a))
	for i := range a {
		b[i] = slices.Clone(a[i])
	}
	return b
}
func normalize(a [][]string) [][]string {
	b := cloneAccounts(a)
	sort.Slice(b, func(i, j int) bool { return slices.Compare(b[i], b[j]) < 0 })
	return b
}
func TestAccountsMerge(t *testing.T) {
	for _, tt := range []struct {
		name           string
		accounts, want [][]string
	}{
		{"single", [][]string{{"A", "a@mail.com"}}, [][]string{{"A", "a@mail.com"}}},
		{"shared email", [][]string{{"A", "b@mail.com", "a@mail.com"}, {"A", "b@mail.com", "c@mail.com"}}, [][]string{{"A", "a@mail.com", "b@mail.com", "c@mail.com"}}},
		{"same name not identity", [][]string{{"A", "a@mail.com"}, {"A", "b@mail.com"}}, [][]string{{"A", "a@mail.com"}, {"A", "b@mail.com"}}},
		{"transitive", [][]string{{"A", "a@mail.com", "b@mail.com"}, {"A", "c@mail.com", "d@mail.com"}, {"A", "b@mail.com", "c@mail.com"}}, [][]string{{"A", "a@mail.com", "b@mail.com", "c@mail.com", "d@mail.com"}}},
		{"duplicate email", [][]string{{"A", "a@mail.com", "a@mail.com"}, {"B", "b@mail.com"}}, [][]string{{"A", "a@mail.com"}, {"B", "b@mail.com"}}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := accountsMerge(cloneAccounts(tt.accounts))
			for _, group := range got {
				if len(group) < 2 {
					t.Fatalf("invalid group %v", group)
				}
				if !sort.StringsAreSorted(group[1:]) {
					t.Errorf("emails not sorted: %v", group)
				}
			}
			if !reflect.DeepEqual(normalize(got), normalize(tt.want)) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
