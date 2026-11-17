# 有向可达性：分别计算每个源点的覆盖

[返回图算法目录](../README.md)

适用于“从一个事件开始，沿有向触发关系能影响多少对象”。**不变量**：固定源点的一次搜索中，每个节点只统计一次；换源点时重新建立访问表。

```go
best := 0
for start := range graph {
    seen := make([]bool, len(graph))
    var dfs func(int) int
    dfs = func(u int) int {
        if seen[u] { return 0 }
        seen[u] = true
        count := 1
        for _, v := range graph[u] { count += dfs(v) }
        return count
    }
    if count := dfs(start); count > best { best = count }
}
```

| 题目 | 原始代码 | 建图与变化点 |
| --- | --- | --- |
| [2101. 引爆最多的炸弹](https://leetcode.com/problems/detonate-the-maximum-bombs/) | [实现](2101.go) | i 的半径覆盖 j 才加 i→j；距离平方需要足够宽的整数 |

这个有向关系通常不对称。并查集会丢失方向，不能表示“从某源点能到达”。枚举所有源点耗时 O(V(V+E))；本题两两建边还需 O(V²)。
