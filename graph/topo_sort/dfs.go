package main

func TopoSortDFS(graph [][]int) ([]int, bool) {
	n := len(graph)
	state := make([]int, n)
	order := make([]int, 0, n)

	var dfs func(int) bool
	dfs = func(u int) bool {
		if state[u] == 1 {
			return false // 回到当前递归路径上的节点：存在环。
		}
		if state[u] == 2 {
			return true // 已处理完成，无需重复访问。
		}

		state[u] = 1

		for _, v := range graph[u] {
			if !dfs(v) {
				return false
			}
		}

		state[u] = 2
		order = append(order, u) // 后序：所有后继处理完再加入。
		return true
	}

	// 遍历所有节点，覆盖不连通的部分。
	for u := 0; u < n; u++ {
		if state[u] == 0 && !dfs(u) {
			return nil, false
		}
	}

	// 反转后序结果，得到拓扑序。
	for l, r := 0, len(order)-1; l < r; l, r = l+1, r-1 {
		order[l], order[r] = order[r], order[l]
	}

	return order, true
}
