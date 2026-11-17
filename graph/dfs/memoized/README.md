# 记忆化 DFS：把递归结果缓存为状态答案

[返回 DFS 目录](../README.md)

适用条件：相同状态的答案不依赖当前进入路径，且依赖关系无环，或另外设计了正确的环处理。

**不变量**：缓存命中意味着该状态的所有后继已经处理，缓存的是完整答案，不是“曾经来过”。

```go
memo := make([]int, n)
done := make([]bool, n)
var dfs func(int) int
dfs = func(u int) int {
    if done[u] { return memo[u] }
    answer := base(u)
    for _, v := range next(u) {
        answer = combine(answer, dfs(v))
    }
    memo[u], done[u] = answer, true
    return answer
}
```

`base/next/combine` 是题目需要填写的部分。若图可能有环，还需要访问中状态检测；仅在返回时置 done 不能阻止递归环。

| 题目 | 原始代码 | 依赖无环的原因与聚合 |
| --- | --- | --- |
| [329. 矩阵中的最长递增路径](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/) | [DFS 笔记](l329/329) · [拓扑版本与测试](../../topo_sort/l329/329.go) | 只走严格更大的格子；长度 = 1 + 后继最大长度 |
| [851. 喧闹和富有](https://leetcode.com/problems/loud-and-rich/) | [DFS 笔记](l851/851) | 按题意的财富偏序查更富的人；缓存最安静者 |

若每个状态只完成一次，复杂度是状态数加依赖边数；329 为 O(mn)，851 为 O(V+E)。DFS 与拓扑 DP 可以解决相同 DAG，但计算方向与初始化要与状态定义保持一致。
