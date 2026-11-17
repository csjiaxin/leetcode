# 最小生成树：连接全部节点的最小总代价

[返回图算法目录](../README.md)

适用于无向带权图；边可以为负。目标是选出 n-1 条无环边连接全部节点，不是求某个源点的最短路。

| 模板 | 固定骨架 | 不变量／变化点 |
| --- | --- | --- |
| [Kruskal](kruskal.go) | 边排序 → 找集合根 → 不同根才选边合并 | 已选边构成森林；用[并查集](../union_find/union_find.go)判断环 |
| [Prim](prim.go) | 从一点开始 → 弹出最小接入边 → 跳过已加入节点 → 扩展边 | 当前已选节点构成一棵树；最便宜的跨割边安全 |
| [稠密隐式图 Prim](l1584/1584.go) | 数组扫描最小 key → 接入节点 → 用 cost 更新其他 key | key[v] 是当前树到 v 的最小单边代价，无需显式存所有边 |

下面两段展示控制流；完整可编译实现见上表链接。通用 `kruskal.go` 与 `prim.go` 组成同一个包：`MSTEdge`、`NewUnionFind` 和 `checkedAdd` 定义在 `kruskal.go`，Prim 复用其中的 `MSTEdge`／`checkedAdd`。单独复制算法时需一并带上所需声明；同包组装时每个顶层声明只保留一次，无需导入仓库内部包。

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
| [1584. 连接所有点的最小费用](https://leetcode.com/problems/min-cost-to-connect-all-points/) · [实现](l1584/1584.go) | 曼哈顿距离作为隐式完全图的 cost，调用同文件 DensePrim；保留演示 main，主解完整独立 |
| Kruskal / Prim 通用文件 | 已补齐 NewUnionFind／checkedAdd，同包可编译；Kruskal 排序后合并，堆 Prim 选跨割边，两者均检查总权溢出及不连通 |

复杂度：Kruskal 为 O(E log E)，堆 Prim 为 O(V+E log(E+1))；稠密图可用 O(V²) 时间、O(V) 额外空间的数组 Prim。

## 稠密隐式图：固定 DensePrim 与题目适配

完全图有 O(n²) 条边，但边权可以按需计算，无需先建立邻接矩阵或把所有边放入堆。先把节点0加入树，令 `key[v]=cost(0,v)`；随后反复选择树外 key 最小的节点，再以新接入节点更新其余 key。

**不变量**：每轮开始时，key[v] 是树内任意节点到树外 v 的最小单边权。它是接入边的代价，不是从节点0出发的整条路径长度；更新为 `min(key[v],cost(u,v))`，不能写成 Dijkstra 的路径加法。

以下 helper 要求节点编号为0..n-1、n非负，cost 对称且为每对节点提供有限边权；完全图在n>0时自然连通。边权可以为负，所有边权和最终总和须在int范围内。若需处理缺边或不连通，使用上面的通用带错误返回值版本，不能拿某个合法权值充当缺边标记。

README 只声明一次 helper，选择一题适配函数接在同一 `package` 中即可；1584 主文件已包含完整副本和演示 main，可独立使用，无需仓库内部包。

```go
package template

// DensePrim returns the MST cost of a complete undirected graph.
// cost(i, j) must be symmetric, and all edge costs and their sum must fit in int.
func DensePrim(n int, cost func(int, int) int) int {
	if n <= 1 {
		return 0
	}
	included := make([]bool, n)
	key := make([]int, n)
	included[0] = true
	for vertex := 1; vertex < n; vertex++ {
		key[vertex] = cost(0, vertex)
	}
	total := 0
	for count := 1; count < n; count++ {
		next := -1
		for vertex := 0; vertex < n; vertex++ {
			if !included[vertex] && (next == -1 || key[vertex] < key[next]) {
				next = vertex
			}
		}
		included[next] = true
		total += key[next]
		// key[v] is the cheapest single edge from the current tree to v.
		for vertex := 0; vertex < n; vertex++ {
			if !included[vertex] {
				weight := cost(next, vertex)
				if weight < key[vertex] {
					key[vertex] = weight
				}
			}
		}
	}
	return total
}
```

### 1584：只提供曼哈顿距离

```go
func minCostConnectPoints(points [][]int) int {
	return DensePrim(len(points), func(i, j int) int {
		dx := points[i][0] - points[j][0]
		if dx < 0 {
			dx = -dx
		}
		dy := points[i][1] - points[j][1]
		if dy < 0 {
			dy = -dy
		}
		return dx + dy
	})
}
```

例子 `[[0,0],[2,2],[3,10],[5,2],[7,0]]` 的最小费用为20。空点集或单点为0。每次加入一点只扫描剩余节点，时间O(n²)、额外空间O(n)，不修改输入。

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
