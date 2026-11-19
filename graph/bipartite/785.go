package main

func isBipartite(graph [][]int) bool {
	_, ok := bipartiteColoring(graph)
	return ok
}

// bipartiteColoring 用 BFS 为无向图染色，成功时返回每个节点所属的组。
// graph 使用邻接表，节点编号为 [0, len(graph))，每条无向边应双向存储。
// 颜色 0 表示未访问，1 和 -1 表示两组；失败时返回 nil, false。
// 时间 O(V+E)，额外空间 O(V)，不修改输入。
func bipartiteColoring(graph [][]int) ([]int, bool) {
	color := make([]int, len(graph))
	queue := make([]int, 0, len(graph))

	// 图可能不连通，每个尚未染色的连通分量都需要一个 BFS 起点。
	for start := range graph {
		if color[start] != 0 {
			continue
		}
		color[start] = 1
		queue = append(queue[:0], start)

		for head := 0; head < len(queue); head++ {
			u := queue[head]
			for _, v := range graph[u] {
				if color[v] == 0 {
					// 首次入队就染色，避免同一节点重复入队。
					color[v] = -color[u]
					queue = append(queue, v)
				} else if color[v] == color[u] {
					// 已染色邻居与当前节点同色，存在冲突（包括自环）。
					return nil, false
				}
			}
		}
	}
	return color, true
}
