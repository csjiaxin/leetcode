# 记忆化 DFS：把递归结果缓存为状态答案

[返回 DFS 目录](../README.md)

适用条件：相同状态的答案不依赖当前进入路径，且依赖关系无环，或另外设计了正确的环处理。

**不变量**：缓存命中意味着该状态的所有后继已经处理，缓存的是完整答案，不是“曾经来过”。

```go
func MemoDFS[S comparable, V any](evaluate func(S, func(S) V) V) func(S) V {
	memo := make(map[S]V)
	var solve func(S) V
	solve = func(state S) V {
		if value, ok := memo[state]; ok {
			return value
		}
		value := evaluate(state, solve)
		memo[state] = value
		return value
	}
	return solve
}
```

在 DAG 中求从 start 出发的最长节点数，只提供状态转移：

```go
func longestFrom(graph [][]int, start int) int {
    solve := MemoDFS(func(u int, recur func(int) int) int {
        answer := 1
        for _, v := range graph[u] {
            answer = max(answer, 1+recur(v))
        }
        return answer
    })
    return solve(start)
}
```

初始化、合法后继与合并规则放在 evaluate 回调中；以 map 的 ok 判断缓存存在性，合法零值不会被误认为未计算。模板与适配放在同一包，实际题解需要的辅助函数保留本地副本。若图可能有环，还需要另外设计状态求解或访问中检测；仅在返回时缓存不能阻止递归环。

| 题目 | 原始代码 | 依赖无环的原因与聚合 |
| --- | --- | --- |
| [329. 矩阵中的最长递增路径](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/) | [DFS 笔记](l329/329) · [拓扑版本与测试](../../topo_sort/l329/329.go) | 只走严格更大的格子；长度 = 1 + 后继最大长度 |
| [851. 喧闹和富有](https://leetcode.com/problems/loud-and-rich/) | [DFS 笔记](l851/851) | 按题意的财富偏序查更富的人；缓存最安静者 |

若每个状态只完成一次，复杂度是状态数加依赖边数；[329](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/) 为 O(mn)，[851](https://leetcode.com/problems/loud-and-rich/) 为 O(V+E)。DFS 与拓扑 DP 可以解决相同 DAG，但计算方向与初始化要与状态定义保持一致。
