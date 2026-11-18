# 图 DFS：可达性与连通分量

[返回 DFS 目录](../README.md)

**适用条件**：关心哪些节点能走到，或者无向图有多少连通块；不要求首次找到的路径最短。

**不变量**：进入节点后立即设置 visited，返回时不撤销。固定一次搜索中，每个节点最多扩展一次。

```go
// ReadyDFS 只在 canEnter 成立时标记；未就绪候选可由后续事件再次尝试。
func ReadyDFS[S comparable](starts []S, canEnter func(S) bool, next func(S) []S, visit func(S)) map[S]bool {
	visited := make(map[S]bool)
	var walk func(S)
	walk = func(state S) {
		if visited[state] || !canEnter(state) {
			return
		}
		visited[state] = true
		if visit != nil {
			visit(state)
		}
		for _, successor := range next(state) {
			walk(successor)
		}
	}
	for _, start := range starts {
		walk(start)
	}
	return visited
}
```

以下适配把静态图可达性写成模板调用；模板与任一适配放在同一包，不需要仓库导入：

```go
func reachable(graph [][]int, start int, blocked []bool) int {
    visited := ReadyDFS([]int{start},
        func(u int) bool { return blocked == nil || !blocked[u] },
        func(u int) []int { return graph[u] },
        nil,
    )
    return len(visited)
}
```

节点编号、起点和可选 blocked 数组按题目约束合法。枚举无向连通分量时，把 seen 提到多次 dfs 调用之外共享；分别比较多个起点的有向覆盖时，每次重置 seen。

| 题目 | 原始代码 | 要改变的部分 |
| --- | --- | --- |
| [1971. 寻找图中是否存在路径](https://leetcode.com/problems/find-if-path-exists-in-graph/) | [笔记](l1971/1971) | 从 source 出发检查 destination 是否被访问 |
| [2368. 受限条件下可到达节点的数目](https://leetcode.com/problems/reachable-nodes-with-restrictions/) | [笔记](l2368/2368) | 进入前排除 restricted 节点 |
| [2316. 统计无向图中无法互相到达点对数](https://leetcode.com/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/) | [笔记](l2316/2316) | 求每个分量大小 s，用 int64 累加 s × 之前分量总大小 |
| [1298. 你能从盒子里获得的最大糖果数](https://leetcode.com/problems/maximum-candies-you-can-get-from-boxes/) | [本地模板实现](l1298/1298.go) | 可打开状态由“持有盒子＋持有钥匙”共同决定，不能只用一次 visited |

[1298](https://leetcode.com/problems/maximum-candies-you-can-get-from-boxes/) 的实际代码也使用上面的本地 ReadyDFS。未就绪候选不能先标记 visited：只有同时持有盒子且可打开时才消费；拿到新钥匙或新盒子会再次触发候选。visit 先更新持有/可开事实，next 再返回受此次事件影响的候选。输入 status 不改写，已消费盒子不会重复计糖果。它没有新增专属测试，下面给出与实际代码一致的适配；可用性传播也可对照 [2115](https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/) · [本地](../../topo_sort/readme.md)。

```go
// https://leetcode.com/problems/maximum-candies-you-can-get-from-boxes/
func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	have := make([]bool, len(status))
	open := make([]bool, len(status))
	for box, value := range status {
		open[box] = value == 1
	}
	for _, box := range initialBoxes {
		have[box] = true
	}
	total := 0
	ReadyDFS(initialBoxes, func(box int) bool { return have[box] && open[box] }, func(box int) []int {
		next := make([]int, 0, len(keys[box])+len(containedBoxes[box]))
		next = append(next, keys[box]...)
		return append(next, containedBoxes[box]...)
	}, func(box int) {
		total += candies[box]
		for _, key := range keys[box] {
			open[key] = true
		}
		for _, child := range containedBoxes[box] {
			have[child] = true
		}
	})
	return total
}
```

时间 O(V+E)，递归空间最坏 O(V)。有环时不能只跳过父节点；那是树的特殊性质。
