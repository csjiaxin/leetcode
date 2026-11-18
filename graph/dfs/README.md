# DFS：按访问状态的生命周期分类

[返回图算法目录](../README.md)

递归只是执行方式，不能用一个 visited 规则套全部题。先判断状态是在**整个搜索、一个子树、还是当前路径**上有效。

| 模板族 | 状态含义与返回行为 | 题目地图 |
| --- | --- | --- |
| [图可达性](reachability/README.md) | 已经处理过的节点；返回后保留 visited | [1971](https://leetcode.com/problems/find-if-path-exists-in-graph/)、[2368](https://leetcode.com/problems/reachable-nodes-with-restrictions/)、[2316](https://leetcode.com/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/)、[1298](https://leetcode.com/problems/maximum-candies-you-can-get-from-boxes/) |
| [网格连通块](grid/readme.md) | 同一合法连通块每格一次；跨起点共享 visited | [733](https://leetcode.com/problems/flood-fill/)、[200](https://leetcode.com/problems/number-of-islands/)、[695](https://leetcode.com/problems/max-area-of-island/)、[1020](https://leetcode.com/problems/number-of-enclaves/)、[1254](https://leetcode.com/problems/number-of-closed-islands/)、[130](https://leetcode.com/problems/surrounded-regions/)、[1905](https://leetcode.com/problems/count-sub-islands/)、[827](https://leetcode.com/problems/making-a-large-island/) |
| [树上后序汇总](tree/README.md) | 子节点返回统计量，父节点合并；无向树跳过父节点 | [1443](https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/)、[1519](https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/)、[1466](https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/)、[2477](https://leetcode.com/problems/minimum-fuel-cost-to-report-to-the-capital/)；[98](https://leetcode.com/problems/validate-binary-search-tree/)、[101](https://leetcode.com/problems/symmetric-tree/)、[114](https://leetcode.com/problems/flatten-binary-tree-to-linked-list/) 旧笔记 |
| [路径回溯与递归构造](backtracking/README.md) | 当前路径上的选择；兄弟分支前恢复共享状态 | [79](https://leetcode.com/problems/word-search/)、[131](https://leetcode.com/problems/palindrome-partitioning/)；[95](https://leetcode.com/problems/unique-binary-search-trees-ii/) 返回左右子树的所有组合 |
| [记忆化 DFS](memoized/README.md) | 缓存某状态的完整答案；相同状态无需重算 | [329](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/)、[851](https://leetcode.com/problems/loud-and-rich/) |

## 一个共同外形，三种不同不变量

```go
func dfs(u int) {
    // 1. 进入：检查终止条件，建立当前状态
    for _, v := range graph[u] {
        // 2. 递归：先判断 v 是否允许进入
        dfs(v)
        // 3. 合并：按题意累计子问题结果
    }
    // 4. 离开：只有当前路径状态需要撤销
}
```

图上可达性通常是 O(V+E)；树上常数次合并为 O(n)；记忆化复杂度是状态数乘转移数；路径枚举和树组合可能是指数规模。不能仅根据“每个函数只写了一个 for”推断为线性。

本目录保留原文件是否带 `.go` 的状态。无扩展名文件是独立旧笔记，未被改名为可编译的已完成解答。[原 DFS 外部练习链接](notes.md)仍保留。
