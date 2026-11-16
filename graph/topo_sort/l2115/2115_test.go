package l2115

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAllRecipes(t *testing.T) {
	for _, tt := range []struct {
		name           string
		recipes        []string
		ingredients    [][]string
		supplies, want []string
	}{
		{"example", []string{"bread"}, [][]string{{"yeast", "flour"}}, []string{"yeast", "flour", "corn"}, []string{"bread"}},
		{"chain", []string{"burger", "bread", "sandwich"}, [][]string{{"sandwich", "bread", "meat"}, {"yeast", "flour"}, {"bread", "meat"}}, []string{"yeast", "flour", "meat"}, []string{"bread", "sandwich", "burger"}},
		{"missing raw", []string{"bread", "cake"}, [][]string{{"flour"}, {"bread", "milk"}}, []string{"flour"}, []string{"bread"}},
		{"cycle", []string{"a", "b"}, [][]string{{"b"}, {"a"}}, []string{"water"}, []string{}},
		{"self cycle", []string{"a"}, [][]string{{"a"}}, nil, []string{}},
		{"duplicate supply", []string{"a"}, [][]string{{"x", "y"}}, []string{"x", "x"}, []string{}},
		{"duplicate edges", []string{"a"}, [][]string{{"x", "x"}}, []string{"x"}, []string{"a"}},
		{"no supplies", []string{"a"}, [][]string{{"x"}}, nil, []string{}},
		{"empty recipe", []string{"a", "b"}, [][]string{{}, {"a"}}, nil, []string{"a", "b"}},
		{"empty", nil, nil, nil, []string{}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			recipes := append([]string(nil), tt.recipes...)
			supplies := append([]string(nil), tt.supplies...)
			ingredients := make([][]string, len(tt.ingredients))
			for i := range ingredients {
				ingredients[i] = append(make([]string, 0, len(tt.ingredients[i])), tt.ingredients[i]...)
			}
			want := append([]string{}, tt.want...)
			sort.Strings(want)
			for call := 0; call < 2; call++ {
				got := findAllRecipes(tt.recipes, tt.ingredients, tt.supplies)
				sort.Strings(got)
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("got %v want %v", got, want)
				}
			}
			if !reflect.DeepEqual(recipes, tt.recipes) || !reflect.DeepEqual(supplies, tt.supplies) {
				t.Fatal("input mutated")
			}
			for i := range ingredients {
				if !reflect.DeepEqual(ingredients[i], tt.ingredients[i]) {
					t.Fatal("ingredients mutated")
				}
			}
		})
	}
}
