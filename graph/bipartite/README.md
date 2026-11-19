# 二分图：BFS 染色模板与 LeetCode 题单

[返回图算法目录](../README.md) · [完整实现](785.go) · [测试](785_test.go)

## 1. 先区分两个问题

- **二分图判定**：能否将所有顶点分成两组，使每条边两端属于不同组？用 BFS / DFS 两色染色。
- **二分图匹配**：左右两组之间挑选一些边，使每个顶点最多被选一次；若要求配对最多，则求最大匹配。需要增广路等算法，不能只靠染色。

本文实现的是第一类：**无向图的 BFS 染色判定**。后面的题单同时列出匹配与进阶建模，供继续学习，并不表示它们都能直接套本模板。

## 2. 如何识别二分图判定

常见描述：分成两组；有冲突的人不能同组；相邻节点必须取不同颜色。

将对象作为顶点，将“必须不同组”的关系作为无向边：

```text
color[u] =  0：尚未访问
color[u] =  1：第一组
color[u] = -1：第二组

对于每条边 u -- v，必须满足 color[v] = -color[u]。
```

**等价性质：无向图是二分图，当且仅当它不包含奇数长度的环。**

- 树、森林、偶环都可以染色。
- 三角形等奇环不能染色。
- 自环要求自己与自己异色，立即失败；重复边不影响结论。
- 空图和只有孤立点的图都返回 `true`，不要求两组都非空。

## 3. 完整 Go 模板：BFS 染色

输入约定：`graph` 是无向图邻接表，节点编号为 `[0, len(graph))`，无向边在两端都存储；不修改输入。这里不对越界编号或非对称邻接表等不符合约定的输入进行校验。

`bipartiteColoring` 成功时返回完整颜色数组；失败时返回 `nil, false`，避免误用局部染色结果。只需要布尔答案时使用 `isBipartite`，可直接用于 785。

```go
func isBipartite(graph [][]int) bool {
    _, ok := bipartiteColoring(graph)
    return ok
}

func bipartiteColoring(graph [][]int) ([]int, bool) {
    color := make([]int, len(graph))
    queue := make([]int, 0, len(graph))

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
                    color[v] = -color[u]
                    queue = append(queue, v)
                } else if color[v] == color[u] {
                    return nil, false
                }
            }
        }
    }
    return color, true
}
```

### 为什么这样写？

1. **遍历所有连通分量**：从一个点 BFS 只能覆盖一个分量，别的分量可能包含奇环。
2. **起点可以任意设为 1**：每个分量的两种颜色整体互换不影响合法性，不同分量无需协调颜色。
3. **入队时染色**：`color` 同时承担 `visited` 的作用，每个顶点至多入队一次。
4. **未染色则赋相反色，已染色则检查冲突**：不能忽略已访问邻居，也不能覆盖它的颜色。
5. **不需要显式逐层循环**：我们只关心颜色，不需要统计层数或距离；相邻层颜色交替已由 `-color[u]` 实现。
6. **按 head 遍历实际队列节点**：当前顶点是 `queue[head]`，不是 `head`。每个分量开始时清空队列长度，复用容量，不清空全局颜色。

### BFS 分层与奇环

从某个起点出发，偶数距离层染为 1，奇数距离层染为 -1。无向图一条边两端的 BFS 距离最多相差 1；如果边连接了同一层的两个点，就产生同色冲突，并能找到奇环。

例如三角形 `0--1--2--0`：先将 0 染为 1，将 1、2 染为 -1；检查边 `1--2` 时冲突。四边形则可以交替染色。

### 正确性与复杂度

- 搜索过程中，已检查且未引发失败的边始终满足两端异色。
- 若发生同色冲突，已有路径施加的异色约束与当前边矛盾；等价地存在奇环，因此无法二分。
- 若全部分量检查结束，每个顶点均已染色、每条边两端均异色，颜色数组就是合法分组。
- 时间 **O(V+E)**：每个顶点最多入队一次，每条无向边被检查两次。
- 额外空间 **O(V)**：颜色数组与队列；不计输入邻接表。

## 4. 模板如何迁移

### 785. 判断二分图：直接调用

输入已经是无向图邻接表，直接调用 `isBipartite(graph)`。见 [实现](785.go) 与 [测试](785_test.go)。

### 886. 可能的二分法：先建图

人编号是 `1..n`，把 `dislikes` 中的关系转换成双向边，再复用同一个 `bipartiteColoring`：

```go
func possibleBipartition(n int, dislikes [][]int) bool {
    graph := make([][]int, n)
    for _, edge := range dislikes {
        u, v := edge[0]-1, edge[1]-1
        graph[u] = append(graph[u], v)
        graph[v] = append(graph[v], u)
    }
    _, ok := bipartiteColoring(graph)
    return ok
}
```

独立提交时，将第 3 节的 `bipartiteColoring` 一并复制，无需 `isBipartite`。建图后的总时间与空间均为 O(n+m)，其中 m 是关系数。

### 如果需要返回两组节点

成功后遍历颜色数组：颜色为 1 的放进第一组，-1 的放进第二组。这个结果只保证约束成立，**不保证人数平衡，也不优化组大小**。

### 2493. 将节点分成尽可能多的组：染色只是第一步

1. 建图并判断二分图；有奇环则返回 -1。
2. 收集每个连通分量，在分量内枚举每个顶点作为起点，分别 BFS 求最大距离。
3. 当前分量最多分组数 = 这些最大距离的最大值 + 1；最终将各分量的结果相加。孤立点贡献 1。

理由：二分图中 BFS 不会出现同层边，因此距离加 1 是合法分组；任意合法分组的最大组号差又不超过图的直径。枚举起点可以求出直径。一般图不能不加证明地用树的“两次 BFS 求直径”替代枚举。

总时间 O(V(V+E)) 是一个上界；染色本身仍只需 O(V+E)。每个起点的距离数组应独立初始化，不能复用染色数组来判断是否访问。

## 5. LeetCode 分类题单

以下是代表性练习，不是穷尽列表。🔒 为会员题，开放情况以平台为准。除 785 外，下表列的是学习建议，不代表本目录已实现。

### A. 二分图判定：BFS / DFS 染色

| 题目 | 难度 | 关键迁移 |
| --- | --- | --- |
| [785. 判断二分图](https://leetcode.com/problems/is-graph-bipartite/) | 中等 | 标准两色染色，遍历所有连通分量 |
| [886. 可能的二分法](https://leetcode.com/problems/possible-bipartition/) | 中等 | 将不喜欢关系建成无向图；也可用并查集 |
| [2493. 将节点分成尽可能多的组](https://leetcode.com/problems/divide-nodes-into-the-maximum-number-of-groups/) | 困难 | 二分图判定 + 每个分量枚举起点做 BFS 分层 |

### B. 最大匹配与消除冲突

| 题目 | 难度 | 核心考点 |
| --- | --- | --- |
| [1820. 最多邀请的个数](https://leetcode.com/problems/maximum-number-of-accepted-invitations/) 🔒 | 中等 | 标准二分图最大匹配，增广路 |
| [1349. 参加考试的最大学生数](https://leetcode.com/problems/maximum-students-taking-exam/) | 困难 | 按列奇偶划分座位，作弊关系连边；最大独立集转最大匹配，也常用状压 DP |
| [2123. Minimum Operations to Remove Adjacent Ones in Matrix](https://leetcode.com/problems/minimum-operations-to-remove-adjacent-ones-in-matrix/) 🔒 | 困难 | 网格按 (行+列) 奇偶划分；删除最少的 1 即最小顶点覆盖 |

对二分图，Kőnig 定理及其推论是：

```text
最小顶点覆盖大小 = 最大匹配大小
最大独立集大小 = 顶点数 - 最小顶点覆盖大小
               = 顶点数 - 最大匹配大小
```

顶点覆盖是选点覆盖所有边；独立集是选点且任意两点之间没有边。所以 1349 是“保留最多不冲突座位”，2123 是“删除最少节点消除全部冲突”。这些是后续匹配知识，本文未实现匹配模板。

### C. 带权匹配与覆盖建模

| 题目 | 难度 | 与二分图的关系 |
| --- | --- | --- |
| [1947. 最大兼容性评分和](https://leetcode.com/problems/maximum-compatibility-score-sum/) | 中等 | 学生与导师一一配对，最大权完美匹配；常用状压 DP |
| [1879. 两个数组最小的异或值之和](https://leetcode.com/problems/minimum-xor-sum-of-two-arrays/) | 困难 | 数组元素一一配对，最小权完美匹配；常用状压 DP |
| [1595. 连通两组点的最小成本](https://leetcode.com/problems/minimum-cost-to-connect-two-groups-of-points/) | 困难 | 最小权边覆盖，每个点至少连接一次；常用状压 DP |
| [3276. 选择矩阵中单元格的最大得分](https://leetcode.com/problems/select-cells-in-grid-with-maximum-score/) | 困难 | 行与数值作为左右两组，带权匹配；也可用状压 DP |

这里不是普通的无权最大匹配。特别注意：1595 的**边覆盖**是选边覆盖顶点，与上面的**顶点覆盖**不同，不能直接套 Kőnig 定理。

## 6. 易错点与推荐顺序

- 只从节点 0 出发，漏掉不连通分量中的奇环。
- 出队才染色，造成重复入队；或者对已染色节点重新赋色，掩盖冲突。
- 以为“有环就不是二分图”：只有奇环不允许，偶环允许。
- 以为只需检查 BFS 树边：非树边同样可能暴露冲突。
- 混淆无向二分图与有向图判环：普通有向环检测应使用拓扑排序或三色 DFS，不能照搬本模板。
- 混淆网格两种划分：四邻接网格通常按 `(r+c)%2` 划分；1349 的冲突边要按列奇偶划分。
- 染色只判定能否二分，不求最大匹配，也不保证两组人数相等。

推荐顺序：

```text
染色判定：785 → 886 → 2493
最大匹配：1820 → 1349 → 2123
带权配对 / 状压 DP：1947 → 1879 → 3276 → 1595
```

时间有限时，优先做 785、886、2493；需要继续学习匹配，再做 1820、1349。

## 7. 本地验证

```sh
go test ./graph/bipartite -count=1 -cover
```

测试覆盖空图、孤立点、不连通图、奇偶环、自环、重复边、队列访问顺序、输入不变性及返回颜色的合法性；还穷举最多 5 个顶点的所有简单无向图，与暴力枚举两组分配的答案对照。
