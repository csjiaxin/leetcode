package l2115

// Kahn 不修改 indegree 和 seeds，返回处理的节点数。
// seeds 必须互不重复且入度为零，next 必须同步枚举每条出边一次。
// indegree 必须与这些边的数量一致（包括重边），节点编号必须合法。
// relax 在递减后继入度前调用；onPop 在全部前置贡献就绪后调用。
// DP 初值和边上的转移由调用方提供，回调不能修改图结构。
func Kahn(indegree []int, seeds []int, next func(u int, yield func(v int)), onPop func(u int), relax func(u, v int)) int {
	remaining := append([]int(nil), indegree...)
	queue := append([]int(nil), seeds...)
	for head := 0; head < len(queue); head++ {
		u := queue[head]
		if onPop != nil {
			onPop(u)
		}
		next(u, func(v int) {
			if relax != nil {
				relax(u, v)
			}
			remaining[v]--
			if remaining[v] == 0 {
				queue = append(queue, v)
			}
		})
	}
	return len(queue)
}

// ZeroIndegree 收集普通 DAG 遍历的全部源点；不修改入度。
// 依赖外部 supplies 的题目应由调用方显式构造 seeds。
func ZeroIndegree(indegree []int) []int {
	seeds := make([]int, 0)
	for u, degree := range indegree {
		if degree == 0 {
			seeds = append(seeds, u)
		}
	}
	return seeds
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
