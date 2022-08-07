package main

import "fmt"

// 0 white
// 1 black
// 2 gray
var color map[string]int

func dfs(at int, recipes []string, ingredients [][]string, supplies []string) bool {

	for _,ir := ingredients[at]{
		
	}
}
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	ret := []string{}
	for at, r := range recipes {
		if color[r] == 0 {
			color[r] = 2
			if dfs(at, recipes, ingredients, supplies) {
				color[r] = 1
				ret = append(ret, r)
			}

		}

	}
	return ret
}
func main() {
	fmt.Println(findAllRecipes([]string{"bread"}, [][]string{{"yeast", "flour"}}, []string{"yeast", "flour", "corn"}))
}
