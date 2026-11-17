# 图 DFS：可达性与连通分量

[返回 DFS 目录](../README.md)

**适用条件**：关心哪些节点能走到，或者无向图有多少连通块；不要求首次找到的路径最短。

**不变量**：进入节点后立即设置 visited，返回时不撤销。固定一次搜索中，每个节点最多扩展一次。

```go
func reachable(graph [][]int, start int, blocked []bool) int {
    seen := make([]bool, len(graph))
    var dfs func(int) int
    dfs = func(u int) int {
        if seen[u] || (blocked != nil && blocked[u]) { return 0 }
        seen[u] = true
        size := 1
        for _, v := range graph[u] { size += dfs(v) }
        return size
    }
    return dfs(start)
}
```

节点编号、起点和可选 blocked 数组按题目约束合法。枚举无向连通分量时，把 seen 提到多次 dfs 调用之外共享；分别比较多个起点的有向覆盖时，每次重置 seen。

| 题目 | 原始代码 | 要改变的部分 |
| --- | --- | --- |
| [1971. 寻找图中是否存在路径](https://leetcode.com/problems/find-if-path-exists-in-graph/) | [笔记](l1971/1971) | 从 source 出发检查 destination 是否被访问 |
| [2368. 受限条件下可到达节点的数目](https://leetcode.com/problems/reachable-nodes-with-restrictions/) | [笔记](l2368/2368) | 进入前排除 restricted 节点 |
| [2316. 统计无向图中无法互相到达点对数](https://leetcode.com/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/) | [笔记](l2316/2316) | 求每个分量大小 s，用 int64 累加 s × 之前分量总大小 |
| [1298. 你能从盒子里获得的最大糖果数](https://leetcode.com/problems/maximum-candies-you-can-get-from-boxes/) | [旧实现](l1298/1298.go) | 可打开状态由“持有盒子＋持有钥匙”共同决定，不能只用一次 visited |

1298 保留原始实现供复盘，尚无测试保证；位掩码是否包含某位应检查按位与 `mask & bit`，原始代码中的按位或条件需要另行纠正。这类可用性变化也可对照 [2115 的事件传播](../../topo_sort/readme.md)。

时间 O(V+E)，递归空间最坏 O(V)。有环时不能只跳过父节点；那是树的特殊性质。
