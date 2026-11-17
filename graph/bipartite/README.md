# 二分图：相邻节点必须异色

[返回图算法目录](../README.md)

适用于无向图能否分成两组、组内没有边。**不变量**：已染色边的两端颜色相反；未染色连通块要另选一个起点。

```go
color := make([]int, len(graph)) // 0 未染色，1 和 -1 两组
for start := range graph {
    if color[start] != 0 { continue }
    color[start] = 1
    queue := []int{start}
    for head := 0; head < len(queue); head++ {
        u := queue[head]
        for _, v := range graph[u] {
            if color[v] == 0 {
                color[v] = -color[u]
                queue = append(queue, v)
            } else if color[v] == color[u] {
                return false
            }
        }
    }
}
return true
```

| 题目 | 代码与测试 | 变化点 |
| --- | --- | --- |
| [785. 判断二分图](https://leetcode.com/problems/is-graph-bipartite/) | [实现](785.go) · [测试](785_test.go) | 遍历所有连通块；自环立即冲突 |

时间 O(V+E)，空间 O(V)。也能用 DFS 染色，颜色约束不变；检测普通有向环应使用拓扑排序或三色 DFS。
