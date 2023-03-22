package main

import "fmt"

type Color int

const (
	White Color = iota
	Gray
	Black
)

var visit map[string]bool
var ret bool
var retArr []string

type AdjType map[string]map[string]bool

var adj AdjType

func insert_edge(adj AdjType, from, to string) {
	if v, ok := adj[from]; ok {
		v[to] = true
	} else {
		adj[from] = make(map[string]bool)
		adj[from][to] = true
	}

}

// func buildAdj(edges [][]int) AdjType {
// 	adj := make(AdjType)
// 	for _, e := range edges {
// 		insert_edge(adj, e[0], e[1])
// 		// insert_edge(adj, e[1], e[0])
// 	}
// 	return adj
// }

func dfs(at string, supplies, recipesMap map[string]bool) bool {
	if v, ok := visit[at]; ok {
		return v
	}
	visit[at] = false
	cook := true
	for to, _ := range adj[at] {
		if _, ok := supplies[to]; ok {
			continue
		}
		if _, ok := recipesMap[to]; ok {
			if dfs(to, supplies, recipesMap) {
				continue
			}
		}
		cook = false
		break
	}
	visit[at] = cook
	return cook
}
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	adj = make(AdjType)
	for i, r := range recipes {
		for _, ingredient := range ingredients[i] {
			insert_edge(adj, r, ingredient)
		}
	}

	retArr = make([]string, 0)
	visit = make(map[string]bool)
	suppliesMap := map[string]bool{}
	for _, s := range supplies {
		suppliesMap[s] = true
	}
	recipesMap := map[string]bool{}
	for _, r := range recipes {
		recipesMap[r] = true
	}

	for _, r := range recipes {
		if dfs(r, suppliesMap, recipesMap) {
			retArr = append(retArr, r)
		}
	}
	return retArr
}

func main() {
	fmt.Println(findAllRecipes([]string{"bread", "sandwich", "burger"}, [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}}, []string{"yeast", "flour", "meat"}))
}
