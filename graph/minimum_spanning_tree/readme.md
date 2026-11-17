# 最小生成树：连接全部节点的最小总代价

[返回图算法目录](../README.md)

适用于无向带权图；边可以为负。目标是选出 n-1 条无环边连接全部节点，不是求某个源点的最短路。

| 模板 | 固定骨架 | 不变量／变化点 |
| --- | --- | --- |
| [Kruskal](kruskal.go) | 边排序 → 找集合根 → 不同根才选边合并 | 已选边构成森林；用[并查集](../union_find/union_find.go)判断环 |
| [Prim](prim.go) | 从一点开始 → 弹出最小接入边 → 跳过已加入节点 → 扩展边 | 当前已选节点构成一棵树；最便宜的跨割边安全 |

```go
// Kruskal 骨架：edges 已按权重升序排列，uf 已初始化。
for _, edge := range edges {
    if !uf.Union(edge.U, edge.V) { continue }
    chosen = append(chosen, edge)
    total += edge.Weight // 根据数值范围检查溢出
    if len(chosen) == n-1 { break }
}
// chosen 少于 n-1，说明图不连通，不能返回完整生成树。
```

```go
// Prim 骨架：队列初始含一条接入起点的 0 权虚拟边。
for queueNotEmpty() {
    edge := popMinimum()
    if included[edge.to] { continue }
    included[edge.to] = true
    total += edge.weight
    // 将 edge.to 指向尚未加入节点的边放入候选堆。
}
```

这里的 Prim 候选是“接入一条边”的权重。Dijkstra 候选是“从源点走来的整条路径”的代价，不能用同一个松弛公式替换。

| 题目／文件 | 现状与变化点 |
| --- | --- |
| [1584. 连接所有点的最小费用](https://leetcode.com/problems/min-cost-to-connect-all-points/) · [实现](l1584/1584.go) | 曼哈顿距离构造完全图的 Prim 旧写法，单独成包保留其 main 和堆类型 |
| Kruskal / Prim 通用文件 | 原代码原样保留，原来缺少 NewUnionFind／checkedAdd 依赖；本次归类不补完旧草稿 |

复杂度：Kruskal 为 O(E log E)，堆 Prim 为 O(V+E log(E+1))；稠密图可考虑 O(V²) 的数组 Prim。

## 原始笔记（保留）

https://leetcode.com/tag/minimum-spanning-tree/

**最小生成树（Minimum Spanning Tree，MST）**用于无向带权图：选择一些边，连接所有节点，使总权重最小。

算法	核心思路	常见时间复杂度
Kruskal	按边权排序，用并查集合并不同集合	O(E log E)
Prim	从一个节点开始，每次用最便宜的边接入一个新节点	数组版 O(V²)；堆版 O(E log V)


一 Prim

Prim:
key[v] = 将 v 接入当前生成树所需的最小边权
更新：key[v] = min(key[v], weight(u, v))

Dijkstra:
dist[v] = 从起点到 v 的最短路径长度
更新：dist[v] = min(dist[v], dist[u] + weight(u, v))


二 Kruskal
对于 n 个节点，生成树：
- 包含 n-1 条边。
- 连接所有节点。
- 不包含环。
可以使用 Kruskal 算法：边按权重排序 + 并查集判环。
1. 按边权从小到大处理。
2. 如果两个端点属于不同集合，就选择这条边并合并集合。
3. 选满 n-1 条边后结束。
