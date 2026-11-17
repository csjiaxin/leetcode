package l2115

import topotemplate "leetcode/graph/topo_sort/template"

// Kahn 按依赖顺序处理节点；seeds 不重复且入度为零，next 同步枚举后继。
// 返回处理数；入度和 seeds 均不修改，DP 状态由调用方初始化。
func Kahn(indegree []int, seeds []int, next func(u int, yield func(v int)), onPop func(u int), relax func(u, v int)) int {
	return topotemplate.Kahn(indegree, seeds, next, onPop, relax)
}

func ZeroIndegree(indegree []int) []int {
	return topotemplate.ZeroIndegree(indegree)
}

// https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	ids := map[string]int{}
	intern := func(name string) int {
		if id, ok := ids[name]; ok {
			return id
		}
		id := len(ids)
		ids[name] = id
		return id
	}
	for _, name := range recipes {
		intern(name)
	}
	for _, list := range ingredients {
		for _, name := range list {
			intern(name)
		}
	}
	for _, name := range supplies {
		intern(name)
	}
	graph := make([][]int, len(ids))
	indegree := make([]int, len(ids))
	recipeNames := map[int]string{}
	for i, name := range recipes {
		v := ids[name]
		recipeNames[v] = name
		for _, ingredient := range ingredients[i] {
			u := ids[ingredient]
			graph[u] = append(graph[u], v)
			indegree[v]++
		}
	}
	seeds := make([]int, 0)
	seen := make([]bool, len(ids))
	addSeed := func(u int) {
		if !seen[u] {
			seen[u] = true
			seeds = append(seeds, u)
		}
	}
	// 缺失原料不能因为入度为零就自动成为可用节点。
	for _, name := range supplies {
		addSeed(ids[name])
	}
	// 支持空配方；原题每个配方至少需要一种原料。
	for i, name := range recipes {
		if len(ingredients[i]) == 0 {
			addSeed(ids[name])
		}
	}
	answer := make([]string, 0)
	Kahn(indegree, seeds,
		func(u int, yield func(int)) {
			for _, v := range graph[u] {
				yield(v)
			}
		},
		func(u int) {
			if name, ok := recipeNames[u]; ok {
				answer = append(answer, name)
			}
		}, nil,
	)
	return answer
}
